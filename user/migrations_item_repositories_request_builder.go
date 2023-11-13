package user

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MigrationsItemRepositoriesRequestBuilder builds and executes requests for operations under \user\migrations\{migration_id}\repositories
type MigrationsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MigrationsItemRepositoriesRequestBuilderGetQueryParameters lists all the repositories for this user migration.
type MigrationsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// MigrationsItemRepositoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MigrationsItemRepositoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MigrationsItemRepositoriesRequestBuilderGetQueryParameters
}
// NewMigrationsItemRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewMigrationsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemRepositoriesRequestBuilder) {
    m := &MigrationsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/migrations/{migration_id}/repositories{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewMigrationsItemRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewMigrationsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MigrationsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMigrationsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all the repositories for this user migration.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/users#list-repositories-for-a-user-migration
func (m *MigrationsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *MigrationsItemRepositoriesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMinimalRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all the repositories for this user migration.
func (m *MigrationsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MigrationsItemRepositoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MigrationsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*MigrationsItemRepositoriesRequestBuilder) {
    return NewMigrationsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
