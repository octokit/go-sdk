package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemWithEnablementItemRequestBuilder builds and executes requests for operations under \orgs\{org}\{security_product}\{enablement}
type ItemItemWithEnablementItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemWithEnablementItemRequestBuilderInternal instantiates a new WithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    m := &ItemItemWithEnablementItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/{security_product}/{enablement}", pathParameters),
    }
    return m
}
// NewItemItemWithEnablementItemRequestBuilder instantiates a new WithEnablementItemRequestBuilder and sets the default values.
func NewItemItemWithEnablementItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithEnablementItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemWithEnablementItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enables or disables the specified security feature for all eligible repositories in an organization.To use this endpoint, you must be an organization owner or be member of a team with the security manager role.A token with the 'write:org' scope is also required.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#enable-or-disable-a-security-feature-for-an-organization
func (m *ItemItemWithEnablementItemRequestBuilder) Post(ctx context.Context, body ItemItemItemWithEnablementPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation enables or disables the specified security feature for all eligible repositories in an organization.To use this endpoint, you must be an organization owner or be member of a team with the security manager role.A token with the 'write:org' scope is also required.GitHub Apps must have the `organization_administration:write` permission to use this endpoint.For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
func (m *ItemItemWithEnablementItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemItemWithEnablementPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemWithEnablementItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemWithEnablementItemRequestBuilder) {
    return NewItemItemWithEnablementItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
