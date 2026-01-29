package app_resource

import (
	"context"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ByTeamCache stores AppResources keyed by team_id for efficient per-team caching.
var ByTeamCache *cache.KeyedCache[AppResourceModel]

// AppInstanceToTeam maps app_instance_id → team_id (shared with app_resource_role).
var AppInstanceToTeam *cache.MappingCache

// importLoadLock prevents thundering herd during parallel imports.
// Only one goroutine will fetch the initial resource to discover the team_id.
var importLoadLock *cache.ImportCache[AppResourceModel]

// InitCache initializes the app resource cache. Call this from provider.Configure().
func InitCache() {
	ByTeamCache = cache.NewKeyedCache[AppResourceModel]()
	AppInstanceToTeam = cache.NewMappingCache()
	importLoadLock = cache.NewImportCache[AppResourceModel]()
}

// FindInLoadedCachesModel returns the model if found in any loaded cache.
// This is useful for Read operations where app_instance_id may not be available.
func FindInLoadedCachesModel(id string) (*AppResourceModel, bool) {
	if ByTeamCache == nil {
		return nil, false
	}
	item, _ := ByTeamCache.FindInLoadedCaches(id)
	return item, item != nil
}

// GetCached retrieves an app resource from the per-team cache, loading via List API if needed.
// Returns (model, found, error). If the cache fails to load, error is non-nil.
// If the resource doesn't exist, found is false.
func GetCached(ctx context.Context, client *serval.Client, id string, appInstanceID string) (*AppResourceModel, bool, error) {
	if ByTeamCache == nil {
		return nil, false, nil // Cache not initialized, caller should fall back to API
	}

	// FAST PATH: Check if already in any loaded cache before doing expensive API lookups
	if item, teamID := ByTeamCache.FindInLoadedCaches(id); item != nil {
		tflog.Debug(ctx, "[AppResource] GetCached FAST PATH hit", map[string]interface{}{"id": id, "teamID": teamID})
		return item, true, nil
	}

	// Step 1: Get the team_id for this app_instance_id
	teamID, err := getTeamIDForAppInstance(ctx, client, appInstanceID)
	if err != nil {
		return nil, false, err
	}

	// Step 2: Get or create the cache for this team
	teamCache := ByTeamCache.GetOrCreateCache(teamID)

	// Step 3: Load all resources for this team if not already loaded
	return teamCache.GetOrLoad(id, func() (map[string]*AppResourceModel, error) {
		return fetchAllAppResourcesForTeam(ctx, client, teamID)
	})
}

// GetCachedForImport retrieves an app resource from cache when we only have the ID.
// Used by ImportState where we don't know the app_instance_id upfront.
// On first import for a team, does a GET to learn the app_instance_id, then loads the team cache.
// Uses ImportCache to prevent thundering herd - only one goroutine fetches the initial resource.
func GetCachedForImport(ctx context.Context, client *serval.Client, id string) (*AppResourceModel, bool, error) {
	if ByTeamCache == nil || importLoadLock == nil {
		tflog.Debug(ctx, "[AppResource] Cache not initialized, falling back to API", map[string]interface{}{"id": id})
		return nil, false, nil // Cache not initialized, caller should fall back to API
	}

	// Step 1: Check if this resource is already in any loaded team cache
	if item, teamID := ByTeamCache.FindInLoadedCaches(id); item != nil {
		tflog.Debug(ctx, "[AppResource] CACHE HIT - found in loaded cache", map[string]interface{}{"id": id, "teamID": teamID})
		// Found in cache - also ensure mappings are populated
		if !item.AppInstanceID.IsNull() && !item.AppInstanceID.IsUnknown() {
			AppInstanceToTeam.Set(item.AppInstanceID.ValueString(), teamID)
		}
		return item, true, nil
	}

	// Step 2: Check if import lock already has the team_id (another goroutine discovered it)
	if importLoadLock.IsInitialized() {
		teamID := importLoadLock.GetParentKey()
		if teamID != "" {
			tflog.Debug(ctx, "[AppResource] Lock initialized, using discovered team_id", map[string]interface{}{"id": id, "teamID": teamID})
			teamCache := ByTeamCache.GetOrCreateCache(teamID)
			return teamCache.GetOrLoad(id, func() (map[string]*AppResourceModel, error) {
				tflog.Info(ctx, "[AppResource] LOADING ALL via List API", map[string]interface{}{"teamID": teamID})
				return fetchAllAppResourcesForTeam(ctx, client, teamID)
			})
		}
	}

	// Step 3: Try to acquire the load lock - only one goroutine should fetch
	if importLoadLock.AcquireLoadLock() {
		tflog.Info(ctx, "[AppResource] ACQUIRED LOCK - will fetch initial resource", map[string]interface{}{"id": id})
		// This goroutine will do the initial fetch to discover team_id
		resource, err := client.AppResources.Get(ctx, id,
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			tflog.Error(ctx, "[AppResource] Initial Get failed", map[string]interface{}{"id": id, "error": err.Error()})
			importLoadLock.CompleteLoad("") // Signal failure
			return nil, false, err
		}

		// Get team_id from app_instance_id
		teamID, err := getTeamIDForAppInstance(ctx, client, resource.AppInstanceID)
		if err != nil {
			tflog.Error(ctx, "[AppResource] Failed to get team_id", map[string]interface{}{"id": id, "error": err.Error()})
			importLoadLock.CompleteLoad("") // Signal failure
			return nil, false, err
		}

		tflog.Info(ctx, "[AppResource] Discovered team_id, completing lock", map[string]interface{}{"id": id, "teamID": teamID})
		// Complete the lock so other goroutines can proceed
		importLoadLock.CompleteLoad(teamID)

		// Now load the full cache and return the item
		return GetCached(ctx, client, id, resource.AppInstanceID)
	}

	// Step 4: Another goroutine is loading - wait for it
	tflog.Debug(ctx, "[AppResource] WAITING for lock holder to complete", map[string]interface{}{"id": id})
	teamID := importLoadLock.WaitForLoad()
	if teamID == "" {
		tflog.Warn(ctx, "[AppResource] Lock holder failed, falling back to individual Get", map[string]interface{}{"id": id})
		// Loading failed, fall back to individual API call
		resource, err := client.AppResources.Get(ctx, id,
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			return nil, false, err
		}
		return GetCached(ctx, client, id, resource.AppInstanceID)
	}

	tflog.Debug(ctx, "[AppResource] Lock holder completed, using cache", map[string]interface{}{"id": id, "teamID": teamID})
	// Use the discovered team_id to get from cache
	teamCache := ByTeamCache.GetOrCreateCache(teamID)
	return teamCache.GetOrLoad(id, func() (map[string]*AppResourceModel, error) {
		tflog.Info(ctx, "[AppResource] LOADING ALL via List API (from waiter)", map[string]interface{}{"teamID": teamID})
		return fetchAllAppResourcesForTeam(ctx, client, teamID)
	})
}

// getTeamIDForAppInstance looks up the team_id for an app_instance_id.
// Uses cached mapping if available, otherwise fetches from API.
func getTeamIDForAppInstance(ctx context.Context, client *serval.Client, appInstanceID string) (string, error) {
	// Check mapping cache first
	if teamID, found := AppInstanceToTeam.Get(appInstanceID); found {
		return teamID, nil
	}

	// Fetch app instance to get team_id
	appInstance, err := client.AppInstances.Get(ctx, appInstanceID,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return "", err
	}

	// Cache the mapping for future use
	AppInstanceToTeam.Set(appInstanceID, appInstance.TeamID)
	return appInstance.TeamID, nil
}

// fetchAllAppResourcesForTeam retrieves all app resources for a team via the List API.
func fetchAllAppResourcesForTeam(ctx context.Context, client *serval.Client, teamID string) (map[string]*AppResourceModel, error) {
	result := make(map[string]*AppResourceModel)
	var pageToken *string

	for {
		params := serval.AppResourceListParams{
			TeamID:   serval.String(teamID),
			PageSize: serval.Int(5000),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.AppResources.List(ctx, params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			return nil, err
		}

		bytes, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}

		var page struct {
			Data          []AppResourceModel `json:"data"`
			NextPageToken *string            `json:"nextPageToken,omitempty"`
		}
		if err := apijson.Unmarshal(bytes, &page); err != nil {
			return nil, err
		}

		for i := range page.Data {
			item := page.Data[i]
			result[item.ID.ValueString()] = &item
			// Also cache the app_instance_id → team_id mapping for each resource
			if !item.AppInstanceID.IsNull() && !item.AppInstanceID.IsUnknown() {
				AppInstanceToTeam.Set(item.AppInstanceID.ValueString(), teamID)
			}
		}

		if page.NextPageToken == nil || *page.NextPageToken == "" {
			break
		}
		pageToken = page.NextPageToken
	}

	return result, nil
}
