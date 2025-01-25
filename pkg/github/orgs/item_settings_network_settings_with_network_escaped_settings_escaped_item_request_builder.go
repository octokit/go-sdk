package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\network-settings\{network_settings_id}
type ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal instantiates a new ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder and sets the default values.
func NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    m := &ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/network-settings/{network_settings_id}", pathParameters),
    }
    return m
}
// NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder instantiates a new ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder and sets the default values.
func NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a hosted compute network settings resource configured for an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a NetworkSettingsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/network-configurations#get-a-hosted-compute-network-settings-resource-for-an-organization
func (m *ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateNetworkSettingsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkSettingsable), nil
}
// ToGetRequestInformation gets a hosted compute network settings resource configured for an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder when successful
func (m *ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    return NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
