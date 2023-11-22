package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}
type ItemMigrationsWithMigration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
type ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // Exclude attributes from the API response to improve performance
    Exclude []string `uriparametername:"exclude"`
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters
}
// Archive the archive property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Archive()(*ItemMigrationsItemArchiveRequestBuilder) {
    return NewItemMigrationsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    m := &ItemMigrationsWithMigration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}{?exclude*}", pathParameters),
    }
    return m
}
// NewItemMigrationsWithMigration_ItemRequestBuilder instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/orgs#get-an-organization-migration-status
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateMigrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Migrationable), nil
}
// Repos the repos property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repos()(*ItemMigrationsItemReposRequestBuilder) {
    return NewItemMigrationsItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repositories()(*ItemMigrationsItemRepositoriesRequestBuilder) {
    return NewItemMigrationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    return NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
