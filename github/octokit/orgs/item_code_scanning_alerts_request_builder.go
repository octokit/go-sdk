package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    i8358ab267907a430f1e13a09cd3c77f84900604fa2581e56ec3f2a2b0471693d "github.com/octokit/go-sdk/github/octokit/orgs/item/codescanning/alerts"
)

// ItemCodeScanningAlertsRequestBuilder builds and executes requests for operations under \orgs\{org}\code-scanning\alerts
type ItemCodeScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCodeScanningAlertsRequestBuilderGetQueryParameters lists code scanning alerts for the default branch for all eligible repositories in an organization. Eligible repositories are repositories that are owned by organizations that you own or for which you are a security manager. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.
type ItemCodeScanningAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i8358ab267907a430f1e13a09cd3c77f84900604fa2581e56ec3f2a2b0471693d.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // If specified, only code scanning alerts with this severity will be returned.
    // Deprecated: This property is deprecated, use severityAsCodeScanningAlertSeverity instead
    Severity *string `uriparametername:"severity"`
    // If specified, only code scanning alerts with this severity will be returned.
    SeverityAsCodeScanningAlertSeverity *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningAlertSeverity `uriparametername:"severity"`
    // The property by which to sort the results.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property by which to sort the results.
    SortAsGetSortQueryParameterType *i8358ab267907a430f1e13a09cd3c77f84900604fa2581e56ec3f2a2b0471693d.GetSortQueryParameterType `uriparametername:"sort"`
    // If specified, only code scanning alerts with this state will be returned.
    // Deprecated: This property is deprecated, use stateAsCodeScanningAlertStateQuery instead
    State *string `uriparametername:"state"`
    // If specified, only code scanning alerts with this state will be returned.
    StateAsCodeScanningAlertStateQuery *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningAlertStateQuery `uriparametername:"state"`
    // The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.
    Tool_guid *string `uriparametername:"tool_guid"`
    // The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.
    Tool_name *string `uriparametername:"tool_name"`
}
// ItemCodeScanningAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodeScanningAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCodeScanningAlertsRequestBuilderGetQueryParameters
}
// NewItemCodeScanningAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemCodeScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningAlertsRequestBuilder) {
    m := &ItemCodeScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/code-scanning/alerts{?tool_name*,tool_guid*,before*,after*,page*,per_page*,direction*,state*,sort*,severity*}", pathParameters),
    }
    return m
}
// NewItemCodeScanningAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemCodeScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodeScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodeScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists code scanning alerts for the default branch for all eligible repositories in an organization. Eligible repositories are repositories that are owned by organizations that you own or for which you are a security manager. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#list-code-scanning-alerts-for-an-organization
func (m *ItemCodeScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCodeScanningAlertsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningOrganizationAlertItemsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "503": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateCodeScanningOrganizationAlertItemsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningOrganizationAlertItemsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CodeScanningOrganizationAlertItemsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists code scanning alerts for the default branch for all eligible repositories in an organization. Eligible repositories are repositories that are owned by organizations that you own or for which you are a security manager. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.
func (m *ItemCodeScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCodeScanningAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemCodeScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemCodeScanningAlertsRequestBuilder) {
    return NewItemCodeScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
