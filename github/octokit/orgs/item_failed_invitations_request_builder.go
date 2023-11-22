package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemFailed_invitationsRequestBuilder builds and executes requests for operations under \orgs\{org}\failed_invitations
type ItemFailed_invitationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFailed_invitationsRequestBuilderGetQueryParameters the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.
type ItemFailed_invitationsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemFailed_invitationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFailed_invitationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFailed_invitationsRequestBuilderGetQueryParameters
}
// NewItemFailed_invitationsRequestBuilderInternal instantiates a new Failed_invitationsRequestBuilder and sets the default values.
func NewItemFailed_invitationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFailed_invitationsRequestBuilder) {
    m := &ItemFailed_invitationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/failed_invitations{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemFailed_invitationsRequestBuilder instantiates a new Failed_invitationsRequestBuilder and sets the default values.
func NewItemFailed_invitationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFailed_invitationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFailed_invitationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#list-failed-organization-invitations
func (m *ItemFailed_invitationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFailed_invitationsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationInvitationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationInvitationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationInvitationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationInvitationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation the return hash contains `failed_at` and `failed_reason` fields which represent the time at which the invitation failed and the reason for the failure.
func (m *ItemFailed_invitationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFailed_invitationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
func (m *ItemFailed_invitationsRequestBuilder) WithUrl(rawUrl string)(*ItemFailed_invitationsRequestBuilder) {
    return NewItemFailed_invitationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
