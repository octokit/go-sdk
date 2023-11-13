package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersWithUsernameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}
type ItemMembersWithUsernameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersWithUsernameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithUsernameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMembersWithUsernameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithUsernameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Codespaces the codespaces property
func (m *ItemMembersWithUsernameItemRequestBuilder) Codespaces()(*ItemMembersItemCodespacesRequestBuilder) {
    return NewItemMembersItemCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMembersWithUsernameItemRequestBuilderInternal instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewItemMembersWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithUsernameItemRequestBuilder) {
    m := &ItemMembersWithUsernameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/members/{username}", pathParameters),
    }
    return m
}
// NewItemMembersWithUsernameItemRequestBuilder instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewItemMembersWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copilot the copilot property
func (m *ItemMembersWithUsernameItemRequestBuilder) Copilot()(*ItemMembersItemCopilotRequestBuilder) {
    return NewItemMembersItemCopilotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete removing a user from this list will remove them from all teams and they will no longer have any access to the organization's repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#remove-an-organization-member
func (m *ItemMembersWithUsernameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMembersWithUsernameItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get check if a user is, publicly or privately, a member of the organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#check-organization-membership-for-a-user
func (m *ItemMembersWithUsernameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithUsernameItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation removing a user from this list will remove them from all teams and they will no longer have any access to the organization's repositories.
func (m *ItemMembersWithUsernameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithUsernameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation check if a user is, publicly or privately, a member of the organization.
func (m *ItemMembersWithUsernameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithUsernameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMembersWithUsernameItemRequestBuilder) WithUrl(rawUrl string)(*ItemMembersWithUsernameItemRequestBuilder) {
    return NewItemMembersWithUsernameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
