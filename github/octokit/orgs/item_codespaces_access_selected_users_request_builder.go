package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemCodespacesAccessSelected_usersRequestBuilder builds and executes requests for operations under \orgs\{org}\codespaces\access\selected_users
type ItemCodespacesAccessSelected_usersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCodespacesAccessSelected_usersRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesAccessSelected_usersRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCodespacesAccessSelected_usersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCodespacesAccessSelected_usersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCodespacesAccessSelected_usersRequestBuilderInternal instantiates a new Selected_usersRequestBuilder and sets the default values.
func NewItemCodespacesAccessSelected_usersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesAccessSelected_usersRequestBuilder) {
    m := &ItemCodespacesAccessSelected_usersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/codespaces/access/selected_users", pathParameters),
    }
    return m
}
// NewItemCodespacesAccessSelected_usersRequestBuilder instantiates a new Selected_usersRequestBuilder and sets the default values.
func NewItemCodespacesAccessSelected_usersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCodespacesAccessSelected_usersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCodespacesAccessSelected_usersRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete codespaces for the specified users will no longer be billed to the organization.To use this endpoint, the access settings for the organization must be set to `selected_members`.For information on how to change this setting, see "[Manage access control for organization codespaces](https://docs.github.com/rest/codespaces/organizations#manage-access-control-for-organization-codespaces)."You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces settings` write permission to use this endpoint.
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#remove-users-from-codespaces-access-for-an-organization
func (m *ItemCodespacesAccessSelected_usersRequestBuilder) Delete(ctx context.Context, body ItemCodespacesAccessSelected_usersDeleteRequestBodyable, requestConfiguration *ItemCodespacesAccessSelected_usersRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Post codespaces for the specified users will be billed to the organization.To use this endpoint, the access settings for the organization must be set to `selected_members`.For information on how to change this setting, see "[Manage access control for organization codespaces](https://docs.github.com/rest/codespaces/organizations#manage-access-control-for-organization-codespaces)."You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces settings` write permission to use this endpoint.
// Deprecated: 
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#add-users-to-codespaces-access-for-an-organization
func (m *ItemCodespacesAccessSelected_usersRequestBuilder) Post(ctx context.Context, body ItemCodespacesAccessSelected_usersPostRequestBodyable, requestConfiguration *ItemCodespacesAccessSelected_usersRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation codespaces for the specified users will no longer be billed to the organization.To use this endpoint, the access settings for the organization must be set to `selected_members`.For information on how to change this setting, see "[Manage access control for organization codespaces](https://docs.github.com/rest/codespaces/organizations#manage-access-control-for-organization-codespaces)."You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces settings` write permission to use this endpoint.
// Deprecated: 
func (m *ItemCodespacesAccessSelected_usersRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemCodespacesAccessSelected_usersDeleteRequestBodyable, requestConfiguration *ItemCodespacesAccessSelected_usersRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation codespaces for the specified users will be billed to the organization.To use this endpoint, the access settings for the organization must be set to `selected_members`.For information on how to change this setting, see "[Manage access control for organization codespaces](https://docs.github.com/rest/codespaces/organizations#manage-access-control-for-organization-codespaces)."You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces settings` write permission to use this endpoint.
// Deprecated: 
func (m *ItemCodespacesAccessSelected_usersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCodespacesAccessSelected_usersPostRequestBodyable, requestConfiguration *ItemCodespacesAccessSelected_usersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
func (m *ItemCodespacesAccessSelected_usersRequestBuilder) WithUrl(rawUrl string)(*ItemCodespacesAccessSelected_usersRequestBuilder) {
    return NewItemCodespacesAccessSelected_usersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
