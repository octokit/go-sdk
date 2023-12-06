package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemBranchesWithBranchItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\branches\{branch}
type ItemItemBranchesWithBranchItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBranchesWithBranchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesWithBranchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemBranchesWithBranchItemRequestBuilderInternal instantiates a new WithBranchItemRequestBuilder and sets the default values.
func NewItemItemBranchesWithBranchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesWithBranchItemRequestBuilder) {
    m := &ItemItemBranchesWithBranchItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/branches/{branch}", pathParameters),
    }
    return m
}
// NewItemItemBranchesWithBranchItemRequestBuilder instantiates a new WithBranchItemRequestBuilder and sets the default values.
func NewItemItemBranchesWithBranchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesWithBranchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesWithBranchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a branch
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branches#get-a-branch
func (m *ItemItemBranchesWithBranchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesWithBranchItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.BranchWithProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBranchWithProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.BranchWithProtectionable), nil
}
// Protection the protection property
func (m *ItemItemBranchesWithBranchItemRequestBuilder) Protection()(*ItemItemBranchesItemProtectionRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rename the rename property
func (m *ItemItemBranchesWithBranchItemRequestBuilder) Rename()(*ItemItemBranchesItemRenameRequestBuilder) {
    return NewItemItemBranchesItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *ItemItemBranchesWithBranchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesWithBranchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemBranchesWithBranchItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesWithBranchItemRequestBuilder) {
    return NewItemItemBranchesWithBranchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
