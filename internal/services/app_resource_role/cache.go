package app_resource_role

import (
	"context"
	"io"
	"net/http"

	"github.com/ServalHQ/serval-go"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/terraform-provider-serval/internal/apijson"
	"github.com/ServalHQ/terraform-provider-serval/internal/cache"
	"github.com/ServalHQ/terraform-provider-serval/internal/logging"
	"github.com/ServalHQ/terraform-provider-serval/internal/services/app_resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ByTeamCache stores AppResourceRoles keyed by team_id for efficient per-team caching.
var ByTeamCache *cache.KeyedCache[AppResourceRoleModel]

// ResourceToAppInstance maps resource_id → app_instance_id.
var ResourceToAppInstance *cache.MappingCache

// importLoadLock prevents thundering herd during parallel imports.
var importLoadLock *cache.ImportCache[AppResourceRoleModel]

// InitCache initializes the app resource role cache. Call this from provider.Configure().
func InitCache() {
	ByTeamCache = cache.NewKeyedCache[AppResourceRoleModel]()
	ResourceToAppInstance = cache.NewMappingCache()
	importLoadLock = cache.NewImportCache[AppResourceRoleModel]()
}

// FindInLoadedCachesModel returns the model if found in any loaded cache.
// This is useful for Read operations where resource_id may not be available.
func FindInLoadedCachesModel(id string) (*AppResourceRoleModel, bool) {
	if ByTeamCache == nil {
		return nil, false
	}
	item, _ := ByTeamCache.FindInLoadedCaches(id)
	return item, item != nil
}

// GetCached retrieves an app resource role from the per-team cache, loading via List API if needed.
// Returns (model, found, error). If the cache fails to load, error is non-nil.
// If the role doesn't exist, found is false.
func GetCached(ctx context.Context, client *serval.Client, id string, resourceID string) (*AppResourceRoleModel, bool, error) {
	if ByTeamCache == nil {
		return nil, false, nil // Cache not initialized, caller should fall back to API
	}

	// FAST PATH: Check if already in any loaded cache before doing expensive API lookups
	if item, teamID := ByTeamCache.FindInLoadedCaches(id); item != nil {
		tflog.Debug(ctx, "[AppResourceRole] GetCached FAST PATH hit", map[string]interface{}{"id": id, "teamID": teamID})
		return item, true, nil
	}

	// Step 1: Get the app_instance_id for this resource
	appInstanceID, err := getAppInstanceIDForResource(ctx, client, resourceID)
	if err != nil {
		return nil, false, err
	}

	// Step 2: Get the team_id for this app_instance_id (uses app_resource's mapping cache)
	teamID, err := getTeamIDForAppInstance(ctx, client, appInstanceID)
	if err != nil {
		return nil, false, err
	}

	// Step 3: Get or create the cache for this team
	teamCache := ByTeamCache.GetOrCreateCache(teamID)

	// Step 4: Load all roles for this team if not already loaded
	return teamCache.GetOrLoad(id, func() (map[string]*AppResourceRoleModel, error) {
		return fetchAllAppResourceRolesForTeam(ctx, client, teamID)
	})
}

// GetCachedForImport retrieves an app resource role from cache when we only have the ID.
// Used by ImportState where we don't know the resource_id upfront.
// On first import for a team, does a GET to learn the resource_id, then loads the team cache.
// Uses ImportCache to prevent thundering herd - only one goroutine fetches the initial resource.
func GetCachedForImport(ctx context.Context, client *serval.Client, id string) (*AppResourceRoleModel, bool, error) {
	if ByTeamCache == nil || importLoadLock == nil {
		tflog.Debug(ctx, "[AppResourceRole] Cache not initialized, falling back to API", map[string]interface{}{"id": id})
		return nil, false, nil // Cache not initialized, caller should fall back to API
	}

	// Step 1: Check if this role is already in any loaded team cache
	if item, teamID := ByTeamCache.FindInLoadedCaches(id); item != nil {
		tflog.Debug(ctx, "[AppResourceRole] CACHE HIT - found in loaded cache", map[string]interface{}{"id": id, "teamID": teamID})
		return item, true, nil
	}

	// Step 2: Check if import lock already has the team_id (another goroutine discovered it)
	if importLoadLock.IsInitialized() {
		teamID := importLoadLock.GetParentKey()
		if teamID != "" {
			tflog.Debug(ctx, "[AppResourceRole] Lock initialized, using discovered team_id", map[string]interface{}{"id": id, "teamID": teamID})
			teamCache := ByTeamCache.GetOrCreateCache(teamID)
			return teamCache.GetOrLoad(id, func() (map[string]*AppResourceRoleModel, error) {
				tflog.Info(ctx, "[AppResourceRole] LOADING ALL via List API", map[string]interface{}{"teamID": teamID})
				return fetchAllAppResourceRolesForTeam(ctx, client, teamID)
			})
		}
	}

	// Step 3: Try to acquire the load lock - only one goroutine should fetch
	if importLoadLock.AcquireLoadLock() {
		tflog.Info(ctx, "[AppResourceRole] ACQUIRED LOCK - will fetch initial resource", map[string]interface{}{"id": id})
		// This goroutine will do the initial fetch to discover team_id
		role, err := client.AppResourceRoles.Get(ctx, id,
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			tflog.Error(ctx, "[AppResourceRole] Initial Get failed", map[string]interface{}{"id": id, "error": err.Error()})
			importLoadLock.CompleteLoad("") // Signal failure
			return nil, false, err
		}

		// Get app_instance_id from resource_id, then team_id from app_instance_id
		appInstanceID, err := getAppInstanceIDForResource(ctx, client, role.ResourceID)
		if err != nil {
			tflog.Error(ctx, "[AppResourceRole] Failed to get app_instance_id", map[string]interface{}{"id": id, "error": err.Error()})
			importLoadLock.CompleteLoad("") // Signal failure
			return nil, false, err
		}

		teamID, err := getTeamIDForAppInstance(ctx, client, appInstanceID)
		if err != nil {
			tflog.Error(ctx, "[AppResourceRole] Failed to get team_id", map[string]interface{}{"id": id, "error": err.Error()})
			importLoadLock.CompleteLoad("") // Signal failure
			return nil, false, err
		}

		tflog.Info(ctx, "[AppResourceRole] Discovered team_id, completing lock", map[string]interface{}{"id": id, "teamID": teamID})
		// Complete the lock so other goroutines can proceed
		importLoadLock.CompleteLoad(teamID)

		// Now load the full cache and return the item
		return GetCached(ctx, client, id, role.ResourceID)
	}

	// Step 4: Another goroutine is loading - wait for it
	tflog.Debug(ctx, "[AppResourceRole] WAITING for lock holder to complete", map[string]interface{}{"id": id})
	teamID := importLoadLock.WaitForLoad()
	if teamID == "" {
		tflog.Warn(ctx, "[AppResourceRole] Lock holder failed, falling back to individual Get", map[string]interface{}{"id": id})
		// Loading failed, fall back to individual API call
		role, err := client.AppResourceRoles.Get(ctx, id,
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			return nil, false, err
		}
		return GetCached(ctx, client, id, role.ResourceID)
	}

	tflog.Debug(ctx, "[AppResourceRole] Lock holder completed, using cache", map[string]interface{}{"id": id, "teamID": teamID})
	// Use the discovered team_id to get from cache
	teamCache := ByTeamCache.GetOrCreateCache(teamID)
	return teamCache.GetOrLoad(id, func() (map[string]*AppResourceRoleModel, error) {
		tflog.Info(ctx, "[AppResourceRole] LOADING ALL via List API (from waiter)", map[string]interface{}{"teamID": teamID})
		return fetchAllAppResourceRolesForTeam(ctx, client, teamID)
	})
}

// getAppInstanceIDForResource looks up the app_instance_id for a resource_id.
// Uses cached mapping if available, otherwise fetches from API.
func getAppInstanceIDForResource(ctx context.Context, client *serval.Client, resourceID string) (string, error) {
	// Check mapping cache first
	if appInstanceID, found := ResourceToAppInstance.Get(resourceID); found {
		return appInstanceID, nil
	}

	// Fetch resource to get app_instance_id
	resource, err := client.AppResources.Get(ctx, resourceID,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return "", err
	}

	// Cache the mapping for future use
	ResourceToAppInstance.Set(resourceID, resource.AppInstanceID)
	return resource.AppInstanceID, nil
}

// getTeamIDForAppInstance looks up the team_id for an app_instance_id.
// Uses the shared mapping cache from app_resource package.
func getTeamIDForAppInstance(ctx context.Context, client *serval.Client, appInstanceID string) (string, error) {
	// Check the shared mapping cache first
	if app_resource.AppInstanceToTeam != nil {
		if teamID, found := app_resource.AppInstanceToTeam.Get(appInstanceID); found {
			return teamID, nil
		}
	}

	// Fetch app instance to get team_id
	appInstance, err := client.AppInstances.Get(ctx, appInstanceID,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return "", err
	}

	// Cache the mapping in the shared cache for future use
	if app_resource.AppInstanceToTeam != nil {
		app_resource.AppInstanceToTeam.Set(appInstanceID, appInstance.TeamID)
	}
	return appInstance.TeamID, nil
}

// fetchAllAppResourceRolesForTeam retrieves all app resource roles for a team via the List API.
func fetchAllAppResourceRolesForTeam(ctx context.Context, client *serval.Client, teamID string) (map[string]*AppResourceRoleModel, error) {
	result := make(map[string]*AppResourceRoleModel)
	var pageToken *string

	for {
		params := serval.AppResourceRoleListParams{
			TeamID:   serval.String(teamID),
			PageSize: serval.Int(5000),
		}
		if pageToken != nil {
			params.PageToken = serval.String(*pageToken)
		}

		res := new(http.Response)
		_, err := client.AppResourceRoles.List(ctx, params,
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
			Data          []AppResourceRoleModel `json:"data"`
			NextPageToken *string                `json:"nextPageToken,omitempty"`
		}
		if err := apijson.UnmarshalComputed(bytes, &page); err != nil {
			return nil, err
		}

		for i := range page.Data {
			item := page.Data[i]
			result[item.ID.ValueString()] = &item
		}

		if page.NextPageToken == nil || *page.NextPageToken == "" {
			break
		}
		pageToken = page.NextPageToken
	}

	return result, nil
}
