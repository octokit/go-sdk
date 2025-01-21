package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\network-configurations\{network_configuration_id}
type ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilderInternal instantiates a new ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder and sets the default values.
func NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) {
    m := &ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/network-configurations/{network_configuration_id}", pathParameters),
    }
    return m
}
// NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder instantiates a new ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder and sets the default values.
func NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a hosted compute network configuration from an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/settings/network-configurations#delete-a-hosted-compute-network-configuration-from-an-organization
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get gets a hosted compute network configuration configured in an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a NetworkConfigurationable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/settings/network-configurations#get-a-hosted-compute-network-configuration-for-an-organization
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateNetworkConfigurationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable), nil
}
// Patch updates a hosted compute network configuration for an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// returns a NetworkConfigurationable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/settings/network-configurations#update-a-hosted-compute-network-configuration-for-an-organization
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) Patch(ctx context.Context, body ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateNetworkConfigurationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable), nil
}
// ToDeleteRequestInformation deletes a hosted compute network configuration from an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation gets a hosted compute network configuration configured in an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a hosted compute network configuration for an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemSettingsNetworkConfigurationsItemWithNetwork_configuration_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder when successful
func (m *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) {
    return NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
