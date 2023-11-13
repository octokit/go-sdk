package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphSbomRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependency-graph\sbom
type ItemItemDependencyGraphSbomRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependencyGraphSbomRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependencyGraphSbomRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemDependencyGraphSbomRequestBuilderInternal instantiates a new SbomRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSbomRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSbomRequestBuilder) {
    m := &ItemItemDependencyGraphSbomRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/dependency-graph/sbom", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphSbomRequestBuilder instantiates a new SbomRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSbomRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSbomRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphSbomRequestBuilderInternal(urlParams, requestAdapter)
}
// Get exports the software bill of materials (SBOM) for a repository in SPDX JSON format.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependency-graph/sboms#export-a-software-bill-of-materials-sbom-for-a-repository
func (m *ItemItemDependencyGraphSbomRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDependencyGraphSbomRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DependencyGraphSpdxSbomable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateDependencyGraphSpdxSbomFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DependencyGraphSpdxSbomable), nil
}
// ToGetRequestInformation exports the software bill of materials (SBOM) for a repository in SPDX JSON format.
func (m *ItemItemDependencyGraphSbomRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDependencyGraphSbomRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemDependencyGraphSbomRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependencyGraphSbomRequestBuilder) {
    return NewItemItemDependencyGraphSbomRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
