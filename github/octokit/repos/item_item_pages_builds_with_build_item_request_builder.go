package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPagesBuildsWithBuild_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\pages\builds\{build_id}
type ItemItemPagesBuildsWithBuild_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal instantiates a new WithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    m := &ItemItemPagesBuildsWithBuild_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/pages/builds/{build_id}", pathParameters),
    }
    return m
}
// NewItemItemPagesBuildsWithBuild_ItemRequestBuilder instantiates a new WithBuild_ItemRequestBuilder and sets the default values.
func NewItemItemPagesBuildsWithBuild_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPagesBuildsWithBuild_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about a GitHub Pages build.A token with the `repo` scope is required. GitHub Apps must have the `pages:read` permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pages/pages#get-apiname-pages-build
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PageBuildable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePageBuildFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PageBuildable), nil
}
// ToGetRequestInformation gets information about a GitHub Pages build.A token with the `repo` scope is required. GitHub Apps must have the `pages:read` permission.
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPagesBuildsWithBuild_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPagesBuildsWithBuild_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemPagesBuildsWithBuild_ItemRequestBuilder) {
    return NewItemItemPagesBuildsWithBuild_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
