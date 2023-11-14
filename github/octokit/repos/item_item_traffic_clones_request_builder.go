package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0902f0564408944fb964ef491d8a10cc920645bb50bffe9ebf870f6604602113 "github.com/octokit/go-sdk/github/octokit/repos/item/item/traffic/clones"
)

// ItemItemTrafficClonesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\traffic\clones
type ItemItemTrafficClonesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTrafficClonesRequestBuilderGetQueryParameters get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
type ItemItemTrafficClonesRequestBuilderGetQueryParameters struct {
    // The time frame to display results for.
    // Deprecated: This property is deprecated, use perAsGetPerQueryParameterType instead
    Per *string `uriparametername:"per"`
    // The time frame to display results for.
    PerAsGetPerQueryParameterType *i0902f0564408944fb964ef491d8a10cc920645bb50bffe9ebf870f6604602113.GetPerQueryParameterType `uriparametername:"per"`
}
// ItemItemTrafficClonesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTrafficClonesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemTrafficClonesRequestBuilderGetQueryParameters
}
// NewItemItemTrafficClonesRequestBuilderInternal instantiates a new ClonesRequestBuilder and sets the default values.
func NewItemItemTrafficClonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficClonesRequestBuilder) {
    m := &ItemItemTrafficClonesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/traffic/clones{?per*}", pathParameters),
    }
    return m
}
// NewItemItemTrafficClonesRequestBuilder instantiates a new ClonesRequestBuilder and sets the default values.
func NewItemItemTrafficClonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTrafficClonesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTrafficClonesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/traffic#get-repository-clones
func (m *ItemItemTrafficClonesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTrafficClonesRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CloneTrafficable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCloneTrafficFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CloneTrafficable), nil
}
// ToGetRequestInformation get the total number of clones and breakdown per day or week for the last 14 days. Timestamps are aligned to UTC midnight of the beginning of the day or week. Week begins on Monday.
func (m *ItemItemTrafficClonesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTrafficClonesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemTrafficClonesRequestBuilder) WithUrl(rawUrl string)(*ItemItemTrafficClonesRequestBuilder) {
    return NewItemItemTrafficClonesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
