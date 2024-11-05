package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiUserStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\user-stats
type ItemInsightsApiUserStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUser_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.userStats.item collection
// returns a *ItemInsightsApiUserStatsWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiUserStatsRequestBuilder) ByUser_id(user_id string)(*ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if user_id != "" {
        urlTplParams["user_id"] = user_id
    }
    return NewItemInsightsApiUserStatsWithUser_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiUserStatsRequestBuilderInternal instantiates a new ItemInsightsApiUserStatsRequestBuilder and sets the default values.
func NewItemInsightsApiUserStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiUserStatsRequestBuilder) {
    m := &ItemInsightsApiUserStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/user-stats", pathParameters),
    }
    return m
}
// NewItemInsightsApiUserStatsRequestBuilder instantiates a new ItemInsightsApiUserStatsRequestBuilder and sets the default values.
func NewItemInsightsApiUserStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiUserStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiUserStatsRequestBuilderInternal(urlParams, requestAdapter)
}
