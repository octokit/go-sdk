package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiSummaryStatsUsersRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\summary-stats\users
type ItemInsightsApiSummaryStatsUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUser_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.summaryStats.users.item collection
// returns a *ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsUsersRequestBuilder) ByUser_id(user_id string)(*ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if user_id != "" {
        urlTplParams["user_id"] = user_id
    }
    return NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiSummaryStatsUsersRequestBuilderInternal instantiates a new ItemInsightsApiSummaryStatsUsersRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsUsersRequestBuilder) {
    m := &ItemInsightsApiSummaryStatsUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/summary-stats/users", pathParameters),
    }
    return m
}
// NewItemInsightsApiSummaryStatsUsersRequestBuilder instantiates a new ItemInsightsApiSummaryStatsUsersRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSummaryStatsUsersRequestBuilderInternal(urlParams, requestAdapter)
}
