package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    id2d3655a1028a42b85c7ad7aeb99baa507e8da22b67d4ad526b25a2aabb84ac3 "github.com/octokit/go-sdk/github/octokit/enterprises/item/secretscanning/alerts"
)

// ItemSecretScanningAlertsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\secret-scanning\alerts
type ItemSecretScanningAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecretScanningAlertsRequestBuilderGetQueryParameters lists secret scanning alerts for eligible repositories in an enterprise, from newest to oldest.To use this endpoint, you must be a member of the enterprise, and you must use an access token with the `repo` scope or `security_events` scope. Alerts are only returned for organizations in the enterprise for which you are an organization owner or a [security manager](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).
type ItemSecretScanningAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *id2d3655a1028a42b85c7ad7aeb99baa507e8da22b67d4ad526b25a2aabb84ac3.GetDirectionQueryParameterType `uriparametername:"direction"`
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
    SortAsGetSortQueryParameterType *id2d3655a1028a42b85c7ad7aeb99baa507e8da22b67d4ad526b25a2aabb84ac3.GetSortQueryParameterType `uriparametername:"sort"`
    // Set to `open` or `resolved` to only list secret scanning alerts in a specific state.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Set to `open` or `resolved` to only list secret scanning alerts in a specific state.
    StateAsGetStateQueryParameterType *id2d3655a1028a42b85c7ad7aeb99baa507e8da22b67d4ad526b25a2aabb84ac3.GetStateQueryParameterType `uriparametername:"state"`
    // A comma-separated list of validities that, when present, will return alerts that match the validities in this list. Valid options are `active`, `inactive`, and `unknown`.
    Validity *string `uriparametername:"validity"`
}
// ItemSecretScanningAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecretScanningAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSecretScanningAlertsRequestBuilderGetQueryParameters
}
// NewItemSecretScanningAlertsRequestBuilderInternal instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemSecretScanningAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningAlertsRequestBuilder) {
    m := &ItemSecretScanningAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/secret-scanning/alerts{?state*,secret_type*,resolution*,sort*,direction*,per_page*,before*,after*,validity*}", pathParameters),
    }
    return m
}
// NewItemSecretScanningAlertsRequestBuilder instantiates a new AlertsRequestBuilder and sets the default values.
func NewItemSecretScanningAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecretScanningAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecretScanningAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists secret scanning alerts for eligible repositories in an enterprise, from newest to oldest.To use this endpoint, you must be a member of the enterprise, and you must use an access token with the `repo` scope or `security_events` scope. Alerts are only returned for organizations in the enterprise for which you are an organization owner or a [security manager](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#list-secret-scanning-alerts-for-an-enterprise
func (m *ItemSecretScanningAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecretScanningAlertsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.OrganizationSecretScanningAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "503": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateAlerts503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateOrganizationSecretScanningAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.OrganizationSecretScanningAlertable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.OrganizationSecretScanningAlertable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists secret scanning alerts for eligible repositories in an enterprise, from newest to oldest.To use this endpoint, you must be a member of the enterprise, and you must use an access token with the `repo` scope or `security_events` scope. Alerts are only returned for organizations in the enterprise for which you are an organization owner or a [security manager](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).
func (m *ItemSecretScanningAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecretScanningAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSecretScanningAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemSecretScanningAlertsRequestBuilder) {
    return NewItemSecretScanningAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
