package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}
type MigrationsWithMigration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/migrations/users#download-a-user-migration-archive).
type MigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // 
    Exclude []string `uriparametername:"exclude"`
}
// MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MigrationsWithMigration_ItemRequestBuilderGetQueryParameters
}
// Archive the archive property
func (m *MigrationsWithMigration_ItemRequestBuilder) Archive()(*MigrationsItemArchiveRequestBuilder) {
    return NewMigrationsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    m := &MigrationsWithMigration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/migrations/{migration_id}{?exclude*}", pathParameters),
    }
    return m
}
// NewMigrationsWithMigration_ItemRequestBuilder instantiates a new WithMigration_ItemRequestBuilder and sets the default values.
func NewMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/migrations/users#download-a-user-migration-archive).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/users#get-a-user-migration-status
func (m *MigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
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
func (m *MigrationsWithMigration_ItemRequestBuilder) Repos()(*MigrationsItemReposRequestBuilder) {
    return NewMigrationsItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
func (m *MigrationsWithMigration_ItemRequestBuilder) Repositories()(*MigrationsItemRepositoriesRequestBuilder) {
    return NewMigrationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation fetches a single user migration. The response includes the `state` of the migration, which can be one of the following values:*   `pending` - the migration hasn't started yet.*   `exporting` - the migration is in progress.*   `exported` - the migration finished successfully.*   `failed` - the migration failed.Once the migration has been `exported` you can [download the migration archive](https://docs.github.com/rest/migrations/users#download-a-user-migration-archive).
func (m *MigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MigrationsWithMigration_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MigrationsWithMigration_ItemRequestBuilder) WithUrl(rawUrl string)(*MigrationsWithMigration_ItemRequestBuilder) {
    return NewMigrationsWithMigration_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
