package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOrganizationRolesRequestBuilder builds and executes requests for operations under \orgs\{org}\organization-roles
type ItemOrganizationRolesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOrganizationRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOrganizationRolesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOrganizationRolesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRole_id gets an item from the octokit.orgs.item.organizationRoles.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemOrganizationRolesRequestBuilder) ByRole_id(role_id string)(*ItemOrganizationRolesWithRole_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if role_id != "" {
        urlTplParams["role_id"] = role_id
    }
    return NewItemOrganizationRolesWithRole_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRole_idInteger gets an item from the octokit.orgs.item.organizationRoles.item collection
func (m *ItemOrganizationRolesRequestBuilder) ByRole_idInteger(role_id int32)(*ItemOrganizationRolesWithRole_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["role_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(role_id), 10)
    return NewItemOrganizationRolesWithRole_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOrganizationRolesRequestBuilderInternal instantiates a new OrganizationRolesRequestBuilder and sets the default values.
func NewItemOrganizationRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesRequestBuilder) {
    m := &ItemOrganizationRolesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/organization-roles", pathParameters),
    }
    return m
}
// NewItemOrganizationRolesRequestBuilder instantiates a new OrganizationRolesRequestBuilder and sets the default values.
func NewItemOrganizationRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOrganizationRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOrganizationRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the organization roles available in this organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `read_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:read` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:read` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// Deprecated: This method is obsolete. Use GetAsOrganizationRolesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#get-all-organization-roles-for-an-organization
func (m *ItemOrganizationRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOrganizationRolesRequestBuilderGetRequestConfiguration)(ItemOrganizationRolesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOrganizationRolesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOrganizationRolesResponseable), nil
}
// GetAsOrganizationRolesGetResponse lists the organization roles available in this organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `read_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:read` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:read` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#get-all-organization-roles-for-an-organization
func (m *ItemOrganizationRolesRequestBuilder) GetAsOrganizationRolesGetResponse(ctx context.Context, requestConfiguration *ItemOrganizationRolesRequestBuilderGetRequestConfiguration)(ItemOrganizationRolesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOrganizationRolesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOrganizationRolesGetResponseable), nil
}
// Post creates a custom organization role that can be assigned to users and teams, granting them specific permissions over the organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/organization-roles#create-a-custom-organization-role
func (m *ItemOrganizationRolesRequestBuilder) Post(ctx context.Context, body ItemOrganizationRolesPostRequestBodyable, requestConfiguration *ItemOrganizationRolesRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
func (m *ItemOrganizationRolesRequestBuilder) Teams()(*ItemOrganizationRolesTeamsRequestBuilder) {
    return NewItemOrganizationRolesTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists the organization roles available in this organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `read_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:read` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:read` organization permission to use this endpoint.For more information on organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOrganizationRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation creates a custom organization role that can be assigned to users and teams, granting them specific permissions over the organization.To use this endpoint, the authenticated user must be one of:- An administrator for the organization.- A user, or a user on a team, with the fine-grained permissions of `write_organization_custom_org_role` in the organization.The authenticated user needs an access token with `admin:org` scope or a fine-grained personal access token with the `organization_custom_roles:write` permission to use this endpoint.GitHub Apps must have the `organization_custom_org_roles:write` organization permission to use this endpoint.For more information on custom organization roles, see "[Managing people's access to your organization with roles](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/about-custom-organization-roles)."
func (m *ItemOrganizationRolesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOrganizationRolesPostRequestBodyable, requestConfiguration *ItemOrganizationRolesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Users the users property
func (m *ItemOrganizationRolesRequestBuilder) Users()(*ItemOrganizationRolesUsersRequestBuilder) {
    return NewItemOrganizationRolesUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOrganizationRolesRequestBuilder) WithUrl(rawUrl string)(*ItemOrganizationRolesRequestBuilder) {
    return NewItemOrganizationRolesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
