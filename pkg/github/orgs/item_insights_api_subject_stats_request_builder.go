package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i66b1523c4c4ef510ea098d616fe90f80d2679e0386fabfe6a2f1742954997624 "github.com/octokit/go-sdk/pkg/github/orgs/item/insights/api/subjectstats"
)

// ItemInsightsApiSubjectStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\subject-stats
type ItemInsightsApiSubjectStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
type ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i66b1523c4c4ef510ea098d616fe90f80d2679e0386fabfe6a2f1742954997624.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    Sort []i66b1523c4c4ef510ea098d616fe90f80d2679e0386fabfe6a2f1742954997624.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewItemInsightsApiSubjectStatsRequestBuilderInternal instantiates a new ItemInsightsApiSubjectStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSubjectStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    m := &ItemInsightsApiSubjectStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/subject-stats?min_timestamp={min_timestamp}{&direction*,max_timestamp*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiSubjectStatsRequestBuilder instantiates a new ItemInsightsApiSubjectStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSubjectStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSubjectStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
// returns a []ApiInsightsSubjectStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/api-insights#get-subject-stats
func (m *ItemInsightsApiSubjectStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsSubjectStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateApiInsightsSubjectStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsSubjectStatsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ApiInsightsSubjectStatsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get API request statistics for all subjects within an organization within a specified time frame. Subjects can be users or GitHub Apps.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiSubjectStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSubjectStatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiSubjectStatsRequestBuilder when successful
func (m *ItemInsightsApiSubjectStatsRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiSubjectStatsRequestBuilder) {
    return NewItemInsightsApiSubjectStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
