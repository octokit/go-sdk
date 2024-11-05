package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemInsightsApiTimeStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\time-stats
type ItemInsightsApiTimeStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiTimeStatsRequestBuilderGetQueryParameters get the number of API requests and rate-limited requests made within an organization over a specified time period.
type ItemInsightsApiTimeStatsRequestBuilderGetQueryParameters struct {
    // The maximum timestamp to query for stats
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The increment of time used to breakdown the query results (5m, 10m, 1h, etc.)
    Timestamp_increment *string `uriparametername:"timestamp_increment"`
}
// ByActor_type gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.timeStats.item collection
// returns a *ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsRequestBuilder) ByActor_type(actor_type string)(*ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if actor_type != "" {
        urlTplParams["actor_type"] = actor_type
    }
    return NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiTimeStatsRequestBuilderInternal instantiates a new ItemInsightsApiTimeStatsRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsRequestBuilder) {
    m := &ItemInsightsApiTimeStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/time-stats?max_timestamp={max_timestamp}&min_timestamp={min_timestamp}&timestamp_increment={timestamp_increment}", pathParameters),
    }
    return m
}
// NewItemInsightsApiTimeStatsRequestBuilder instantiates a new ItemInsightsApiTimeStatsRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiTimeStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of API requests and rate-limited requests made within an organization over a specified time period.
// returns a []ApiInsightsTimeStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/api-insights#get-time-stats
func (m *ItemInsightsApiTimeStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiTimeStatsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsTimeStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateApiInsightsTimeStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsTimeStatsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsTimeStatsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get the number of API requests and rate-limited requests made within an organization over a specified time period.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiTimeStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiTimeStatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// Users the users property
// returns a *ItemInsightsApiTimeStatsUsersRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsRequestBuilder) Users()(*ItemInsightsApiTimeStatsUsersRequestBuilder) {
    return NewItemInsightsApiTimeStatsUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiTimeStatsRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiTimeStatsRequestBuilder) {
    return NewItemInsightsApiTimeStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
