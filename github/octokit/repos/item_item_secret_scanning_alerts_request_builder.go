package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib9b85e8b6fc335d534455584e6c960195e07c1a8ac8ba2b04321dd3402c1f98a "github.com/octokit/go-sdk/github/octokit/repos/item/item/secretscanning/alerts"
)

// ItemItemSecretScanningAlertsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\secret-scanning\alerts
type ItemItemSecretScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters lists secret scanning alerts for an eligible repository, from newest to oldest.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
type ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty "after" query string.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty "before" query string.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *ib9b85e8b6fc335d534455584e6c960195e07c1a8ac8ba2b04321dd3402c1f98a.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`.
    Resolution *string `uriparametername:"resolution"`
    // A comma-separated list of secret types to return. By default all secret types are returned.See "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)"for a complete list of secret types.
    Secret_type *string `uriparametername:"secret_type"`
    // The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved.
    SortAsGetSortQueryParameterType *ib9b85e8b6fc335d534455584e6c960195e07c1a8ac8ba2b04321dd3402c1f98a.GetSortQueryParameterType `uriparametername:"sort"`
    // Set to `open` or `resolved` to only list secret scanning alerts in a specific state.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Set to `open` or `resolved` to only list secret scanning alerts in a specific state.
    StateAsGetStateQueryParameterType *ib9b85e8b6fc335d534455584e6c960195e07c1a8ac8ba2b04321dd3402c1f98a.GetStateQueryParameterType `uriparametername:"state"`
    // A comma-separated list of validities that, when present, will return alerts that match the validities in this list. Valid options are `active`, `inactive`, and `unknown`.
    Validity *string `uriparametername:"validity"`
}
// ItemItemSecretScanningAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecretScanningAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemSecretScanningAlertsRequestBuilderGetQueryParameters
}
// ByAlert_number gets an item from the octokit.repos.item.item.secretScanning.alerts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemSecretScanningAlertsRequestBuilder) ByAlert_number(alert_number string)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alert_number != "" {
        urlTplParams["alert_number"] = alert_number
    }
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAlert_numberInteger gets an item from the octokit.repos.item.item.secretScanning.alerts.item collection
func (m *ItemItemSecretScanningAlertsRequestBuilder) ByAlert_numberInteger(alert_number int32)(*ItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemSecretScanningAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemSecretScanningAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsRequestBuilder) {
    m := &ItemItemSecretScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/secret-scanning/alerts{?state*,secret_type*,resolution*,sort*,direction*,page*,per_page*,before*,after*,validity*}", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists secret scanning alerts for an eligible repository, from newest to oldest.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#list-secret-scanning-alerts-for-a-repository
func (m *ItemItemSecretScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemSecretScanningAlertsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSecretScanningAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningAlertable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningAlertable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists secret scanning alerts for an eligible repository, from newest to oldest.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
func (m *ItemItemSecretScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecretScanningAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemSecretScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecretScanningAlertsRequestBuilder) {
    return NewItemItemSecretScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
