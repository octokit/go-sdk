package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WithOrgItemRequestBuilder builds and executes requests for operations under \orgs\{org}
type WithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OrganizationFull422Error composed type wrapper for classes validationError, validationErrorSimple
type OrganizationFull422Error struct {
    // Composed type representation for type validationError
    validationError i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable
    // Composed type representation for type validationErrorSimple
    validationErrorSimple i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable
}
// NewOrganizationFull422Error instantiates a new organizationFull422Error and sets the default values.
func NewOrganizationFull422Error()(*OrganizationFull422Error) {
    m := &OrganizationFull422Error{
    }
    return m
}
// CreateOrganizationFull422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationFull422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOrganizationFull422Error()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationFull422Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *OrganizationFull422Error) GetIsComposedType()(bool) {
    return true
}
// GetValidationError gets the validationError property value. Composed type representation for type validationError
func (m *OrganizationFull422Error) GetValidationError()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable) {
    return m.validationError
}
// GetValidationErrorSimple gets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *OrganizationFull422Error) GetValidationErrorSimple()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable) {
    return m.validationErrorSimple
}
// Serialize serializes information the current object
func (m *OrganizationFull422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValidationError() != nil {
        err := writer.WriteObjectValue("", m.GetValidationError())
        if err != nil {
            return err
        }
    } else if m.GetValidationErrorSimple() != nil {
        err := writer.WriteObjectValue("", m.GetValidationErrorSimple())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValidationError sets the validationError property value. Composed type representation for type validationError
func (m *OrganizationFull422Error) SetValidationError(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)() {
    m.validationError = value
}
// SetValidationErrorSimple sets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *OrganizationFull422Error) SetValidationErrorSimple(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)() {
    m.validationErrorSimple = value
}
// WithOrgItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithOrgItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithOrgItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithOrgItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithOrgItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithOrgItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationFull422Errorable 
type OrganizationFull422Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValidationError()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)
    GetValidationErrorSimple()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)
    SetValidationError(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)()
    SetValidationErrorSimple(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)()
}
// Actions the actions property
func (m *WithOrgItemRequestBuilder) Actions()(*ItemActionsRequestBuilder) {
    return NewItemActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Blocks the blocks property
func (m *WithOrgItemRequestBuilder) Blocks()(*ItemBlocksRequestBuilder) {
    return NewItemBlocksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BySecurity_product gets an item from the octokit.orgs.item.item collection
func (m *WithOrgItemRequestBuilder) BySecurity_product(security_product string)(*ItemWithSecurity_productItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if security_product != "" {
        urlTplParams["security_product"] = security_product
    }
    return NewItemWithSecurity_productItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CodeScanning the codeScanning property
func (m *WithOrgItemRequestBuilder) CodeScanning()(*ItemCodeScanningRequestBuilder) {
    return NewItemCodeScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codespaces the codespaces property
func (m *WithOrgItemRequestBuilder) Codespaces()(*ItemCodespacesRequestBuilder) {
    return NewItemCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithOrgItemRequestBuilderInternal instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    m := &WithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}", pathParameters),
    }
    return m
}
// NewWithOrgItemRequestBuilder instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copilot the copilot property
func (m *WithOrgItemRequestBuilder) Copilot()(*ItemCopilotRequestBuilder) {
    return NewItemCopilotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete deletes an organization and all its repositories.The organization login will be unavailable for 90 days after deletion.Please review the Terms of Service regarding account deletion before using this endpoint:https://docs.github.com/site-policy/github-terms/github-terms-of-service
// Deprecated: This method is obsolete. Use DeleteAsWithOrgDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#delete-an-organization
func (m *WithOrgItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderDeleteRequestConfiguration)(ItemWithOrgResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemWithOrgResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemWithOrgResponseable), nil
}
// DeleteAsWithOrgDeleteResponse deletes an organization and all its repositories.The organization login will be unavailable for 90 days after deletion.Please review the Terms of Service regarding account deletion before using this endpoint:https://docs.github.com/site-policy/github-terms/github-terms-of-service
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#delete-an-organization
func (m *WithOrgItemRequestBuilder) DeleteAsWithOrgDeleteResponse(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderDeleteRequestConfiguration)(ItemWithOrgDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemWithOrgDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemWithOrgDeleteResponseable), nil
}
// Dependabot the dependabot property
func (m *WithOrgItemRequestBuilder) Dependabot()(*ItemDependabotRequestBuilder) {
    return NewItemDependabotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Docker the docker property
func (m *WithOrgItemRequestBuilder) Docker()(*ItemDockerRequestBuilder) {
    return NewItemDockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *WithOrgItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Failed_invitations the failed_invitations property
func (m *WithOrgItemRequestBuilder) Failed_invitations()(*ItemFailed_invitationsRequestBuilder) {
    return NewItemFailed_invitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get to see many of the organization response values, you need to be an authenticated organization owner with the `admin:org` scope. When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/).GitHub Apps with the `Organization plan` permission can use this endpoint to retrieve information about an organization's GitHub plan. See "[Authenticating with GitHub Apps](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/)" for details. For an example response, see 'Response with GitHub plan information' below."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#get-an-organization
func (m *WithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationFullable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationFullable), nil
}
// Hooks the hooks property
func (m *WithOrgItemRequestBuilder) Hooks()(*ItemHooksRequestBuilder) {
    return NewItemHooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *WithOrgItemRequestBuilder) Installation()(*ItemInstallationRequestBuilder) {
    return NewItemInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installations the installations property
func (m *WithOrgItemRequestBuilder) Installations()(*ItemInstallationsRequestBuilder) {
    return NewItemInstallationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InteractionLimits the interactionLimits property
func (m *WithOrgItemRequestBuilder) InteractionLimits()(*ItemInteractionLimitsRequestBuilder) {
    return NewItemInteractionLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Invitations the invitations property
func (m *WithOrgItemRequestBuilder) Invitations()(*ItemInvitationsRequestBuilder) {
    return NewItemInvitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *WithOrgItemRequestBuilder) Issues()(*ItemIssuesRequestBuilder) {
    return NewItemIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members the members property
func (m *WithOrgItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Memberships the memberships property
func (m *WithOrgItemRequestBuilder) Memberships()(*ItemMembershipsRequestBuilder) {
    return NewItemMembershipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Migrations the migrations property
func (m *WithOrgItemRequestBuilder) Migrations()(*ItemMigrationsRequestBuilder) {
    return NewItemMigrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Outside_collaborators the outside_collaborators property
func (m *WithOrgItemRequestBuilder) Outside_collaborators()(*ItemOutside_collaboratorsRequestBuilder) {
    return NewItemOutside_collaboratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Packages the packages property
func (m *WithOrgItemRequestBuilder) Packages()(*ItemPackagesRequestBuilder) {
    return NewItemPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch **Parameter Deprecation Notice:** GitHub will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Enables an authenticated organization owner with the `admin:org` scope or the `repo` scope to update the organization's profile and member privileges.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/orgs#update-an-organization
func (m *WithOrgItemRequestBuilder) Patch(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *WithOrgItemRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationFullable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "409": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": CreateOrganizationFull422ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationFullFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationFullable), nil
}
// PersonalAccessTokenRequests the personalAccessTokenRequests property
func (m *WithOrgItemRequestBuilder) PersonalAccessTokenRequests()(*ItemPersonalAccessTokenRequestsRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PersonalAccessTokens the personalAccessTokens property
func (m *WithOrgItemRequestBuilder) PersonalAccessTokens()(*ItemPersonalAccessTokensRequestBuilder) {
    return NewItemPersonalAccessTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *WithOrgItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Properties the properties property
func (m *WithOrgItemRequestBuilder) Properties()(*ItemPropertiesRequestBuilder) {
    return NewItemPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Public_members the public_members property
func (m *WithOrgItemRequestBuilder) Public_members()(*ItemPublic_membersRequestBuilder) {
    return NewItemPublic_membersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *WithOrgItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rulesets the rulesets property
func (m *WithOrgItemRequestBuilder) Rulesets()(*ItemRulesetsRequestBuilder) {
    return NewItemRulesetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecretScanning the secretScanning property
func (m *WithOrgItemRequestBuilder) SecretScanning()(*ItemSecretScanningRequestBuilder) {
    return NewItemSecretScanningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityAdvisories the securityAdvisories property
func (m *WithOrgItemRequestBuilder) SecurityAdvisories()(*ItemSecurityAdvisoriesRequestBuilder) {
    return NewItemSecurityAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityManagers the securityManagers property
func (m *WithOrgItemRequestBuilder) SecurityManagers()(*ItemSecurityManagersRequestBuilder) {
    return NewItemSecurityManagersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
func (m *WithOrgItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *WithOrgItemRequestBuilder) Teams()(*ItemTeamsRequestBuilder) {
    return NewItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes an organization and all its repositories.The organization login will be unavailable for 90 days after deletion.Please review the Terms of Service regarding account deletion before using this endpoint:https://docs.github.com/site-policy/github-terms/github-terms-of-service
func (m *WithOrgItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToGetRequestInformation to see many of the organization response values, you need to be an authenticated organization owner with the `admin:org` scope. When the value of `two_factor_requirement_enabled` is `true`, the organization requires all members, billing managers, and outside collaborators to enable [two-factor authentication](https://docs.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/).GitHub Apps with the `Organization plan` permission can use this endpoint to retrieve information about an organization's GitHub plan. See "[Authenticating with GitHub Apps](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/)" for details. For an example response, see 'Response with GitHub plan information' below."
func (m *WithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithOrgItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Parameter Deprecation Notice:** GitHub will replace and discontinue `members_allowed_repository_creation_type` in favor of more granular permissions. The new input parameters are `members_can_create_public_repositories`, `members_can_create_private_repositories` for all organizations and `members_can_create_internal_repositories` for organizations associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see the [blog post](https://developer.github.com/changes/2019-12-03-internal-visibility-changes).Enables an authenticated organization owner with the `admin:org` scope or the `repo` scope to update the organization's profile and member privileges.
func (m *WithOrgItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithOrgPatchRequestBodyable, requestConfiguration *WithOrgItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WithOrgItemRequestBuilder) WithUrl(rawUrl string)(*WithOrgItemRequestBuilder) {
    return NewWithOrgItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
