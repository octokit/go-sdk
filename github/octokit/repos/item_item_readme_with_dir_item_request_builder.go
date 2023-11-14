package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemReadmeWithDirItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\readme\{dir}
type ItemItemReadmeWithDirItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters gets the README from a repository directory.READMEs support [custom media types](https://docs.github.com/rest/overview/media-types) for retrieving the raw content or rendered HTML.
type ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters struct {
    // The name of the commit/branch/tag. Default: the repository’s default branch.
    Ref *string `uriparametername:"ref"`
}
// ItemItemReadmeWithDirItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemReadmeWithDirItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemReadmeWithDirItemRequestBuilderGetQueryParameters
}
// NewItemItemReadmeWithDirItemRequestBuilderInternal instantiates a new WithDirItemRequestBuilder and sets the default values.
func NewItemItemReadmeWithDirItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReadmeWithDirItemRequestBuilder) {
    m := &ItemItemReadmeWithDirItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/readme/{dir}{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemReadmeWithDirItemRequestBuilder instantiates a new WithDirItemRequestBuilder and sets the default values.
func NewItemItemReadmeWithDirItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReadmeWithDirItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReadmeWithDirItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the README from a repository directory.READMEs support [custom media types](https://docs.github.com/rest/overview/media-types) for retrieving the raw content or rendered HTML.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#get-a-repository-readme-for-a-directory
func (m *ItemItemReadmeWithDirItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemReadmeWithDirItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ContentFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ContentFileable), nil
}
// ToGetRequestInformation gets the README from a repository directory.READMEs support [custom media types](https://docs.github.com/rest/overview/media-types) for retrieving the raw content or rendered HTML.
func (m *ItemItemReadmeWithDirItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemReadmeWithDirItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemReadmeWithDirItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemReadmeWithDirItemRequestBuilder) {
    return NewItemItemReadmeWithDirItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
