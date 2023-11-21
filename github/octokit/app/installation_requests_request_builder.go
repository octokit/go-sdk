package app

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstallationRequestsRequestBuilder builds and executes requests for operations under \app\installation-requests
type InstallationRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InstallationRequestsRequestBuilderGetQueryParameters lists all the pending installation requests for the authenticated GitHub App.
type InstallationRequestsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// InstallationRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstallationRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstallationRequestsRequestBuilderGetQueryParameters
}
// NewInstallationRequestsRequestBuilderInternal instantiates a new InstallationRequestsRequestBuilder and sets the default values.
func NewInstallationRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestsRequestBuilder) {
    m := &InstallationRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/installation-requests{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewInstallationRequestsRequestBuilder instantiates a new InstallationRequestsRequestBuilder and sets the default values.
func NewInstallationRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstallationRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstallationRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all the pending installation requests for the authenticated GitHub App.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/apps#list-installation-requests-for-the-authenticated-app
func (m *InstallationRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *InstallationRequestsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IntegrationInstallationRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateIntegrationInstallationRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IntegrationInstallationRequestable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.IntegrationInstallationRequestable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all the pending installation requests for the authenticated GitHub App.
func (m *InstallationRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InstallationRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *InstallationRequestsRequestBuilder) WithUrl(rawUrl string)(*InstallationRequestsRequestBuilder) {
    return NewInstallationRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
