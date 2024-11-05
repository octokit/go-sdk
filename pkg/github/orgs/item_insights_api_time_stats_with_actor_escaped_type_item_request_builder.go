package orgs

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\time-stats\{actor_type}
type ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByActor_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.timeStats.item.item collection
// returns a *ItemInsightsApiTimeStatsItemWithActor_ItemRequestBuilder when successful
func (m *ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder) ByActor_id(actor_id int32)(*ItemInsightsApiTimeStatsItemWithActor_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["actor_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(actor_id), 10)
    return NewItemInsightsApiTimeStatsItemWithActor_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilderInternal instantiates a new ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder) {
    m := &ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/time-stats/{actor_type}", pathParameters),
    }
    return m
}
// NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder instantiates a new ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiTimeStatsWithActor_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiTimeStatsWithActor_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
