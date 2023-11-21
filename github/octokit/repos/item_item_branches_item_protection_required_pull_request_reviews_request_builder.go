package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\branches\{branch}\protection\required_pull_request_reviews
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderInternal instantiates a new Required_pull_request_reviewsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder instantiates a new Required_pull_request_reviewsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#delete-pull-request-review-protection
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-pull-request-review-protection
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProtectedBranchPullRequestReviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateProtectedBranchPullRequestReviewFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProtectedBranchPullRequestReviewable), nil
}
// Patch protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Updating pull request review enforcement requires admin or owner permissions to the repository and branch protection to be enabled.**Note**: Passing new arrays of `users` and `teams` replaces their previous values.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#update-pull-request-review-protection
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) Patch(ctx context.Context, body ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderPatchRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProtectedBranchPullRequestReviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateProtectedBranchPullRequestReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProtectedBranchPullRequestReviewable), nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Updating pull request review enforcement requires admin or owner permissions to the repository and branch protection to be enabled.**Note**: Passing new arrays of `users` and `teams` replaces their previous values.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemBranchesItemProtectionRequired_pull_request_reviewsPatchRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
