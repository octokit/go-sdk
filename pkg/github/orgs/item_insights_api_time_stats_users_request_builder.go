package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiTimeStatsUsersRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\time-stats\users
type ItemInsightsApiTimeStatsUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByUser_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.timeStats.users.item collection
// returns a *ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsUsersRequestBuilder) ByUser_id(user_id string)(*ItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if user_id != "" {
        urlTplParams["user_id"] = user_id
    }
    return NewItemInsightsApiTimeStatsUsersWithUser_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiTimeStatsUsersRequestBuilderInternal instantiates a new ItemInsightsApiTimeStatsUsersRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsUsersRequestBuilder) {
    m := &ItemInsightsApiTimeStatsUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/time-stats/users", pathParameters),
    }
    return m
}
// NewItemInsightsApiTimeStatsUsersRequestBuilder instantiates a new ItemInsightsApiTimeStatsUsersRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiTimeStatsUsersRequestBuilderInternal(urlParams, requestAdapter)
}
