package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemMergeUpstreamRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\merge-upstream
type ItemItemMergeUpstreamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemMergeUpstreamRequestBuilderInternal instantiates a new MergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    m := &ItemItemMergeUpstreamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/merge-upstream", pathParameters),
    }
    return m
}
// NewItemItemMergeUpstreamRequestBuilder instantiates a new MergeUpstreamRequestBuilder and sets the default values.
func NewItemItemMergeUpstreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemMergeUpstreamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemMergeUpstreamRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sync a branch of a forked repository to keep it up-to-date with the upstream repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branches#sync-a-fork-branch-with-the-upstream-repository
func (m *ItemItemMergeUpstreamRequestBuilder) Post(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.MergedUpstreamable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateMergedUpstreamFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.MergedUpstreamable), nil
}
// ToPostRequestInformation sync a branch of a forked repository to keep it up-to-date with the upstream repository.
func (m *ItemItemMergeUpstreamRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemMergeUpstreamPostRequestBodyable, requestConfiguration *ItemItemMergeUpstreamRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemMergeUpstreamRequestBuilder) WithUrl(rawUrl string)(*ItemItemMergeUpstreamRequestBuilder) {
    return NewItemItemMergeUpstreamRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
