package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemPullsWithPull_numberItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\pulls\{pull_number}
type ItemItemPullsWithPull_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPullsWithPull_numberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsWithPull_numberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemPullsWithPull_numberItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPullsWithPull_numberItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Codespaces the codespaces property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Codespaces()(*ItemItemPullsItemCodespacesRequestBuilder) {
    return NewItemItemPullsItemCodespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Comments the comments property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Comments()(*ItemItemPullsItemCommentsRequestBuilder) {
    return NewItemItemPullsItemCommentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commits the commits property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Commits()(*ItemItemPullsItemCommitsRequestBuilder) {
    return NewItemItemPullsItemCommitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemPullsWithPull_numberItemRequestBuilderInternal instantiates a new WithPull_numberItemRequestBuilder and sets the default values.
func NewItemItemPullsWithPull_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsWithPull_numberItemRequestBuilder) {
    m := &ItemItemPullsWithPull_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/pulls/{pull_number}", pathParameters),
    }
    return m
}
// NewItemItemPullsWithPull_numberItemRequestBuilder instantiates a new WithPull_numberItemRequestBuilder and sets the default values.
func NewItemItemPullsWithPull_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPullsWithPull_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPullsWithPull_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Files the files property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Files()(*ItemItemPullsItemFilesRequestBuilder) {
    return NewItemItemPullsItemFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists details of a pull request by providing its number.When you get, [create](https://docs.github.com/rest/pulls/pulls/#create-a-pull-request), or [edit](https://docs.github.com/rest/pulls/pulls#update-a-pull-request) a pull request, GitHub creates a merge commit to test whether the pull request can be automatically merged into the base branch. This test commit is not added to the base branch or the head branch. You can review the status of the test commit using the `mergeable` key. For more information, see "[Checking mergeability of pull requests](https://docs.github.com/rest/guides/getting-started-with-the-git-database-api#checking-mergeability-of-pull-requests)".The value of the `mergeable` attribute can be `true`, `false`, or `null`. If the value is `null`, then GitHub has started a background job to compute the mergeability. After giving the job time to complete, resubmit the request. When the job finishes, you will see a non-`null` value for the `mergeable` attribute in the response. If `mergeable` is `true`, then `merge_commit_sha` will be the SHA of the _test_ merge commit.The value of the `merge_commit_sha` attribute changes depending on the state of the pull request. Before merging a pull request, the `merge_commit_sha` attribute holds the SHA of the _test_ merge commit. After merging a pull request, the `merge_commit_sha` attribute changes depending on how you merged the pull request:*   If merged as a [merge commit](https://docs.github.com/articles/about-merge-methods-on-github/), `merge_commit_sha` represents the SHA of the merge commit.*   If merged via a [squash](https://docs.github.com/articles/about-merge-methods-on-github/#squashing-your-merge-commits), `merge_commit_sha` represents the SHA of the squashed commit on the base branch.*   If [rebased](https://docs.github.com/articles/about-merge-methods-on-github/#rebasing-and-merging-your-commits), `merge_commit_sha` represents the commit that the base branch was updated to.Pass the appropriate [media type](https://docs.github.com/rest/overview/media-types/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/pulls#get-a-pull-request
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPullsWithPull_numberItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.PullRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "503": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreatePullRequest503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreatePullRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.PullRequestable), nil
}
// Merge the merge property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Merge()(*ItemItemPullsItemMergeRequestBuilder) {
    return NewItemItemPullsItemMergeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.To open or update a pull request in a public repository, you must have write access to the head or the source branch. For organization-owned repositories, you must be a member of the organization that owns the repository to open or update a pull request.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/pulls/pulls#update-a-pull-request
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Patch(ctx context.Context, body ItemItemPullsItemWithPull_numberPatchRequestBodyable, requestConfiguration *ItemItemPullsWithPull_numberItemRequestBuilderPatchRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.PullRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreatePullRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.PullRequestable), nil
}
// Requested_reviewers the requested_reviewers property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Requested_reviewers()(*ItemItemPullsItemRequested_reviewersRequestBuilder) {
    return NewItemItemPullsItemRequested_reviewersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reviews the reviews property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) Reviews()(*ItemItemPullsItemReviewsRequestBuilder) {
    return NewItemItemPullsItemReviewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists details of a pull request by providing its number.When you get, [create](https://docs.github.com/rest/pulls/pulls/#create-a-pull-request), or [edit](https://docs.github.com/rest/pulls/pulls#update-a-pull-request) a pull request, GitHub creates a merge commit to test whether the pull request can be automatically merged into the base branch. This test commit is not added to the base branch or the head branch. You can review the status of the test commit using the `mergeable` key. For more information, see "[Checking mergeability of pull requests](https://docs.github.com/rest/guides/getting-started-with-the-git-database-api#checking-mergeability-of-pull-requests)".The value of the `mergeable` attribute can be `true`, `false`, or `null`. If the value is `null`, then GitHub has started a background job to compute the mergeability. After giving the job time to complete, resubmit the request. When the job finishes, you will see a non-`null` value for the `mergeable` attribute in the response. If `mergeable` is `true`, then `merge_commit_sha` will be the SHA of the _test_ merge commit.The value of the `merge_commit_sha` attribute changes depending on the state of the pull request. Before merging a pull request, the `merge_commit_sha` attribute holds the SHA of the _test_ merge commit. After merging a pull request, the `merge_commit_sha` attribute changes depending on how you merged the pull request:*   If merged as a [merge commit](https://docs.github.com/articles/about-merge-methods-on-github/), `merge_commit_sha` represents the SHA of the merge commit.*   If merged via a [squash](https://docs.github.com/articles/about-merge-methods-on-github/#squashing-your-merge-commits), `merge_commit_sha` represents the SHA of the squashed commit on the base branch.*   If [rebased](https://docs.github.com/articles/about-merge-methods-on-github/#rebasing-and-merging-your-commits), `merge_commit_sha` represents the commit that the base branch was updated to.Pass the appropriate [media type](https://docs.github.com/rest/overview/media-types/#commits-commit-comparison-and-pull-requests) to fetch diff and patch formats.
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPullsWithPull_numberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation draft pull requests are available in public repositories with GitHub Free and GitHub Free for organizations, GitHub Pro, and legacy per-repository billing plans, and in public and private repositories with GitHub Team and GitHub Enterprise Cloud. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.To open or update a pull request in a public repository, you must have write access to the head or the source branch. For organization-owned repositories, you must be a member of the organization that owns the repository to open or update a pull request.
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemPullsItemWithPull_numberPatchRequestBodyable, requestConfiguration *ItemItemPullsWithPull_numberItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UpdateBranch the updateBranch property
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) UpdateBranch()(*ItemItemPullsItemUpdateBranchRequestBuilder) {
    return NewItemItemPullsItemUpdateBranchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemPullsWithPull_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemPullsWithPull_numberItemRequestBuilder) {
    return NewItemItemPullsWithPull_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
