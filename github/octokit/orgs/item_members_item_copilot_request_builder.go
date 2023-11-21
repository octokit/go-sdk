package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersItemCopilotRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}\copilot
type ItemMembersItemCopilotRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersItemCopilotRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemCopilotRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMembersItemCopilotRequestBuilderInternal instantiates a new CopilotRequestBuilder and sets the default values.
func NewItemMembersItemCopilotRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCopilotRequestBuilder) {
    m := &ItemMembersItemCopilotRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/members/{username}/copilot", pathParameters),
    }
    return m
}
// NewItemMembersItemCopilotRequestBuilder instantiates a new CopilotRequestBuilder and sets the default values.
func NewItemMembersItemCopilotRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCopilotRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemCopilotRequestBuilderInternal(urlParams, requestAdapter)
}
// Get **Note**: This endpoint is in beta and is subject to change.Gets the GitHub Copilot Business seat assignment details for a member of an organization who currently has access to GitHub Copilot.Organization owners and members with admin permissions can view GitHub Copilot seat assignment details for members in their organization. You must authenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-business#get-copilot-business-seat-assignment-details-for-a-user
func (m *ItemMembersItemCopilotRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersItemCopilotRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCopilotSeatDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable), nil
}
// ToGetRequestInformation **Note**: This endpoint is in beta and is subject to change.Gets the GitHub Copilot Business seat assignment details for a member of an organization who currently has access to GitHub Copilot.Organization owners and members with admin permissions can view GitHub Copilot seat assignment details for members in their organization. You must authenticate using an access token with the `manage_billing:copilot` scope to use this endpoint.
func (m *ItemMembersItemCopilotRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemCopilotRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMembersItemCopilotRequestBuilder) WithUrl(rawUrl string)(*ItemMembersItemCopilotRequestBuilder) {
    return NewItemMembersItemCopilotRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
