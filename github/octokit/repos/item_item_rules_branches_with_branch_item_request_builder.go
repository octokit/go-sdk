package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemRulesBranchesWithBranchItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\rules\branches\{branch}
type ItemItemRulesBranchesWithBranchItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
type ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemRulesBranchesWithBranchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemRulesBranchesWithBranchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemRulesBranchesWithBranchItemRequestBuilderGetQueryParameters
}
// NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal instantiates a new WithBranchItemRequestBuilder and sets the default values.
func NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    m := &ItemItemRulesBranchesWithBranchItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/rules/branches/{branch}{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemRulesBranchesWithBranchItemRequestBuilder instantiates a new WithBranchItemRequestBuilder and sets the default values.
func NewItemItemRulesBranchesWithBranchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemRulesBranchesWithBranchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/rules#get-rules-for-a-branch
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemRulesBranchesWithBranchItemRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleDetailedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRepositoryRuleDetailedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleDetailedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.RepositoryRuleDetailedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns all active rules that apply to the specified branch. The branch does not need to exist; rules that would applyto a branch with that name will be returned. All active rules that apply will be returned, regardless of the levelat which they are configured (e.g. repository or organization). Rules in rulesets with "evaluate" or "disabled"enforcement statuses are not returned.
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemRulesBranchesWithBranchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemRulesBranchesWithBranchItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemRulesBranchesWithBranchItemRequestBuilder) {
    return NewItemItemRulesBranchesWithBranchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
