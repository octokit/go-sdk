package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemGitBlobsWithFile_shaItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\git\blobs\{file_sha}
type ItemItemGitBlobsWithFile_shaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemGitBlobsWithFile_shaItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemGitBlobsWithFile_shaItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemGitBlobsWithFile_shaItemRequestBuilderInternal instantiates a new WithFile_shaItemRequestBuilder and sets the default values.
func NewItemItemGitBlobsWithFile_shaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitBlobsWithFile_shaItemRequestBuilder) {
    m := &ItemItemGitBlobsWithFile_shaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/git/blobs/{file_sha}", pathParameters),
    }
    return m
}
// NewItemItemGitBlobsWithFile_shaItemRequestBuilder instantiates a new WithFile_shaItemRequestBuilder and sets the default values.
func NewItemItemGitBlobsWithFile_shaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemGitBlobsWithFile_shaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemGitBlobsWithFile_shaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the `content` in the response will always be Base64 encoded._Note_: This API supports blobs up to 100 megabytes in size.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/git/blobs#get-a-blob
func (m *ItemItemGitBlobsWithFile_shaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemGitBlobsWithFile_shaItemRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Blobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBlobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Blobable), nil
}
// ToGetRequestInformation the `content` in the response will always be Base64 encoded._Note_: This API supports blobs up to 100 megabytes in size.
func (m *ItemItemGitBlobsWithFile_shaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemGitBlobsWithFile_shaItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemGitBlobsWithFile_shaItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemGitBlobsWithFile_shaItemRequestBuilder) {
    return NewItemItemGitBlobsWithFile_shaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
