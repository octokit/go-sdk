package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatsPunch_cardRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\stats\punch_card
type ItemItemStatsPunch_cardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatsPunch_cardRequestBuilderInternal instantiates a new Punch_cardRequestBuilder and sets the default values.
func NewItemItemStatsPunch_cardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsPunch_cardRequestBuilder) {
    m := &ItemItemStatsPunch_cardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/stats/punch_card", pathParameters),
    }
    return m
}
// NewItemItemStatsPunch_cardRequestBuilder instantiates a new Punch_cardRequestBuilder and sets the default values.
func NewItemItemStatsPunch_cardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatsPunch_cardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatsPunch_cardRequestBuilderInternal(urlParams, requestAdapter)
}
// Get each array contains the day number, hour number, and number of commits:*   `0-6`: Sunday - Saturday*   `0-23`: Hour of day*   Number of commitsFor example, `[2, 14, 25]` indicates that there were 25 total commits, during the 2:00pm hour on Tuesdays. All times are based on the time zone of individual commits.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/metrics/statistics#get-the-hourly-commit-count-for-each-day
func (m *ItemItemStatsPunch_cardRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration)([]int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "int32", nil)
    if err != nil {
        return nil, err
    }
    val := make([]int32, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*int32))
        }
    }
    return val, nil
}
// ToGetRequestInformation each array contains the day number, hour number, and number of commits:*   `0-6`: Sunday - Saturday*   `0-23`: Hour of day*   Number of commitsFor example, `[2, 14, 25]` indicates that there were 25 total commits, during the 2:00pm hour on Tuesdays. All times are based on the time zone of individual commits.
func (m *ItemItemStatsPunch_cardRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemStatsPunch_cardRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemStatsPunch_cardRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatsPunch_cardRequestBuilder) {
    return NewItemItemStatsPunch_cardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
