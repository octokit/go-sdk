package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemSecretScanningAlertsItemLocationsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\secret-scanning\alerts\{alert_number}\locations
type ItemItemSecretScanningAlertsItemLocationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetQueryParameters lists all locations for a given secret scanning alert for an eligible repository.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
type ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetQueryParameters
}
// NewItemItemSecretScanningAlertsItemLocationsRequestBuilderInternal instantiates a new LocationsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsItemLocationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsItemLocationsRequestBuilder) {
    m := &ItemItemSecretScanningAlertsItemLocationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}/locations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningAlertsItemLocationsRequestBuilder instantiates a new LocationsRequestBuilder and sets the default values.
func NewItemItemSecretScanningAlertsItemLocationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningAlertsItemLocationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningAlertsItemLocationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all locations for a given secret scanning alert for an eligible repository.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#list-locations-for-a-secret-scanning-alert
func (m *ItemItemSecretScanningAlertsItemLocationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningLocationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateLocations503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSecretScanningLocationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningLocationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SecretScanningLocationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all locations for a given secret scanning alert for an eligible repository.To use this endpoint, you must be an administrator for the repository or for the organization that owns the repository, and you must use a personal access token with the `repo` scope or `security_events` scope.For public repositories, you may instead use the `public_repo` scope.GitHub Apps must have the `secret_scanning_alerts` read permission to use this endpoint.
func (m *ItemItemSecretScanningAlertsItemLocationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecretScanningAlertsItemLocationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemSecretScanningAlertsItemLocationsRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecretScanningAlertsItemLocationsRequestBuilder) {
    return NewItemItemSecretScanningAlertsItemLocationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
