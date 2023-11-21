package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTrafficPopularReferrersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic\popular\referrers
type ItemItemTrafficPopularReferrersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTrafficPopularReferrersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTrafficPopularReferrersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemTrafficPopularReferrersRequestBuilderInternal instantiates a new ReferrersRequestBuilder and sets the default values.
func NewItemItemTrafficPopularReferrersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularReferrersRequestBuilder) {
    m := &ItemItemTrafficPopularReferrersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/traffic/popular/referrers", pathParameters),
    }
    return m
}
// NewItemItemTrafficPopularReferrersRequestBuilder instantiates a new ReferrersRequestBuilder and sets the default values.
func NewItemItemTrafficPopularReferrersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficPopularReferrersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficPopularReferrersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the top 10 referrers over the last 14 days.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/traffic#get-top-referral-sources
func (m *ItemItemTrafficPopularReferrersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTrafficPopularReferrersRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ReferrerTrafficable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReferrerTrafficFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ReferrerTrafficable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ReferrerTrafficable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get the top 10 referrers over the last 14 days.
func (m *ItemItemTrafficPopularReferrersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTrafficPopularReferrersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *ItemItemTrafficPopularReferrersRequestBuilder) WithUrl(rawUrl string)(*ItemItemTrafficPopularReferrersRequestBuilder) {
    return NewItemItemTrafficPopularReferrersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
