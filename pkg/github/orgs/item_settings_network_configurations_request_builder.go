package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemSettingsNetworkConfigurationsRequestBuilder builds and executes requests for operations under \orgs\{org}\settings\network-configurations
type ItemSettingsNetworkConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsNetworkConfigurationsRequestBuilderGetQueryParameters lists all hosted compute network configurations configured in an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
type ItemSettingsNetworkConfigurationsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByNetwork_configuration_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.settings.networkConfigurations.item collection
// returns a *ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder when successful
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) ByNetwork_configuration_id(network_configuration_id string)(*ItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if network_configuration_id != "" {
        urlTplParams["network_configuration_id"] = network_configuration_id
    }
    return NewItemSettingsNetworkConfigurationsWithNetwork_configuration_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSettingsNetworkConfigurationsRequestBuilderInternal instantiates a new ItemSettingsNetworkConfigurationsRequestBuilder and sets the default values.
func NewItemSettingsNetworkConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkConfigurationsRequestBuilder) {
    m := &ItemSettingsNetworkConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/settings/network-configurations{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemSettingsNetworkConfigurationsRequestBuilder instantiates a new ItemSettingsNetworkConfigurationsRequestBuilder and sets the default values.
func NewItemSettingsNetworkConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsNetworkConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsNetworkConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all hosted compute network configurations configured in an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a ItemSettingsNetworkConfigurationsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/settings/network-configurations#list-hosted-compute-network-configurations-for-an-organization
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsNetworkConfigurationsRequestBuilderGetQueryParameters])(ItemSettingsNetworkConfigurationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSettingsNetworkConfigurationsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSettingsNetworkConfigurationsGetResponseable), nil
}
// Post creates a hosted compute network configuration for an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// returns a NetworkConfigurationable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/settings/network-configurations#create-a-hosted-compute-network-configuration-for-an-organization
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) Post(ctx context.Context, body ItemSettingsNetworkConfigurationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.NetworkConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation lists all hosted compute network configurations configured in an organization.OAuth app tokens and personal access tokens (classic) need the `read:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSettingsNetworkConfigurationsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a hosted compute network configuration for an organization.OAuth app tokens and personal access tokens (classic) need the `write:network_configurations` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSettingsNetworkConfigurationsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSettingsNetworkConfigurationsRequestBuilder when successful
func (m *ItemSettingsNetworkConfigurationsRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsNetworkConfigurationsRequestBuilder) {
    return NewItemSettingsNetworkConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
