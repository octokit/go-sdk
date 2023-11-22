package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCopilotBillingSeatsRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing\seats
type ItemCopilotBillingSeatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingSeatsRequestBuilderGetQueryParameters **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners and members with admin permissions can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
type ItemCopilotBillingSeatsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemCopilotBillingSeatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingSeatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCopilotBillingSeatsRequestBuilderGetQueryParameters
}
// NewItemCopilotBillingSeatsRequestBuilderInternal instantiates a new SeatsRequestBuilder and sets the default values.
func NewItemCopilotBillingSeatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSeatsRequestBuilder) {
    m := &ItemCopilotBillingSeatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/copilot/billing/seats{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemCopilotBillingSeatsRequestBuilder instantiates a new SeatsRequestBuilder and sets the default values.
func NewItemCopilotBillingSeatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSeatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingSeatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners and members with admin permissions can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsSeatsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#list-all-copilot-business-seat-assignments-for-an-organization
func (m *ItemCopilotBillingSeatsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCopilotBillingSeatsRequestBuilderGetRequestConfiguration)(ItemCopilotBillingSeatsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSeatsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSeatsResponseable), nil
}
// GetAsSeatsGetResponse **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners and members with admin permissions can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#list-all-copilot-business-seat-assignments-for-an-organization
func (m *ItemCopilotBillingSeatsRequestBuilder) GetAsSeatsGetResponse(ctx context.Context, requestConfiguration *ItemCopilotBillingSeatsRequestBuilderGetRequestConfiguration)(ItemCopilotBillingSeatsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSeatsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSeatsGetResponseable), nil
}
// ToGetRequestInformation **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners and members with admin permissions can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingSeatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCopilotBillingSeatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemCopilotBillingSeatsRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingSeatsRequestBuilder) {
    return NewItemCopilotBillingSeatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
