package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInvitationsWithInvitation_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\invitations\{invitation_id}
type ItemInvitationsWithInvitation_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInvitationsWithInvitation_ItemRequestBuilderInternal instantiates a new WithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    m := &ItemInvitationsWithInvitation_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/invitations/{invitation_id}", pathParameters),
    }
    return m
}
// NewItemInvitationsWithInvitation_ItemRequestBuilder instantiates a new WithInvitation_ItemRequestBuilder and sets the default values.
func NewItemInvitationsWithInvitation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#cancel-an-organization-invitation
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Teams the teams property
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) Teams()(*ItemInvitationsItemTeamsRequestBuilder) {
    return NewItemInvitationsItemTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation cancel an organization invitation. In order to cancel an organization invitation, the authenticated user must be an organization owner.This endpoint triggers [notifications](https://docs.github.com/github/managing-subscriptions-and-notifications-on-github/about-notifications).
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInvitationsWithInvitation_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemInvitationsWithInvitation_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInvitationsWithInvitation_ItemRequestBuilder) {
    return NewItemInvitationsWithInvitation_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
