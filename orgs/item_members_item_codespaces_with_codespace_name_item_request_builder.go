package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}\codespaces\{codespace_name}
type ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    m := &ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/members/{username}/codespaces/{codespace_name}", pathParameters),
    }
    return m
}
// NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder instantiates a new WithCodespace_nameItemRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a user's codespace.To use this endpoint you must authenticate using one of the following methods:- An access token with the `admin:org` scope- An access token with write permissions for `Codespaces` on the specific repository and write permissions for `Organization codespaces`
// Deprecated: This method is obsolete. Use DeleteAsWithCodespace_nameDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#delete-a-codespace-from-the-organization
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(ItemMembersItemCodespacesItemWithCodespace_nameResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMembersItemCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMembersItemCodespacesItemWithCodespace_nameResponseable), nil
}
// DeleteAsWithCodespace_nameDeleteResponse deletes a user's codespace.To use this endpoint you must authenticate using one of the following methods:- An access token with the `admin:org` scope- An access token with write permissions for `Codespaces` on the specific repository and write permissions for `Organization codespaces`
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#delete-a-codespace-from-the-organization
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) DeleteAsWithCodespace_nameDeleteResponse(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(ItemMembersItemCodespacesItemWithCodespace_nameDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMembersItemCodespacesItemWithCodespace_nameDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMembersItemCodespacesItemWithCodespace_nameDeleteResponseable), nil
}
// Stop the stop property
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) Stop()(*ItemMembersItemCodespacesItemStopRequestBuilder) {
    return NewItemMembersItemCodespacesItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation deletes a user's codespace.To use this endpoint you must authenticate using one of the following methods:- An access token with the `admin:org` scope- An access token with write permissions for `Codespaces` on the specific repository and write permissions for `Organization codespaces`
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    return NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
