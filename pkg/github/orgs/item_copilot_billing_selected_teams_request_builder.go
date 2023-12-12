package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCopilotBillingSelected_teamsRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing\selected_teams
type ItemCopilotBillingSelected_teamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingSelected_teamsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingSelected_teamsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCopilotBillingSelected_teamsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingSelected_teamsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCopilotBillingSelected_teamsRequestBuilderInternal instantiates a new Selected_teamsRequestBuilder and sets the default values.
func NewItemCopilotBillingSelected_teamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSelected_teamsRequestBuilder) {
    m := &ItemCopilotBillingSelected_teamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/copilot/billing/selected_teams", pathParameters),
    }
    return m
}
// NewItemCopilotBillingSelected_teamsRequestBuilder instantiates a new Selected_teamsRequestBuilder and sets the default values.
func NewItemCopilotBillingSelected_teamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSelected_teamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingSelected_teamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot Business seat assignment for all members of each team specified.This will cause the members of the specified team(s) to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)".For more information about disabling access to Copilot Business, see "[Revoking access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-access-for-copilot-business-in-your-organization#revoking-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// Deprecated: This method is obsolete. Use DeleteAsSelected_teamsDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#remove-teams-from-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) Delete(ctx context.Context, body ItemCopilotBillingSelected_teamsDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderDeleteRequestConfiguration)(ItemCopilotBillingSelected_teamsResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_teamsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_teamsResponseable), nil
}
// DeleteAsSelected_teamsDeleteResponse **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot Business seat assignment for all members of each team specified.This will cause the members of the specified team(s) to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)".For more information about disabling access to Copilot Business, see "[Revoking access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-access-for-copilot-business-in-your-organization#revoking-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#remove-teams-from-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) DeleteAsSelected_teamsDeleteResponse(ctx context.Context, body ItemCopilotBillingSelected_teamsDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderDeleteRequestConfiguration)(ItemCopilotBillingSelected_teamsDeleteResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_teamsDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_teamsDeleteResponseable), nil
}
// Post **Note**: This endpoint is in beta and is subject to change. Purchases a GitHub Copilot Business seat for all users within each specified team. The organization will be billed accordingly. For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)". Only organization owners can configure GitHub Copilot in their organization. You must authenticate using an access token with the `manage_billing:copilot` scope to use this endpoint. In order for an admin to use this endpoint, the organization must have a Copilot Business subscription and a configured suggestion matching policy. For more information about setting up a Copilot Business subscription, see "[Setting up a Copilot Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-business-subscription-for-your-organization)". For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-policies-for-copilot-business-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
// Deprecated: This method is obsolete. Use PostAsSelected_teamsPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#add-teams-to-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) Post(ctx context.Context, body ItemCopilotBillingSelected_teamsPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderPostRequestConfiguration)(ItemCopilotBillingSelected_teamsResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_teamsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_teamsResponseable), nil
}
// PostAsSelected_teamsPostResponse **Note**: This endpoint is in beta and is subject to change. Purchases a GitHub Copilot Business seat for all users within each specified team. The organization will be billed accordingly. For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)". Only organization owners can configure GitHub Copilot in their organization. You must authenticate using an access token with the `manage_billing:copilot` scope to use this endpoint. In order for an admin to use this endpoint, the organization must have a Copilot Business subscription and a configured suggestion matching policy. For more information about setting up a Copilot Business subscription, see "[Setting up a Copilot Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-business-subscription-for-your-organization)". For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-policies-for-copilot-business-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#add-teams-to-the-copilot-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) PostAsSelected_teamsPostResponse(ctx context.Context, body ItemCopilotBillingSelected_teamsPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderPostRequestConfiguration)(ItemCopilotBillingSelected_teamsPostResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_teamsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_teamsPostResponseable), nil
}
// ToDeleteRequestInformation **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot Business seat assignment for all members of each team specified.This will cause the members of the specified team(s) to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)".For more information about disabling access to Copilot Business, see "[Revoking access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-access-for-copilot-business-in-your-organization#revoking-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_teamsDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPostRequestInformation **Note**: This endpoint is in beta and is subject to change. Purchases a GitHub Copilot Business seat for all users within each specified team. The organization will be billed accordingly. For more information about Copilot Business pricing, see "[Pricing for GitHub Copilot Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-business)". Only organization owners can configure GitHub Copilot in their organization. You must authenticate using an access token with the `manage_billing:copilot` scope to use this endpoint. In order for an admin to use this endpoint, the organization must have a Copilot Business subscription and a configured suggestion matching policy. For more information about setting up a Copilot Business subscription, see "[Setting up a Copilot Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-business-subscription-for-your-organization)". For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/managing-copilot-business/managing-policies-for-copilot-business-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_teamsPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_teamsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCopilotBillingSelected_teamsRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingSelected_teamsRequestBuilder) {
    return NewItemCopilotBillingSelected_teamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
