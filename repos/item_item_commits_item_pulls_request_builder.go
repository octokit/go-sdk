package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsItemPullsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\commits\{commit_sha}\pulls
type ItemItemCommitsItemPullsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommitsItemPullsRequestBuilderGetQueryParameters lists the merged pull request that introduced the commit to the repository. If the commit is not present in the default branch, will only return open pull requests associated with the commit.To list the open or merged pull requests associated with a branch, you can set the `commit_sha` parameter to the branch name.
type ItemItemCommitsItemPullsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCommitsItemPullsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommitsItemPullsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommitsItemPullsRequestBuilderGetQueryParameters
}
// NewItemItemCommitsItemPullsRequestBuilderInternal instantiates a new PullsRequestBuilder and sets the default values.
func NewItemItemCommitsItemPullsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemPullsRequestBuilder) {
    m := &ItemItemCommitsItemPullsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/commits/{commit_sha}/pulls{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCommitsItemPullsRequestBuilder instantiates a new PullsRequestBuilder and sets the default values.
func NewItemItemCommitsItemPullsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemPullsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemPullsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the merged pull request that introduced the commit to the repository. If the commit is not present in the default branch, will only return open pull requests associated with the commit.To list the open or merged pull requests associated with a branch, you can set the `commit_sha` parameter to the branch name.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/commits/commits#list-pull-requests-associated-with-a-commit
func (m *ItemItemCommitsItemPullsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommitsItemPullsRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PullRequestSimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreatePullRequestSimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PullRequestSimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.PullRequestSimpleable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the merged pull request that introduced the commit to the repository. If the commit is not present in the default branch, will only return open pull requests associated with the commit.To list the open or merged pull requests associated with a branch, you can set the `commit_sha` parameter to the branch name.
func (m *ItemItemCommitsItemPullsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommitsItemPullsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
func (m *ItemItemCommitsItemPullsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommitsItemPullsRequestBuilder) {
    return NewItemItemCommitsItemPullsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
