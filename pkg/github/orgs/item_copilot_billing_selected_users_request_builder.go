package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCopilotBillingSelected_usersRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing\selected_users
type ItemCopilotBillingSelected_usersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemCopilotBillingSelected_usersRequestBuilderInternal instantiates a new Selected_usersRequestBuilder and sets the default values.
func NewItemCopilotBillingSelected_usersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSelected_usersRequestBuilder) {
    m := &ItemCopilotBillingSelected_usersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/copilot/billing/selected_users", pathParameters),
    }
    return m
}
// NewItemCopilotBillingSelected_usersRequestBuilder instantiates a new Selected_usersRequestBuilder and sets the default values.
func NewItemCopilotBillingSelected_usersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSelected_usersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingSelected_usersRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot Business seat assignment for each user specified.This will cause the specified users to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)"For more information about disabling access to Copilot Business, see "[Revoking access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-access-for-copilot-business-in-your-organization#revoking-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#remove-users-from-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) Delete(ctx context.Context, body ItemCopilotBillingSelected_usersDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemCopilotBillingSelected_usersDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_usersDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_usersDeleteResponseable), nil
}
// Post **Note**: This endpoint is in beta and is subject to change.Purchases a GitHub Copilot Business seat for each user specified.The organization will be billed accordingly. For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.In order for an admin to use this endpoint, the organization must have a Copilot Business subscription and a configured suggestion matching policy.For more information about setting up a Copilot Business subscription, see "[Setting up a Copilot Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-business-subscription-for-your-organization)".For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-policies-for-copilot-business-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#add-users-to-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) Post(ctx context.Context, body ItemCopilotBillingSelected_usersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ItemCopilotBillingSelected_usersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_usersPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_usersPostResponseable), nil
}
// ToDeleteRequestInformation **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot Business seat assignment for each user specified.This will cause the specified users to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)"For more information about disabling access to Copilot Business, see "[Revoking access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-access-for-copilot-business-in-your-organization#revoking-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingSelected_usersRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_usersDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPostRequestInformation **Note**: This endpoint is in beta and is subject to change.Purchases a GitHub Copilot Business seat for each user specified.The organization will be billed accordingly. For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.In order for an admin to use this endpoint, the organization must have a Copilot Business subscription and a configured suggestion matching policy.For more information about setting up a Copilot Business subscription, see "[Setting up a Copilot Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-business-subscription-for-your-organization)".For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-policies-for-copilot-business-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
func (m *ItemCopilotBillingSelected_usersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_usersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemCopilotBillingSelected_usersRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingSelected_usersRequestBuilder) {
    return NewItemCopilotBillingSelected_usersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
