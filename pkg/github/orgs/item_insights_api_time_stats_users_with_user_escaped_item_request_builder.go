package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\time-stats\users\{user_id}
type ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderGetQueryParameters get the number of API requests and rate-limited requests made within an organization by a specific user over a specified time period.
type ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderGetQueryParameters struct {
    // The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The increment of time used to breakdown the query results (5m, 10m, 1h, etc.)
    Timestamp_increment *string `uriparametername:"timestamp_increment"`
}
// NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderInternal instantiates a new ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) {
    m := &ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/time-stats/users/{user_id}?min_timestamp={min_timestamp}&timestamp_increment={timestamp_increment}{&max_timestamp*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder instantiates a new ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of API requests and rate-limited requests made within an organization by a specific user over a specified time period.
// returns a []ApiInsightsTimeStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/api-insights#get-time-stats-by-user
func (m *ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsTimeStatsable, error) {
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
// ToGetRequestInformation get the number of API requests and rate-limited requests made within an organization by a specific user over a specified time period.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) {
    return NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
