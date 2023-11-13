package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodespacesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces
type ItemItemCodespacesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesRequestBuilderGetQueryParameters lists the codespaces associated to a specified repository and the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
type ItemItemCodespacesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCodespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesRequestBuilderGetQueryParameters
}
// ItemItemCodespacesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCodespacesRequestBuilderInternal instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemItemCodespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesRequestBuilder) {
    m := &ItemItemCodespacesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codespaces{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesRequestBuilder instantiates a new CodespacesRequestBuilder and sets the default values.
func NewItemItemCodespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Devcontainers the devcontainers property
func (m *ItemItemCodespacesRequestBuilder) Devcontainers()(*ItemItemCodespacesDevcontainersRequestBuilder) {
    return NewItemItemCodespacesDevcontainersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get lists the codespaces associated to a specified repository and the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsCodespacesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-codespaces-in-a-repository-for-the-authenticated-user
func (m *ItemItemCodespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesRequestBuilderGetRequestConfiguration)(ItemItemCodespacesResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesResponseable), nil
}
// GetAsCodespacesGetResponse lists the codespaces associated to a specified repository and the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-codespaces-in-a-repository-for-the-authenticated-user
func (m *ItemItemCodespacesRequestBuilder) GetAsCodespacesGetResponse(ctx context.Context, requestConfiguration *ItemItemCodespacesRequestBuilderGetRequestConfiguration)(ItemItemCodespacesGetResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesGetResponseable), nil
}
// Machines the machines property
func (m *ItemItemCodespacesRequestBuilder) Machines()(*ItemItemCodespacesMachinesRequestBuilder) {
    return NewItemItemCodespacesMachinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// New the new property
func (m *ItemItemCodespacesRequestBuilder) New()(*ItemItemCodespacesNewRequestBuilder) {
    return NewItemItemCodespacesNewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions_check the permissions_check property
func (m *ItemItemCodespacesRequestBuilder) Permissions_check()(*ItemItemCodespacesPermissions_checkRequestBuilder) {
    return NewItemItemCodespacesPermissions_checkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates a codespace owned by the authenticated user in the specified repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#create-a-codespace-in-a-repository
func (m *ItemItemCodespacesRequestBuilder) Post(ctx context.Context, body ItemItemCodespacesPostRequestBodyable, requestConfiguration *ItemItemCodespacesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "503": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodespace503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Codespaceable), nil
}
// Secrets the secrets property
func (m *ItemItemCodespacesRequestBuilder) Secrets()(*ItemItemCodespacesSecretsRequestBuilder) {
    return NewItemItemCodespacesSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists the codespaces associated to a specified repository and the authenticated user.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces` repository permission to use this endpoint.
func (m *ItemItemCodespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a codespace owned by the authenticated user in the specified repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces` repository permission to use this endpoint.
func (m *ItemItemCodespacesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemCodespacesPostRequestBodyable, requestConfiguration *ItemItemCodespacesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCodespacesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesRequestBuilder) {
    return NewItemItemCodespacesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
