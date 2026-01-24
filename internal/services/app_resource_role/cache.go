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
)

// ByTeamCache stores AppResourceRoles keyed by team_id for efficient per-team caching.
var ByTeamCache *cache.KeyedCache[AppResourceRoleModel]

// ResourceToAppInstance maps resource_id → app_instance_id.
var ResourceToAppInstance *cache.MappingCache

// InitCache initializes the app resource role cache. Call this from provider.Configure().
func InitCache() {
	ByTeamCache = cache.NewKeyedCache[AppResourceRoleModel]()
	ResourceToAppInstance = cache.NewMappingCache()
}

// GetCached retrieves an app resource role from the per-team cache, loading via List API if needed.
// Returns (model, found, error). If the cache fails to load, error is non-nil.
// If the role doesn't exist, found is false.
func GetCached(ctx context.Context, client *serval.Client, id string, resourceID string) (*AppResourceRoleModel, bool, error) {
	if ByTeamCache == nil {
		return nil, false, nil // Cache not initialized, caller should fall back to API
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
