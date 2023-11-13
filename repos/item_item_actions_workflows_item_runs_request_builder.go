package repos

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i947ca746c0f4f514db23f9b4557533da5768da4d06e4a1fd8edc71f58bad5c33 "github.com/octokit/go-sdk/repos/item/item/actions/workflows/item/runs"
)

// ItemItemActionsWorkflowsItemRunsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\workflows\{workflow_id}\runs
type ItemItemActionsWorkflowsItemRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
type ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters struct {
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
    StatusAsGetStatusQueryParameterType *i947ca746c0f4f514db23f9b4557533da5768da4d06e4a1fd8edc71f58bad5c33.GetStatusQueryParameterType `uriparametername:"status"`
}
// ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsWorkflowsItemRunsRequestBuilderGetQueryParameters
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    m := &ItemItemActionsWorkflowsItemRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs{?actor*,branch*,event*,status*,per_page*,page*,created*,exclude_pull_requests*,check_suite_id*,head_sha*}", pathParameters),
    }
    return m
}
// NewItemItemActionsWorkflowsItemRunsRequestBuilder instantiates a new RunsRequestBuilder and sets the default values.
func NewItemItemActionsWorkflowsItemRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsWorkflowsItemRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
// Deprecated: This method is obsolete. Use GetAsRunsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#list-workflow-runs-for-a-workflow
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(ItemItemActionsWorkflowsItemRunsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsWorkflowsItemRunsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsWorkflowsItemRunsResponseable), nil
}
// GetAsRunsGetResponse list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#list-workflow-runs-for-a-workflow
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) GetAsRunsGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(ItemItemActionsWorkflowsItemRunsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsWorkflowsItemRunsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsWorkflowsItemRunsGetResponseable), nil
}
// ToGetRequestInformation list all workflow runs for a workflow. You can replace `workflow_id` with the workflow file name. For example, you could use `main.yaml`. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope.
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsWorkflowsItemRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsWorkflowsItemRunsRequestBuilder) {
    return NewItemItemActionsWorkflowsItemRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
