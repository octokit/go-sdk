package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodespacesDevcontainersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces\devcontainers
type ItemItemCodespacesDevcontainersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesDevcontainersRequestBuilderGetQueryParameters lists the devcontainer.json files associated with a specified repository and the authenticated user. These filesspecify launchpoint configurations for codespaces created within the repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
type ItemItemCodespacesDevcontainersRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCodespacesDevcontainersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesDevcontainersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesDevcontainersRequestBuilderGetQueryParameters
}
// NewItemItemCodespacesDevcontainersRequestBuilderInternal instantiates a new DevcontainersRequestBuilder and sets the default values.
func NewItemItemCodespacesDevcontainersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesDevcontainersRequestBuilder) {
    m := &ItemItemCodespacesDevcontainersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codespaces/devcontainers{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesDevcontainersRequestBuilder instantiates a new DevcontainersRequestBuilder and sets the default values.
func NewItemItemCodespacesDevcontainersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesDevcontainersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesDevcontainersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the devcontainer.json files associated with a specified repository and the authenticated user. These filesspecify launchpoint configurations for codespaces created within the repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsDevcontainersGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-devcontainer-configurations-in-a-repository-for-the-authenticated-user
func (m *ItemItemCodespacesDevcontainersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesDevcontainersRequestBuilderGetRequestConfiguration)(ItemItemCodespacesDevcontainersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesDevcontainersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesDevcontainersResponseable), nil
}
// GetAsDevcontainersGetResponse lists the devcontainer.json files associated with a specified repository and the authenticated user. These filesspecify launchpoint configurations for codespaces created within the repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-devcontainer-configurations-in-a-repository-for-the-authenticated-user
func (m *ItemItemCodespacesDevcontainersRequestBuilder) GetAsDevcontainersGetResponse(ctx context.Context, requestConfiguration *ItemItemCodespacesDevcontainersRequestBuilderGetRequestConfiguration)(ItemItemCodespacesDevcontainersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "401": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesDevcontainersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesDevcontainersGetResponseable), nil
}
// ToGetRequestInformation lists the devcontainer.json files associated with a specified repository and the authenticated user. These filesspecify launchpoint configurations for codespaces created within the repository.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
func (m *ItemItemCodespacesDevcontainersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesDevcontainersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodespacesDevcontainersRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesDevcontainersRequestBuilder) {
    return NewItemItemCodespacesDevcontainersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
