package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesWithRole_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles\{role_id}
type ItemOrganizationRolesWithRole_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesWithRole_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesWithRole_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOrganizationRolesWithRole_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesWithRole_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOrganizationRolesWithRole_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesWithRole_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOrganizationRolesWithRole_ItemRequestBuilderInternal instantiates a new WithRole_ItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesWithRole_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesWithRole_ItemRequestBuilder) {
    m := &ItemOrganizationRolesWithRole_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles/{role_id}", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesWithRole_ItemRequestBuilder instantiates a new WithRole_ItemRequestBuilder and sets the default values.
func NewItemOrganizationRolesWithRole_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesWithRole_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesWithRole_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a custom organization role.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#delete-a-custom-organization-role
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get gets an organization role that is available to this organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `read_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:read` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:read` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#get-an-organization-role
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable), nil
}
// Patch updates an existing custom organization role. Permission changes will apply to all assignees.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#update-a-custom-organization-role
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) Patch(ctx context.Context, body ItemOrganizationRolesItemWithRole_PatchRequestBodyable, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "409": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable), nil
}
// Teams the teams property
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) Teams()(*ItemOrganizationRolesItemTeamsRequestBuilder) {
    return NewItemOrganizationRolesItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a custom organization role.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation gets an organization role that is available to this organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `read_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:read` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:read` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates an existing custom organization role. Permission changes will apply to all assignees.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemOrganizationRolesItemWithRole_PatchRequestBodyable, requestConfiguration *ItemOrganizationRolesWithRole_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Users the users property
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) Users()(*ItemOrganizationRolesItemUsersRequestBuilder) {
    return NewItemOrganizationRolesItemUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOrganizationRolesWithRole_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesWithRole_ItemRequestBuilder) {
    return NewItemOrganizationRolesWithRole_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
