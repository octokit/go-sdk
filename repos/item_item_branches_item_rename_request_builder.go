package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBranchesItemRenameRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\rename
type ItemItemBranchesItemRenameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBranchesItemRenameRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemRenameRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemBranchesItemRenameRequestBuilderInternal instantiates a new RenameRequestBuilder and sets the default values.
func NewItemItemBranchesItemRenameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemRenameRequestBuilder) {
    m := &ItemItemBranchesItemRenameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/rename", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemRenameRequestBuilder instantiates a new RenameRequestBuilder and sets the default values.
func NewItemItemBranchesItemRenameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemRenameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemRenameRequestBuilderInternal(urlParams, requestAdapter)
}
// Post renames a branch in a repository.**Note:** Although the API responds immediately, the branch rename process might take some extra time to complete in the background. You won't be able to push to the old branch name while the rename process is in progress. For more information, see "[Renaming a branch](https://docs.github.com/github/administering-a-repository/renaming-a-branch)".The permissions required to use this endpoint depends on whether you are renaming the default branch.To rename a non-default branch:* Users must have push access.* GitHub Apps must have the `contents:write` repository permission.To rename the default branch:* Users must have admin or owner permissions.* GitHub Apps must have the `administration:write` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branches#rename-a-branch
func (m *ItemItemBranchesItemRenameRequestBuilder) Post(ctx context.Context, body ItemItemBranchesItemRenamePostRequestBodyable, requestConfiguration *ItemItemBranchesItemRenameRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.BranchWithProtectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBranchWithProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.BranchWithProtectionable), nil
}
// ToPostRequestInformation renames a branch in a repository.**Note:** Although the API responds immediately, the branch rename process might take some extra time to complete in the background. You won't be able to push to the old branch name while the rename process is in progress. For more information, see "[Renaming a branch](https://docs.github.com/github/administering-a-repository/renaming-a-branch)".The permissions required to use this endpoint depends on whether you are renaming the default branch.To rename a non-default branch:* Users must have push access.* GitHub Apps must have the `contents:write` repository permission.To rename the default branch:* Users must have admin or owner permissions.* GitHub Apps must have the `administration:write` repository permission.
func (m *ItemItemBranchesItemRenameRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemBranchesItemRenamePostRequestBodyable, requestConfiguration *ItemItemBranchesItemRenameRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemBranchesItemRenameRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemRenameRequestBuilder) {
    return NewItemItemBranchesItemRenameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
