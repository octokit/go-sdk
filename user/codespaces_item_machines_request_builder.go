package user

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesItemMachinesRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\machines
type CodespacesItemMachinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesItemMachinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemMachinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesItemMachinesRequestBuilderInternal instantiates a new MachinesRequestBuilder and sets the default values.
func NewCodespacesItemMachinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemMachinesRequestBuilder) {
    m := &CodespacesItemMachinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/machines", pathParameters),
    }
    return m
}
// NewCodespacesItemMachinesRequestBuilder instantiates a new MachinesRequestBuilder and sets the default values.
func NewCodespacesItemMachinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemMachinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemMachinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the machine types a codespace can transition to use.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsMachinesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/machines#list-machine-types-for-a-codespace
func (m *CodespacesItemMachinesRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesItemMachinesRequestBuilderGetRequestConfiguration)(CodespacesItemMachinesResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesItemMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesItemMachinesResponseable), nil
}
// GetAsMachinesGetResponse list the machine types a codespace can transition to use.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/machines#list-machine-types-for-a-codespace
func (m *CodespacesItemMachinesRequestBuilder) GetAsMachinesGetResponse(ctx context.Context, requestConfiguration *CodespacesItemMachinesRequestBuilderGetRequestConfiguration)(CodespacesItemMachinesGetResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesItemMachinesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesItemMachinesGetResponseable), nil
}
// ToGetRequestInformation list the machine types a codespace can transition to use.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_metadata` repository permission to use this endpoint.
func (m *CodespacesItemMachinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesItemMachinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *CodespacesItemMachinesRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemMachinesRequestBuilder) {
    return NewCodespacesItemMachinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
