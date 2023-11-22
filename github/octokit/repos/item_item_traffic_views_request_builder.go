package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic966bd8fcb01cfb93ece5d243c1cba5fc18d3a5aabcb8395774cc5bafa3508f4 "github.com/octokit/go-sdk/github/octokit/repos/item/item/traffic/views"
)

// ItemItemTrafficViewsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic\views
type ItemItemTrafficViewsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTrafficViewsRequestBuilderGetQueryParameters get the total number of views and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
type ItemItemTrafficViewsRequestBuilderGetQueryParameters struct {
    // The time frame to display results for.
    // Deprecated: This property is deprecated, use perAsGetPerQueryParameterType instead
    Per *string `uriparametername:"per"`
    // The time frame to display results for.
    PerAsGetPerQueryParameterType *ic966bd8fcb01cfb93ece5d243c1cba5fc18d3a5aabcb8395774cc5bafa3508f4.GetPerQueryParameterType `uriparametername:"per"`
}
// ItemItemTrafficViewsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTrafficViewsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemTrafficViewsRequestBuilderGetQueryParameters
}
// NewItemItemTrafficViewsRequestBuilderInternal instantiates a new ViewsRequestBuilder and sets the default values.
func NewItemItemTrafficViewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficViewsRequestBuilder) {
    m := &ItemItemTrafficViewsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/traffic/views{?per*}", pathParameters),
    }
    return m
}
// NewItemItemTrafficViewsRequestBuilder instantiates a new ViewsRequestBuilder and sets the default values.
func NewItemItemTrafficViewsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficViewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficViewsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the total number of views and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/traffic#get-page-views
func (m *ItemItemTrafficViewsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTrafficViewsRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ViewTrafficable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateViewTrafficFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ViewTrafficable), nil
}
// ToGetRequestInformation get the total number of views and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
func (m *ItemItemTrafficViewsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTrafficViewsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemTrafficViewsRequestBuilder) WithUrl(rawUrl string)(*ItemItemTrafficViewsRequestBuilder) {
    return NewItemItemTrafficViewsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
