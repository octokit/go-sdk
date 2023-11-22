package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemReleasesTagsWithTagItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\releases\tags\{tag}
type ItemItemReleasesTagsWithTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemReleasesTagsWithTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemReleasesTagsWithTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemReleasesTagsWithTagItemRequestBuilderInternal instantiates a new WithTagItemRequestBuilder and sets the default values.
func NewItemItemReleasesTagsWithTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesTagsWithTagItemRequestBuilder) {
    m := &ItemItemReleasesTagsWithTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/releases/tags/{tag}", pathParameters),
    }
    return m
}
// NewItemItemReleasesTagsWithTagItemRequestBuilder instantiates a new WithTagItemRequestBuilder and sets the default values.
func NewItemItemReleasesTagsWithTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemReleasesTagsWithTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemReleasesTagsWithTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a published release with the specified tag.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/releases/releases#get-a-release-by-tag-name
func (m *ItemItemReleasesTagsWithTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemReleasesTagsWithTagItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Releaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateReleaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Releaseable), nil
}
// ToGetRequestInformation get a published release with the specified tag.
func (m *ItemItemReleasesTagsWithTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemReleasesTagsWithTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *ItemItemReleasesTagsWithTagItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemReleasesTagsWithTagItemRequestBuilder) {
    return NewItemItemReleasesTagsWithTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
