package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCommitsItemCheckSuitesRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\commits\{commits-id}\check-suites
type ItemItemCommitsItemCheckSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.
type ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters struct {
    // Filters check suites by GitHub App `id`.
    App_id *int32 `uriparametername:"app_id"`
    // Returns check runs with the specified `name`.
    Check_name *string `uriparametername:"check_name"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCommitsItemCheckSuitesRequestBuilderGetQueryParameters
}
// NewItemItemCommitsItemCheckSuitesRequestBuilderInternal instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    m := &ItemItemCommitsItemCheckSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/commits/{commits%2Did}/check-suites{?app_id*,check_name*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCommitsItemCheckSuitesRequestBuilder instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCommitsItemCheckSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCommitsItemCheckSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.
// Deprecated: This method is obsolete. Use GetAsCheckSuitesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#list-check-suites-for-a-git-reference
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration)(ItemItemCommitsItemCheckSuitesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckSuitesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckSuitesResponseable), nil
}
// GetAsCheckSuitesGetResponse lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#list-check-suites-for-a-git-reference
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) GetAsCheckSuitesGetResponse(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration)(ItemItemCommitsItemCheckSuitesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemCommitsItemCheckSuitesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemCommitsItemCheckSuitesGetResponseable), nil
}
// ToGetRequestInformation lists check suites for a commit `ref`. The `ref` can be a SHA, branch name, or a tag name.GitHub Apps must have the `checks:read` permission on a private repository or pull access to a public repository to list check suites. OAuth apps and authenticated users must have the `repo` scope to get check suites in a private repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCommitsItemCheckSuitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCommitsItemCheckSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCommitsItemCheckSuitesRequestBuilder) {
    return NewItemItemCommitsItemCheckSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
