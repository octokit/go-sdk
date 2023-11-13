package users

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id5afee440d3b2c0d817934948e5c81cd01d6117e1c97204dd063f65c061893d4 "github.com/octokit/go-sdk/users/item/repos"
)

// ItemReposRequestBuilder builds and executes requests for operations under \users\{username}\repos
type ItemReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReposRequestBuilderGetQueryParameters lists public repositories for the specified user. Note: For GitHub AE, this endpoint will list internal repositories for the specified user.
type ItemReposRequestBuilderGetQueryParameters struct {
    // The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`.
    DirectionAsGetDirectionQueryParameterType *id5afee440d3b2c0d817934948e5c81cd01d6117e1c97204dd063f65c061893d4.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *id5afee440d3b2c0d817934948e5c81cd01d6117e1c97204dd063f65c061893d4.GetSortQueryParameterType `uriparametername:"sort"`
    // Limit results to repositories of the specified type.
    // Deprecated: This property is deprecated, use typeAsGetTypeQueryParameterType instead
    Type *string `uriparametername:"type"`
    // Limit results to repositories of the specified type.
    TypeAsGetTypeQueryParameterType *id5afee440d3b2c0d817934948e5c81cd01d6117e1c97204dd063f65c061893d4.GetTypeQueryParameterType `uriparametername:"type"`
}
// ItemReposRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReposRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemReposRequestBuilderGetQueryParameters
}
// NewItemReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposRequestBuilder) {
    m := &ItemReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/repos{?type*,sort*,direction*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists public repositories for the specified user. Note: For GitHub AE, this endpoint will list internal repositories for the specified user.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#list-repositories-for-a-user
func (m *ItemReposRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemReposRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMinimalRepositoryFromDiscriminatorValue, nil)
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
// ToGetRequestInformation lists public repositories for the specified user. Note: For GitHub AE, this endpoint will list internal repositories for the specified user.
func (m *ItemReposRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemReposRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemReposRequestBuilder) WithUrl(rawUrl string)(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
