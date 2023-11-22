package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRequestBuilder builds and executes requests for operations under \user
type UserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserGetResponse composed type wrapper for classes privateUser, publicUser
type UserGetResponse struct {
    // Composed type representation for type privateUser
    privateUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable
    // Composed type representation for type publicUser
    publicUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable
}
// NewUserGetResponse instantiates a new userGetResponse and sets the default values.
func NewUserGetResponse()(*UserGetResponse) {
    m := &UserGetResponse{
    }
    return m
}
// CreateUserGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserGetResponse()
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
func (m *UserGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *UserGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *UserGetResponse) GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *UserGetResponse) GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *UserGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPrivateUser() != nil {
        err := writer.WriteObjectValue("", m.GetPrivateUser())
        if err != nil {
            return err
        }
    } else if m.GetPublicUser() != nil {
        err := writer.WriteObjectValue("", m.GetPublicUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPrivateUser sets the privateUser property value. Composed type representation for type privateUser
func (m *UserGetResponse) SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *UserGetResponse) SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)() {
    m.publicUser = value
}
// UserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserResponse composed type wrapper for classes privateUser, publicUser
type UserResponse struct {
    // Composed type representation for type privateUser
    privateUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable
    // Composed type representation for type publicUser
    publicUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable
}
// NewUserResponse instantiates a new userResponse and sets the default values.
func NewUserResponse()(*UserResponse) {
    m := &UserResponse{
    }
    return m
}
// CreateUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserResponse()
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
func (m *UserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *UserResponse) GetIsComposedType()(bool) {
    return true
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *UserResponse) GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *UserResponse) GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *UserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPrivateUser() != nil {
        err := writer.WriteObjectValue("", m.GetPrivateUser())
        if err != nil {
            return err
        }
    } else if m.GetPublicUser() != nil {
        err := writer.WriteObjectValue("", m.GetPublicUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPrivateUser sets the privateUser property value. Composed type representation for type privateUser
func (m *UserResponse) SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *UserResponse) SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)() {
    m.publicUser = value
}
// UserGetResponseable 
type UserGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)
    GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)
    SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)()
    SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)()
}
// UserResponseable 
type UserResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)
    GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)
    SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)()
    SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)()
}
// Blocks the blocks property
func (m *UserRequestBuilder) Blocks()(*BlocksRequestBuilder) {
    return NewBlocksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codespaces the codespaces property
func (m *UserRequestBuilder) Codespaces()(*CodespacesRequestBuilder) {
    return NewCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user", pathParameters),
    }
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Docker the docker property
func (m *UserRequestBuilder) Docker()(*DockerRequestBuilder) {
    return NewDockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Email the email property
func (m *UserRequestBuilder) Email()(*EmailRequestBuilder) {
    return NewEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Emails the emails property
func (m *UserRequestBuilder) Emails()(*EmailsRequestBuilder) {
    return NewEmailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Followers the followers property
func (m *UserRequestBuilder) Followers()(*FollowersRequestBuilder) {
    return NewFollowersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Following the following property
func (m *UserRequestBuilder) Following()(*FollowingRequestBuilder) {
    return NewFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get if the authenticated user is authenticated with an OAuth token with the `user` scope, then the response lists public and private profile information.If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.
// Deprecated: This method is obsolete. Use GetAsUserGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-the-authenticated-user
func (m *UserRequestBuilder) Get(ctx context.Context, requestConfiguration *UserRequestBuilderGetRequestConfiguration)(UserResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserResponseable), nil
}
// GetAsUserGetResponse if the authenticated user is authenticated with an OAuth token with the `user` scope, then the response lists public and private profile information.If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-the-authenticated-user
func (m *UserRequestBuilder) GetAsUserGetResponse(ctx context.Context, requestConfiguration *UserRequestBuilderGetRequestConfiguration)(UserGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserGetResponseable), nil
}
// Gpg_keys the gpg_keys property
func (m *UserRequestBuilder) Gpg_keys()(*Gpg_keysRequestBuilder) {
    return NewGpg_keysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installations the installations property
func (m *UserRequestBuilder) Installations()(*InstallationsRequestBuilder) {
    return NewInstallationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InteractionLimits the interactionLimits property
func (m *UserRequestBuilder) InteractionLimits()(*InteractionLimitsRequestBuilder) {
    return NewInteractionLimitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *UserRequestBuilder) Issues()(*IssuesRequestBuilder) {
    return NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys the keys property
func (m *UserRequestBuilder) Keys()(*KeysRequestBuilder) {
    return NewKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_purchases the marketplace_purchases property
func (m *UserRequestBuilder) Marketplace_purchases()(*Marketplace_purchasesRequestBuilder) {
    return NewMarketplace_purchasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Memberships the memberships property
func (m *UserRequestBuilder) Memberships()(*MembershipsRequestBuilder) {
    return NewMembershipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Migrations the migrations property
func (m *UserRequestBuilder) Migrations()(*MigrationsRequestBuilder) {
    return NewMigrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
func (m *UserRequestBuilder) Orgs()(*OrgsRequestBuilder) {
    return NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Packages the packages property
func (m *UserRequestBuilder) Packages()(*PackagesRequestBuilder) {
    return NewPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch **Note:** If your email is set to private and you send an `email` parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#update-the-authenticated-user
func (m *UserRequestBuilder) Patch(ctx context.Context, body UserPatchRequestBodyable, requestConfiguration *UserRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePrivateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable), nil
}
// Projects the projects property
func (m *UserRequestBuilder) Projects()(*ProjectsRequestBuilder) {
    return NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Public_emails the public_emails property
func (m *UserRequestBuilder) Public_emails()(*Public_emailsRequestBuilder) {
    return NewPublic_emailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *UserRequestBuilder) Repos()(*ReposRequestBuilder) {
    return NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repository_invitations the repository_invitations property
func (m *UserRequestBuilder) Repository_invitations()(*Repository_invitationsRequestBuilder) {
    return NewRepository_invitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Social_accounts the social_accounts property
func (m *UserRequestBuilder) Social_accounts()(*Social_accountsRequestBuilder) {
    return NewSocial_accountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ssh_signing_keys the ssh_signing_keys property
func (m *UserRequestBuilder) Ssh_signing_keys()(*Ssh_signing_keysRequestBuilder) {
    return NewSsh_signing_keysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Starred the starred property
func (m *UserRequestBuilder) Starred()(*StarredRequestBuilder) {
    return NewStarredRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscriptions the subscriptions property
func (m *UserRequestBuilder) Subscriptions()(*SubscriptionsRequestBuilder) {
    return NewSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *UserRequestBuilder) Teams()(*TeamsRequestBuilder) {
    return NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation if the authenticated user is authenticated with an OAuth token with the `user` scope, then the response lists public and private profile information.If the authenticated user is authenticated through OAuth without the `user` scope, then the response lists only public profile information.
func (m *UserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation **Note:** If your email is set to private and you send an `email` parameter as part of this request to update your profile, your privacy settings are still enforced: the email address will not be displayed on your public profile or via the API.
func (m *UserRequestBuilder) ToPatchRequestInformation(ctx context.Context, body UserPatchRequestBodyable, requestConfiguration *UserRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserRequestBuilder) WithUrl(rawUrl string)(*UserRequestBuilder) {
    return NewUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
