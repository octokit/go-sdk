package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\users\{username}\{role_id}
type ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderInternal instantiates a new WithRole_ItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) {
    m := &ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/users/{username}/{role_id}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder instantiates a new WithRole_ItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an organization role from a user.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members:write` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#remove-an-organization-role-from-a-user
func (m *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Put assigns an organization role to a member of an organization.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read-write permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#assign-an-organization-role-to-a-user
func (m *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) Put(ctx context.Context, requestConfiguration *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation remove an organization role from a user.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members:write` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToPutRequestInformation assigns an organization role to a member of an organization.To use this endpoint, you must be an administrator for the organization, and you must use an access token with the `admin:org` scope.GitHub Apps must have the `members` organization read-write permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder) {
    return NewItemOrganizationRolesUsersItemWithRole_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
