package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemDependencyGraphSnapshotsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\dependency-graph\snapshots
type ItemItemDependencyGraphSnapshotsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependencyGraphSnapshotsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemDependencyGraphSnapshotsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemDependencyGraphSnapshotsRequestBuilderInternal instantiates a new SnapshotsRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    m := &ItemItemDependencyGraphSnapshotsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/dependency-graph/snapshots", pathParameters),
    }
    return m
}
// NewItemItemDependencyGraphSnapshotsRequestBuilder instantiates a new SnapshotsRequestBuilder and sets the default values.
func NewItemItemDependencyGraphSnapshotsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependencyGraphSnapshotsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new snapshot of a repository's dependencies. You must authenticate using an access token with the `repo` scope to use this endpoint for a repository that the requesting user has access to.
// Deprecated: This method is obsolete. Use PostAsSnapshotsPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependency-graph/dependency-submission#create-a-snapshot-of-dependencies-for-a-repository
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) Post(ctx context.Context, body i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Snapshotable, requestConfiguration *ItemItemDependencyGraphSnapshotsRequestBuilderPostRequestConfiguration)(ItemItemDependencyGraphSnapshotsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemDependencyGraphSnapshotsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemDependencyGraphSnapshotsResponseable), nil
}
// PostAsSnapshotsPostResponse create a new snapshot of a repository's dependencies. You must authenticate using an access token with the `repo` scope to use this endpoint for a repository that the requesting user has access to.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependency-graph/dependency-submission#create-a-snapshot-of-dependencies-for-a-repository
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) PostAsSnapshotsPostResponse(ctx context.Context, body i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Snapshotable, requestConfiguration *ItemItemDependencyGraphSnapshotsRequestBuilderPostRequestConfiguration)(ItemItemDependencyGraphSnapshotsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemDependencyGraphSnapshotsPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemDependencyGraphSnapshotsPostResponseable), nil
}
// ToPostRequestInformation create a new snapshot of a repository's dependencies. You must authenticate using an access token with the `repo` scope to use this endpoint for a repository that the requesting user has access to.
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Snapshotable, requestConfiguration *ItemItemDependencyGraphSnapshotsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemDependencyGraphSnapshotsRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependencyGraphSnapshotsRequestBuilder) {
    return NewItemItemDependencyGraphSnapshotsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
