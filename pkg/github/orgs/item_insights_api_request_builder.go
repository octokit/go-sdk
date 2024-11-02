package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api
type ItemInsightsApiRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemInsightsApiRequestBuilderInternal instantiates a new ItemInsightsApiRequestBuilder and sets the default values.
func NewItemInsightsApiRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRequestBuilder) {
    m := &ItemInsightsApiRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api", pathParameters),
    }
    return m
}
// NewItemInsightsApiRequestBuilder instantiates a new ItemInsightsApiRequestBuilder and sets the default values.
func NewItemInsightsApiRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiRequestBuilderInternal(urlParams, requestAdapter)
}
// RouteStats the routeStats property
// returns a *ItemInsightsApiRouteStatsRequestBuilder when successful
func (m *ItemInsightsApiRequestBuilder) RouteStats()(*ItemInsightsApiRouteStatsRequestBuilder) {
    return NewItemInsightsApiRouteStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubjectStats the subjectStats property
// returns a *ItemInsightsApiSubjectStatsRequestBuilder when successful
func (m *ItemInsightsApiRequestBuilder) SubjectStats()(*ItemInsightsApiSubjectStatsRequestBuilder) {
    return NewItemInsightsApiSubjectStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SummaryStats the summaryStats property
// returns a *ItemInsightsApiSummaryStatsRequestBuilder when successful
func (m *ItemInsightsApiRequestBuilder) SummaryStats()(*ItemInsightsApiSummaryStatsRequestBuilder) {
    return NewItemInsightsApiSummaryStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeStats the timeStats property
// returns a *ItemInsightsApiTimeStatsRequestBuilder when successful
func (m *ItemInsightsApiRequestBuilder) TimeStats()(*ItemInsightsApiTimeStatsRequestBuilder) {
    return NewItemInsightsApiTimeStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStats the userStats property
// returns a *ItemInsightsApiUserStatsRequestBuilder when successful
func (m *ItemInsightsApiRequestBuilder) UserStats()(*ItemInsightsApiUserStatsRequestBuilder) {
    return NewItemInsightsApiUserStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
