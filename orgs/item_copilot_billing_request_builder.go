package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCopilotBillingRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing
type ItemCopilotBillingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCopilotBillingRequestBuilderInternal instantiates a new BillingRequestBuilder and sets the default values.
func NewItemCopilotBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingRequestBuilder) {
    m := &ItemCopilotBillingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/copilot/billing", pathParameters),
    }
    return m
}
// NewItemCopilotBillingRequestBuilder instantiates a new BillingRequestBuilder and sets the default values.
func NewItemCopilotBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note**: This endpoint is in beta and is subject to change.Gets information about an organization's Copilot for Business subscription, including seat breakdownand code matching policies. To configure these settings, go to your organization's settings on GitHub.com.For more information, see "[Configuring GitHub Copilot settings in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization)".Only organization owners and members with admin permissions can configure and view details about the organization's Copilot for Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-for-business#get-copilot-for-business-seat-information-and-settings-for-an-organization
func (m *ItemCopilotBillingRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCopilotBillingRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CopilotOrganizationDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCopilotOrganizationDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CopilotOrganizationDetailsable), nil
}
// Seats the seats property
func (m *ItemCopilotBillingRequestBuilder) Seats()(*ItemCopilotBillingSeatsRequestBuilder) {
    return NewItemCopilotBillingSeatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Selected_teams the selected_teams property
func (m *ItemCopilotBillingRequestBuilder) Selected_teams()(*ItemCopilotBillingSelected_teamsRequestBuilder) {
    return NewItemCopilotBillingSelected_teamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Selected_users the selected_users property
func (m *ItemCopilotBillingRequestBuilder) Selected_users()(*ItemCopilotBillingSelected_usersRequestBuilder) {
    return NewItemCopilotBillingSelected_usersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation **Note**: This endpoint is in beta and is subject to change.Gets information about an organization's Copilot for Business subscription, including seat breakdownand code matching policies. To configure these settings, go to your organization's settings on GitHub.com.For more information, see "[Configuring GitHub Copilot settings in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization)".Only organization owners and members with admin permissions can configure and view details about the organization's Copilot for Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCopilotBillingRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCopilotBillingRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingRequestBuilder) {
    return NewItemCopilotBillingRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
