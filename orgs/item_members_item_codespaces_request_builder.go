package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersItemCodespacesRequestBuilder builds and executes requests for operations under \orgs\{org}\members\{username}\codespaces
type ItemMembersItemCodespacesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersItemCodespacesRequestBuilderGetQueryParameters lists the codespaces that a member of an organization has for repositories in that organization.You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces` read permission to use this endpoint.
type ItemMembersItemCodespacesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemMembersItemCodespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersItemCodespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersItemCodespacesRequestBuilderGetQueryParameters
}
// ByCodespace_name gets an item from the go-sdk.orgs.item.members.item.codespaces.item collection
func (m *ItemMembersItemCodespacesRequestBuilder) ByCodespace_name(codespace_name string)(*ItemMembersItemCodespacesWithCodespace_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if codespace_name != "" {
        urlTplParams["codespace_name"] = codespace_name
    }
    return NewItemMembersItemCodespacesWithCodespace_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMembersItemCodespacesRequestBuilderInternal instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesRequestBuilder) {
    m := &ItemMembersItemCodespacesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/members/{username}/codespaces{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemMembersItemCodespacesRequestBuilder instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemMembersItemCodespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersItemCodespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersItemCodespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the codespaces that a member of an organization has for repositories in that organization.You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces` read permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsCodespacesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#list-codespaces-for-a-user-in-organization
func (m *ItemMembersItemCodespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesRequestBuilderGetRequestConfiguration)(ItemMembersItemCodespacesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMembersItemCodespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMembersItemCodespacesResponseable), nil
}
// GetAsCodespacesGetResponse lists the codespaces that a member of an organization has for repositories in that organization.You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces` read permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/organizations#list-codespaces-for-a-user-in-organization
func (m *ItemMembersItemCodespacesRequestBuilder) GetAsCodespacesGetResponse(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesRequestBuilderGetRequestConfiguration)(ItemMembersItemCodespacesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemMembersItemCodespacesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMembersItemCodespacesGetResponseable), nil
}
// ToGetRequestInformation lists the codespaces that a member of an organization has for repositories in that organization.You must authenticate using an access token with the `admin:org` scope or the `Organization codespaces` read permission to use this endpoint.
func (m *ItemMembersItemCodespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersItemCodespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMembersItemCodespacesRequestBuilder) WithUrl(rawUrl string)(*ItemMembersItemCodespacesRequestBuilder) {
    return NewItemMembersItemCodespacesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
