package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i4239b9a99f590fb557f6cee5c0b7a464cc1a12da753e0582e737b6a06899504e "github.com/octokit/go-sdk/pkg/github/repos/item/item/dependabot/alerts"
)

// ItemItemDependabotAlertsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\dependabot\alerts
type ItemItemDependabotAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependabotAlertsRequestBuilderGetQueryParameters you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
type ItemItemDependabotAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i4239b9a99f590fb557f6cee5c0b7a464cc1a12da753e0582e737b6a06899504e.GetDirectionQueryParameterType `uriparametername:"direction"`
    // A comma-separated list of ecosystems. If specified, only alerts for these ecosystems will be returned.Can be: `composer`, `go`, `maven`, `npm`, `nuget`, `pip`, `pub`, `rubygems`, `rust`
    Ecosystem *string `uriparametername:"ecosystem"`
    // **Deprecated**. The number of results per page (max 100), starting from the first matching result.This parameter must not be used in combination with `last`.Instead, use `per_page` in combination with `after` to fetch the first page of results.
    First *int32 `uriparametername:"first"`
    // **Deprecated**. The number of results per page (max 100), starting from the last matching result.This parameter must not be used in combination with `first`.Instead, use `per_page` in combination with `before` to fetch the last page of results.
    Last *int32 `uriparametername:"last"`
    // A comma-separated list of full manifest paths. If specified, only alerts for these manifests will be returned.
    Manifest *string `uriparametername:"manifest"`
    // A comma-separated list of package names. If specified, only alerts for these packages will be returned.
    Package *string `uriparametername:"package"`
    // **Deprecated**. Page number of the results to fetch. Use cursor-based pagination with `before` or `after` instead.
    // Deprecated: 
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    // Deprecated: 
    Per_page *int32 `uriparametername:"per_page"`
    // The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.
    // Deprecated: This property is deprecated, use scopeAsGetScopeQueryParameterType instead
    Scope *string `uriparametername:"scope"`
    // The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.
    ScopeAsGetScopeQueryParameterType *i4239b9a99f590fb557f6cee5c0b7a464cc1a12da753e0582e737b6a06899504e.GetScopeQueryParameterType `uriparametername:"scope"`
    // A comma-separated list of severities. If specified, only alerts with these severities will be returned.Can be: `low`, `medium`, `high`, `critical`
    Severity *string `uriparametername:"severity"`
    // The property by which to sort the results.`created` means when the alert was created.`updated` means when the alert's state last changed.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property by which to sort the results.`created` means when the alert was created.`updated` means when the alert's state last changed.
    SortAsGetSortQueryParameterType *i4239b9a99f590fb557f6cee5c0b7a464cc1a12da753e0582e737b6a06899504e.GetSortQueryParameterType `uriparametername:"sort"`
    // A comma-separated list of states. If specified, only alerts with these states will be returned.Can be: `auto_dismissed`, `dismissed`, `fixed`, `open`
    State *string `uriparametername:"state"`
}
// ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemDependabotAlertsRequestBuilderGetQueryParameters
}
// ByAlert_number gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.dependabot.alerts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemDependabotAlertsRequestBuilder) ByAlert_number(alert_number string)(*ItemItemDependabotAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alert_number != "" {
        urlTplParams["alert_number"] = alert_number
    }
    return NewItemItemDependabotAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAlert_numberInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.repos.item.item.dependabot.alerts.item collection
func (m *ItemItemDependabotAlertsRequestBuilder) ByAlert_numberInteger(alert_number int32)(*ItemItemDependabotAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemDependabotAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDependabotAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    m := &ItemItemDependabotAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/dependabot/alerts{?state*,severity*,ecosystem*,package*,manifest*,scope*,sort*,direction*,page*,per_page*,before*,after*,first*,last*}", pathParameters),
    }
    return m
}
// NewItemItemDependabotAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependabotAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/alerts#list-dependabot-alerts-for-a-repository
func (m *ItemItemDependabotAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DependabotAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDependabotAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DependabotAlertable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DependabotAlertable)
        }
    }
    return val, nil
}
// ToGetRequestInformation you must use an access token with the `security_events` scope to use this endpoint with private repositories.You can also use tokens with the `public_repo` scope for public repositories only.GitHub Apps must have **Dependabot alerts** read permission to use this endpoint.
func (m *ItemItemDependabotAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDependabotAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemDependabotAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependabotAlertsRequestBuilder) {
    return NewItemItemDependabotAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
