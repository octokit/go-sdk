package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodespacesPermissions_checkRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces\permissions_check
type ItemItemCodespacesPermissions_checkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
type ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters struct {
    // Path to the devcontainer.json configuration to use for the permission check.
    Devcontainer_path *string `uriparametername:"devcontainer_path"`
    // The git reference that points to the location of the devcontainer configuration to use for the permission check. The value of `ref` will typically be a branch name (`heads/BRANCH_NAME`). For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
    Ref *string `uriparametername:"ref"`
}
// ItemItemCodespacesPermissions_checkRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesPermissions_checkRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesPermissions_checkRequestBuilderGetQueryParameters
}
// NewItemItemCodespacesPermissions_checkRequestBuilderInternal instantiates a new Permissions_checkRequestBuilder and sets the default values.
func NewItemItemCodespacesPermissions_checkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    m := &ItemItemCodespacesPermissions_checkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codespaces/permissions_check{?ref*,devcontainer_path*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesPermissions_checkRequestBuilder instantiates a new Permissions_checkRequestBuilder and sets the default values.
func NewItemItemCodespacesPermissions_checkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesPermissions_checkRequestBuilderInternal(urlParams, requestAdapter)
}
// Get checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#check-if-permissions-defined-by-a-devcontainer-have-been-accepted-by-the-authenticated-user
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesPermissions_checkRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesPermissionsCheckForDevcontainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespacesPermissionsCheckForDevcontainer503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespacesPermissionsCheckForDevcontainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespacesPermissionsCheckForDevcontainerable), nil
}
// ToGetRequestInformation checks whether the permissions defined by a given devcontainer configuration have been accepted by the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesPermissions_checkRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodespacesPermissions_checkRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    return NewItemItemCodespacesPermissions_checkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
