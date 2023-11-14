package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBranchesWithBranchItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/branches/{branch}", pathParameters),
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
func (m *ItemItemBranchesWithBranchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesWithBranchItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.BranchWithProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBranchWithProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.BranchWithProtectionable), nil
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
func (m *ItemItemBranchesWithBranchItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesWithBranchItemRequestBuilder) {
    return NewItemItemBranchesWithBranchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
