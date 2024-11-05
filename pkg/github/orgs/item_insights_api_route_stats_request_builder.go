package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInsightsApiRouteStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\route-stats
type ItemInsightsApiRouteStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByActor_type gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.insights.api.routeStats.item collection
// returns a *ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder when successful
func (m *ItemInsightsApiRouteStatsRequestBuilder) ByActor_type(actor_type string)(*ItemInsightsApiRouteStatsWithActor_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if actor_type != "" {
        urlTplParams["actor_type"] = actor_type
    }
    return NewItemInsightsApiRouteStatsWithActor_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiRouteStatsRequestBuilderInternal instantiates a new ItemInsightsApiRouteStatsRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsRequestBuilder) {
    m := &ItemInsightsApiRouteStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/route-stats", pathParameters),
    }
    return m
}
// NewItemInsightsApiRouteStatsRequestBuilder instantiates a new ItemInsightsApiRouteStatsRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiRouteStatsRequestBuilderInternal(urlParams, requestAdapter)
}
