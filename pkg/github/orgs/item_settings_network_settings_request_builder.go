package orgs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSettingsNetworkSettingsRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\network-settings
type ItemSettingsNetworkSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByNetwork_settings_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.settings.networkSettings.item collection
// returns a *ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder when successful
func (m *ItemSettingsNetworkSettingsRequestBuilder) ByNetwork_settings_id(network_settings_id string)(*ItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if network_settings_id != "" {
        urlTplParams["network_settings_id"] = network_settings_id
    }
    return NewItemSettingsNetworkSettingsWithNetwork_settings_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSettingsNetworkSettingsRequestBuilderInternal instantiates a new ItemSettingsNetworkSettingsRequestBuilder and sets the default values.
func NewItemSettingsNetworkSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkSettingsRequestBuilder) {
    m := &ItemSettingsNetworkSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/network-settings", pathParameters),
    }
    return m
}
// NewItemSettingsNetworkSettingsRequestBuilder instantiates a new ItemSettingsNetworkSettingsRequestBuilder and sets the default values.
func NewItemSettingsNetworkSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsNetworkSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
