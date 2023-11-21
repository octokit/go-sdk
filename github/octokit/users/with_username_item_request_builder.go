package users

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WithUsernameItemRequestBuilder builds and executes requests for operations under \users\{username}
type WithUsernameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithUsernameGetResponse composed type wrapper for classes privateUser, publicUser
type WithUsernameGetResponse struct {
    // Composed type representation for type privateUser
    privateUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable
    // Composed type representation for type publicUser
    publicUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable
}
// NewWithUsernameGetResponse instantiates a new WithUsernameGetResponse and sets the default values.
func NewWithUsernameGetResponse()(*WithUsernameGetResponse) {
    m := &WithUsernameGetResponse{
    }
    return m
}
// CreateWithUsernameGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithUsernameGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithUsernameGetResponse()
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
func (m *WithUsernameGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *WithUsernameGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *WithUsernameGetResponse) GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameGetResponse) GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *WithUsernameGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WithUsernameGetResponse) SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameGetResponse) SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)() {
    m.publicUser = value
}
// WithUsernameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithUsernameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithUsernameResponse composed type wrapper for classes privateUser, publicUser
type WithUsernameResponse struct {
    // Composed type representation for type privateUser
    privateUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable
    // Composed type representation for type publicUser
    publicUser i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable
}
// NewWithUsernameResponse instantiates a new WithUsernameResponse and sets the default values.
func NewWithUsernameResponse()(*WithUsernameResponse) {
    m := &WithUsernameResponse{
    }
    return m
}
// CreateWithUsernameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithUsernameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithUsernameResponse()
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
func (m *WithUsernameResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *WithUsernameResponse) GetIsComposedType()(bool) {
    return true
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type privateUser
func (m *WithUsernameResponse) GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameResponse) GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *WithUsernameResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WithUsernameResponse) SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type publicUser
func (m *WithUsernameResponse) SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)() {
    m.publicUser = value
}
// WithUsernameGetResponseable 
type WithUsernameGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)
    GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)
    SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)()
    SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)()
}
// WithUsernameResponseable 
type WithUsernameResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)
    GetPublicUser()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)
    SetPrivateUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PrivateUserable)()
    SetPublicUser(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PublicUserable)()
}
// NewWithUsernameItemRequestBuilderInternal instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUsernameItemRequestBuilder) {
    m := &WithUsernameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}", pathParameters),
    }
    return m
}
// NewWithUsernameItemRequestBuilder instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Docker the docker property
func (m *WithUsernameItemRequestBuilder) Docker()(*ItemDockerRequestBuilder) {
    return NewItemDockerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *WithUsernameItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Followers the followers property
func (m *WithUsernameItemRequestBuilder) Followers()(*ItemFollowersRequestBuilder) {
    return NewItemFollowersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Following the following property
func (m *WithUsernameItemRequestBuilder) Following()(*ItemFollowingRequestBuilder) {
    return NewItemFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get provides publicly available information about someone with a GitHub account.GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://docs.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see 'Response with GitHub plan information' below"The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/overview/resources-in-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/users/emails)".
// Deprecated: This method is obsolete. Use GetAsWithUsernameGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-a-user
func (m *WithUsernameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithUsernameItemRequestBuilderGetRequestConfiguration)(WithUsernameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithUsernameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithUsernameResponseable), nil
}
// GetAsWithUsernameGetResponse provides publicly available information about someone with a GitHub account.GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://docs.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see 'Response with GitHub plan information' below"The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/overview/resources-in-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/users/emails)".
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-a-user
func (m *WithUsernameItemRequestBuilder) GetAsWithUsernameGetResponse(ctx context.Context, requestConfiguration *WithUsernameItemRequestBuilderGetRequestConfiguration)(WithUsernameGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithUsernameGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithUsernameGetResponseable), nil
}
// Gists the gists property
func (m *WithUsernameItemRequestBuilder) Gists()(*ItemGistsRequestBuilder) {
    return NewItemGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gpg_keys the gpg_keys property
func (m *WithUsernameItemRequestBuilder) Gpg_keys()(*ItemGpg_keysRequestBuilder) {
    return NewItemGpg_keysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hovercard the hovercard property
func (m *WithUsernameItemRequestBuilder) Hovercard()(*ItemHovercardRequestBuilder) {
    return NewItemHovercardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *WithUsernameItemRequestBuilder) Installation()(*ItemInstallationRequestBuilder) {
    return NewItemInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys the keys property
func (m *WithUsernameItemRequestBuilder) Keys()(*ItemKeysRequestBuilder) {
    return NewItemKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
func (m *WithUsernameItemRequestBuilder) Orgs()(*ItemOrgsRequestBuilder) {
    return NewItemOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Packages the packages property
func (m *WithUsernameItemRequestBuilder) Packages()(*ItemPackagesRequestBuilder) {
    return NewItemPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *WithUsernameItemRequestBuilder) Projects()(*ItemProjectsRequestBuilder) {
    return NewItemProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Received_events the received_events property
func (m *WithUsernameItemRequestBuilder) Received_events()(*ItemReceived_eventsRequestBuilder) {
    return NewItemReceived_eventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *WithUsernameItemRequestBuilder) Repos()(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings the settings property
func (m *WithUsernameItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Social_accounts the social_accounts property
func (m *WithUsernameItemRequestBuilder) Social_accounts()(*ItemSocial_accountsRequestBuilder) {
    return NewItemSocial_accountsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ssh_signing_keys the ssh_signing_keys property
func (m *WithUsernameItemRequestBuilder) Ssh_signing_keys()(*ItemSsh_signing_keysRequestBuilder) {
    return NewItemSsh_signing_keysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Starred the starred property
func (m *WithUsernameItemRequestBuilder) Starred()(*ItemStarredRequestBuilder) {
    return NewItemStarredRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscriptions the subscriptions property
func (m *WithUsernameItemRequestBuilder) Subscriptions()(*ItemSubscriptionsRequestBuilder) {
    return NewItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation provides publicly available information about someone with a GitHub account.GitHub Apps with the `Plan` user permission can use this endpoint to retrieve information about a user's GitHub plan. The GitHub App must be authenticated as a user. See "[Identifying and authorizing users for GitHub Apps](https://docs.github.com/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps/)" for details about authentication. For an example response, see 'Response with GitHub plan information' below"The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/overview/resources-in-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/users/emails)".
func (m *WithUsernameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithUsernameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WithUsernameItemRequestBuilder) WithUrl(rawUrl string)(*WithUsernameItemRequestBuilder) {
    return NewWithUsernameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
