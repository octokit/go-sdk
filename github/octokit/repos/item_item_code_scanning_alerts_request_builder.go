package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i310c7fa33fa1c48c1b3e7a6088c5751691d82691dc25386b504b94581d877993 "github.com/octokit/go-sdk/github/octokit/repos/item/item/codescanning/alerts"
)

// ItemItemCodeScanningAlertsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\alerts
type ItemItemCodeScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters lists code scanning alerts.To use this endpoint, you must use an access token with the `security_events` scope or, for alerts from public repositories only, an access token with the `public_repo` scope.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).
type ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i310c7fa33fa1c48c1b3e7a6088c5751691d82691dc25386b504b94581d877993.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The Git reference for the results you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // If specified, only code scanning alerts with this severity will be returned.
    // Deprecated: This property is deprecated, use severityAsCodeScanningAlertSeverity instead
    Severity *string `uriparametername:"severity"`
    // If specified, only code scanning alerts with this severity will be returned.
    SeverityAsCodeScanningAlertSeverity *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningAlertSeverity `uriparametername:"severity"`
    // The property by which to sort the results.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property by which to sort the results.
    SortAsGetSortQueryParameterType *i310c7fa33fa1c48c1b3e7a6088c5751691d82691dc25386b504b94581d877993.GetSortQueryParameterType `uriparametername:"sort"`
    // If specified, only code scanning alerts with this state will be returned.
    // Deprecated: This property is deprecated, use stateAsCodeScanningAlertStateQuery instead
    State *string `uriparametername:"state"`
    // If specified, only code scanning alerts with this state will be returned.
    StateAsCodeScanningAlertStateQuery *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningAlertStateQuery `uriparametername:"state"`
    // The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.
    Tool_guid *string `uriparametername:"tool_guid"`
    // The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.
    Tool_name *string `uriparametername:"tool_name"`
}
// ItemItemCodeScanningAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodeScanningAlertsRequestBuilderGetQueryParameters
}
// ByAlert_number gets an item from the octokit.repos.item.item.codeScanning.alerts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemCodeScanningAlertsRequestBuilder) ByAlert_number(alert_number string)(*ItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alert_number != "" {
        urlTplParams["alert_number"] = alert_number
    }
    return NewItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAlert_numberInteger gets an item from the octokit.repos.item.item.codeScanning.alerts.item collection
func (m *ItemItemCodeScanningAlertsRequestBuilder) ByAlert_numberInteger(alert_number int32)(*ItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemCodeScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsRequestBuilder) {
    m := &ItemItemCodeScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/code-scanning/alerts{?tool_name*,tool_guid*,page*,per_page*,ref*,direction*,sort*,state*,severity*}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists code scanning alerts.To use this endpoint, you must use an access token with the `security_events` scope or, for alerts from public repositories only, an access token with the `public_repo` scope.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-code-scanning-alerts-for-a-repository
func (m *ItemItemCodeScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningAlertsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningAlertItemsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodeScanningAlertItemsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningAlertItemsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningAlertItemsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists code scanning alerts.To use this endpoint, you must use an access token with the `security_events` scope or, for alerts from public repositories only, an access token with the `public_repo` scope.The response includes a `most_recent_instance` object.This provides details of the most recent instance of this alertfor the default branch (or for the specified Git reference if you used `ref` in the request).
func (m *ItemItemCodeScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCodeScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAlertsRequestBuilder) {
    return NewItemItemCodeScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
