package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemPullsItemUpdateBranchRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\{pull_number}\update-branch
type ItemItemPullsItemUpdateBranchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsItemUpdateBranchRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsItemUpdateBranchRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPullsItemUpdateBranchRequestBuilderInternal instantiates a new UpdateBranchRequestBuilder and sets the default values.
func NewItemItemPullsItemUpdateBranchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    m := &ItemItemPullsItemUpdateBranchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/{pull_number}/update-branch", pathParameters),
    }
    return m
}
// NewItemItemPullsItemUpdateBranchRequestBuilder instantiates a new UpdateBranchRequestBuilder and sets the default values.
func NewItemItemPullsItemUpdateBranchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsItemUpdateBranchRequestBuilderInternal(urlParams, requestAdapter)
}
// Put updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.
// Deprecated: This method is obsolete. Use PutAsUpdateBranchPutResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/pulls#update-a-pull-request-branch
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) Put(ctx context.Context, body ItemItemPullsItemUpdateBranchPutRequestBodyable, requestConfiguration *ItemItemPullsItemUpdateBranchRequestBuilderPutRequestConfiguration)(ItemItemPullsItemUpdateBranchResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemPullsItemUpdateBranchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemPullsItemUpdateBranchResponseable), nil
}
// PutAsUpdateBranchPutResponse updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/pulls#update-a-pull-request-branch
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) PutAsUpdateBranchPutResponse(ctx context.Context, body ItemItemPullsItemUpdateBranchPutRequestBodyable, requestConfiguration *ItemItemPullsItemUpdateBranchRequestBuilderPutRequestConfiguration)(ItemItemPullsItemUpdateBranchPutResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemPullsItemUpdateBranchPutResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemPullsItemUpdateBranchPutResponseable), nil
}
// ToPutRequestInformation updates the pull request branch with the latest upstream changes by merging HEAD from the base branch into the pull request branch.
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemPullsItemUpdateBranchPutRequestBodyable, requestConfiguration *ItemItemPullsItemUpdateBranchRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *ItemItemPullsItemUpdateBranchRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    return NewItemItemPullsItemUpdateBranchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
