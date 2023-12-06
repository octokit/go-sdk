package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie1c6ee8184c8a0498f517f5358d5b5276a5f94b8d80136352d0b23c8a035838c "github.com/octokit/go-sdk/github/octokit/repos/item/item/actions/runs"
)

// ItemItemActionsRunsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\runs
type ItemItemActionsRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsRequestBuilderGetQueryParameters lists all workflow runs for a repository. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemActionsRunsRequestBuilderGetQueryParameters struct {
    // Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run.
    Actor *string `uriparametername:"actor"`
    // Returns workflow runs associated with a branch. Use the name of the branch of the `push`.
    Branch *string `uriparametername:"branch"`
    // Returns workflow runs with the `check_suite_id` that you specify.
    Check_suite_id *int32 `uriparametername:"check_suite_id"`
    // Returns workflow runs created within the given date-time range. For more information on the syntax, see "[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates)."
    Created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"created"`
    // Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see "[Events that trigger workflows](https://docs.github.com/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows)."
    Event *string `uriparametername:"event"`
    // If `true` pull requests are omitted from the response (empty array).
    Exclude_pull_requests *bool `uriparametername:"exclude_pull_requests"`
    // Only returns workflow runs that are associated with the specified `head_sha`.
    Head_sha *string `uriparametername:"head_sha"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub can set a status of `waiting` or `requested`.
    // Deprecated: This property is deprecated, use statusAsGetStatusQueryParameterType instead
    Status *string `uriparametername:"status"`
    // Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub can set a status of `waiting` or `requested`.
    StatusAsGetStatusQueryParameterType *ie1c6ee8184c8a0498f517f5358d5b5276a5f94b8d80136352d0b23c8a035838c.GetStatusQueryParameterType `uriparametername:"status"`
}
// ItemItemActionsRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsRequestBuilderGetQueryParameters
}
// ByRun_id gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.actions.runs.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsRunsRequestBuilder) ByRun_id(run_id string)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if run_id != "" {
        urlTplParams["run_id"] = run_id
    }
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRun_idInteger gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.actions.runs.item collection
func (m *ItemItemActionsRunsRequestBuilder) ByRun_idInteger(run_id int32)(*ItemItemActionsRunsWithRun_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["run_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(run_id), 10)
    return NewItemItemActionsRunsWithRun_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunsRequestBuilderInternal instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsRequestBuilder) {
    m := &ItemItemActionsRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/runs{?actor*,branch*,event*,status*,per_page*,page*,created*,exclude_pull_requests*,check_suite_id*,head_sha*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsRequestBuilder instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all workflow runs for a repository. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsRunsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
func (m *ItemItemActionsRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsRequestBuilderGetRequestConfiguration)(ItemItemActionsRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunsResponseable), nil
}
// GetAsRunsGetResponse lists all workflow runs for a repository. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
func (m *ItemItemActionsRunsRequestBuilder) GetAsRunsGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsRunsRequestBuilderGetRequestConfiguration)(ItemItemActionsRunsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunsGetResponseable), nil
}
// ToGetRequestInformation lists all workflow runs for a repository. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsRequestBuilder) {
    return NewItemItemActionsRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
