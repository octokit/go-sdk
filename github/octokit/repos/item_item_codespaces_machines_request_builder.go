package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodespacesMachinesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codespaces\machines
type ItemItemCodespacesMachinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodespacesMachinesRequestBuilderGetQueryParameters list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
type ItemItemCodespacesMachinesRequestBuilderGetQueryParameters struct {
    // IP for location auto-detection when proxying a request
    Client_ip *string `uriparametername:"client_ip"`
    // The location to check for available machines. Assigned by IP if not provided.
    Location *string `uriparametername:"location"`
    // The branch or commit to check for prebuild availability and devcontainer restrictions.
    Ref *string `uriparametername:"ref"`
}
// ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodespacesMachinesRequestBuilderGetQueryParameters
}
// NewItemItemCodespacesMachinesRequestBuilderInternal instantiates a new MachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    m := &ItemItemCodespacesMachinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codespaces/machines{?location*,client_ip*,ref*}", pathParameters),
    }
    return m
}
// NewItemItemCodespacesMachinesRequestBuilder instantiates a new MachinesRequestBuilder and sets the default values.
func NewItemItemCodespacesMachinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodespacesMachinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodespacesMachinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsMachinesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/machines#list-available-machine-types-for-a-repository
func (m *ItemItemCodespacesMachinesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration)(ItemItemCodespacesMachinesResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesMachinesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesMachinesResponseable), nil
}
// GetAsMachinesGetResponse list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/machines#list-available-machine-types-for-a-repository
func (m *ItemItemCodespacesMachinesRequestBuilder) GetAsMachinesGetResponse(ctx context.Context, requestConfiguration *ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration)(ItemItemCodespacesMachinesGetResponseable, error) {
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCodespacesMachinesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCodespacesMachinesGetResponseable), nil
}
// ToGetRequestInformation list the machine types available for a given repository based on its configuration.You must authenticate using an access token with the `codespace` scope to use this endpoint.GitHub Apps must have write access to the `codespaces_metadata` repository permission to use this endpoint.
func (m *ItemItemCodespacesMachinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodespacesMachinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodespacesMachinesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodespacesMachinesRequestBuilder) {
    return NewItemItemCodespacesMachinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
