package orgs

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\summary-stats\{actor_type}
type ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByActor_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.summaryStats.item.item collection
// returns a *ItemInsightsApiSummaryStatsItemWithActor_ItemRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder) ByActor_id(actor_id int32)(*ItemInsightsApiSummaryStatsItemWithActor_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["actor_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(actor_id), 10)
    return NewItemInsightsApiSummaryStatsItemWithActor_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilderInternal instantiates a new ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder) {
    m := &ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/summary-stats/{actor_type}", pathParameters),
    }
    return m
}
// NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder instantiates a new ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
