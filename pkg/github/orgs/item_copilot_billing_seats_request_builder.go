package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCopilotBillingSeatsRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing\seats
type ItemCopilotBillingSeatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingSeatsRequestBuilderGetQueryParameters **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
type ItemCopilotBillingSeatsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
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
// Get **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#list-all-copilot-business-seat-assignments-for-an-organization
func (m *ItemCopilotBillingSeatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCopilotBillingSeatsRequestBuilderGetQueryParameters])(ItemCopilotBillingSeatsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
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
// ToGetRequestInformation **Note**: This endpoint is in beta and is subject to change.Lists all Copilot Business seat assignments for an organization that are currently being billed (either active or pending cancellation at the start of the next billing cycle).Only organization owners can configure and view details about the organization's Copilot Business subscription. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingSeatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCopilotBillingSeatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCopilotBillingSeatsRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingSeatsRequestBuilder) {
    return NewItemCopilotBillingSeatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
