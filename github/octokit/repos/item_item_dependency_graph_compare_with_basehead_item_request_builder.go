package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\dependency-graph\compare\{basehead}
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters struct {
    // The full path, relative to the repository root, of the dependency manifest file.
    Name *string `uriparametername:"name"`
}
// ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetQueryParameters
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    m := &ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/dependency-graph/compare/{basehead}{?name*}", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder instantiates a new WithBaseheadItemRequestBuilder and sets the default values.
func NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependency-graph/dependency-review#get-a-diff-of-the-dependencies-between-commits
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DependencyGraphDiffable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateDependencyGraphDiffFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DependencyGraphDiffable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.DependencyGraphDiffable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets the diff of the dependency changes between two commits of a repository, based on the changes to the dependency manifests made in those commits.
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder) {
    return NewItemItemDependencyGraphCompareWithBaseheadItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
