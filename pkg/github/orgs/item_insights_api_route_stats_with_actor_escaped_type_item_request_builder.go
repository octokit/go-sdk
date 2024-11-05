package orgs

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\route-stats\{actor_type}
type ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByActor_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.routeStats.item.item collection
// returns a *ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder when successful
func (m *ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder) ByActor_id(actor_id int32)(*ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["actor_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(actor_id), 10)
    return NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilderInternal instantiates a new ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder) {
    m := &ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/route-stats/{actor_type}", pathParameters),
    }
    return m
}
// NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder instantiates a new ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
