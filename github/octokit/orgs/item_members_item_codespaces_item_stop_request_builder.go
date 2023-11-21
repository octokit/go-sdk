package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersItemCodespacesItemStopRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}\codespaces\{codespace_name}\stop
type ItemMembersItemCodespacesItemStopRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersItemCodespacesItemStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemCodespacesItemStopRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMembersItemCodespacesItemStopRequestBuilderInternal instantiates a new StopRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesItemStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesItemStopRequestBuilder) {
    m := &ItemMembersItemCodespacesItemStopRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/members/{username}/codespaces/{codespace_name}/stop", pathParameters),
    }
    return m
}
// NewItemMembersItemCodespacesItemStopRequestBuilder instantiates a new StopRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesItemStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesItemStopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemCodespacesItemStopRequestBuilderInternal(urlParams, requestAdapter)
}
// Post stops a user's codespace.To use this endpoint you must authenticate using one of the following methods:- An access token with the `admin:org` scope- An access token with write permissions for `Codespaces lifecycle admin` on the specific repository and write permissions for `Organization codespaces`
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#stop-a-codespace-for-an-organization-user
func (m *ItemMembersItemCodespacesItemStopRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesItemStopRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "500": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Codespaceable), nil
}
// ToPostRequestInformation stops a user's codespace.To use this endpoint you must authenticate using one of the following methods:- An access token with the `admin:org` scope- An access token with write permissions for `Codespaces lifecycle admin` on the specific repository and write permissions for `Organization codespaces`
func (m *ItemMembersItemCodespacesItemStopRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesItemStopRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMembersItemCodespacesItemStopRequestBuilder) WithUrl(rawUrl string)(*ItemMembersItemCodespacesItemStopRequestBuilder) {
    return NewItemMembersItemCodespacesItemStopRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
