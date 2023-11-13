package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCopilotBillingSelected_usersRequestBuilder builds and executes requests for operations under \orgs\{org}\copilot\billing\selected_users
type ItemCopilotBillingSelected_usersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingSelected_usersRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingSelected_usersRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCopilotBillingSelected_usersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCopilotBillingSelected_usersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
// Delete **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot for Business seat assignment for each user specified.This will cause the specified users to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)"For more information about disabling access to Copilot for Business, see "[Disabling access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#disabling-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// Deprecated: This method is obsolete. Use DeleteAsSelected_usersDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-for-business#remove-users-from-the-copilot-for-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) Delete(ctx context.Context, body ItemCopilotBillingSelected_usersDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderDeleteRequestConfiguration)(ItemCopilotBillingSelected_usersResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_usersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_usersResponseable), nil
}
// DeleteAsSelected_usersDeleteResponse **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot for Business seat assignment for each user specified.This will cause the specified users to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)"For more information about disabling access to Copilot for Business, see "[Disabling access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#disabling-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-for-business#remove-users-from-the-copilot-for-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) DeleteAsSelected_usersDeleteResponse(ctx context.Context, body ItemCopilotBillingSelected_usersDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderDeleteRequestConfiguration)(ItemCopilotBillingSelected_usersDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
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
// Post **Note**: This endpoint is in beta and is subject to change.Purchases a GitHub Copilot for Business seat for each user specified.The organization will be billed accordingly. For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.In order for an admin to use this endpoint, the organization must have a Copilot for Business subscription and a configured suggestion matching policy.For more information about setting up a Copilot for Business subscription, see "[Setting up a Copilot for Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-for-business-subscription-for-your-organization)".For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
// Deprecated: This method is obsolete. Use PostAsSelected_usersPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-for-business#add-users-to-the-copilot-for-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) Post(ctx context.Context, body ItemCopilotBillingSelected_usersPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderPostRequestConfiguration)(ItemCopilotBillingSelected_usersResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSelected_usersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSelected_usersResponseable), nil
}
// PostAsSelected_usersPostResponse **Note**: This endpoint is in beta and is subject to change.Purchases a GitHub Copilot for Business seat for each user specified.The organization will be billed accordingly. For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.In order for an admin to use this endpoint, the organization must have a Copilot for Business subscription and a configured suggestion matching policy.For more information about setting up a Copilot for Business subscription, see "[Setting up a Copilot for Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-for-business-subscription-for-your-organization)".For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-for-business#add-users-to-the-copilot-for-business-subscription-for-an-organization
func (m *ItemCopilotBillingSelected_usersRequestBuilder) PostAsSelected_usersPostResponse(ctx context.Context, body ItemCopilotBillingSelected_usersPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderPostRequestConfiguration)(ItemCopilotBillingSelected_usersPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation **Note**: This endpoint is in beta and is subject to change.Cancels the Copilot for Business seat assignment for each user specified.This will cause the specified users to lose access to GitHub Copilot at the end of the current billing cycle, and the organization will not be billed further for those users.For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)"For more information about disabling access to Copilot for Business, see "[Disabling access to GitHub Copilot for specific users in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#disabling-access-to-github-copilot-for-specific-users-in-your-organization)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemCopilotBillingSelected_usersRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_usersDeleteRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPostRequestInformation **Note**: This endpoint is in beta and is subject to change.Purchases a GitHub Copilot for Business seat for each user specified.The organization will be billed accordingly. For more information about Copilot for Business pricing, see "[About billing for GitHub Copilot for Business](https://docs.github.com/billing/managing-billing-for-github-copilot/about-billing-for-github-copilot#pricing-for-github-copilot-for-business)".Only organization owners and members with admin permissions can configure GitHub Copilot in their organization. You mustauthenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.In order for an admin to use this endpoint, the organization must have a Copilot for Business subscription and a configured suggestion matching policy.For more information about setting up a Copilot for Business subscription, see "[Setting up a Copilot for Business subscription for your organization](https://docs.github.com/billing/managing-billing-for-github-copilot/managing-your-github-copilot-subscription-for-your-organization-or-enterprise#setting-up-a-copilot-for-business-subscription-for-your-organization)".For more information about setting a suggestion matching policy, see "[Configuring suggestion matching policies for GitHub Copilot in your organization](https://docs.github.com/copilot/configuring-github-copilot/configuring-github-copilot-settings-in-your-organization#configuring-suggestion-matching-policies-for-github-copilot-in-your-organization)".
func (m *ItemCopilotBillingSelected_usersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCopilotBillingSelected_usersPostRequestBodyable, requestConfiguration *ItemCopilotBillingSelected_usersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
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
