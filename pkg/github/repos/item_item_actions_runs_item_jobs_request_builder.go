package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i08b4bf52b4de47221f213f980c03d04a2d7557c63c7158cc63f99ebaac398e57 "github.com/octokit/go-sdk/pkg/github/repos/item/item/actions/runs/item/jobs"
)

// ItemItemActionsRunsItemJobsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\runs\{run_id}\jobs
type ItemItemActionsRunsItemJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunsItemJobsRequestBuilderGetQueryParameters lists jobs for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).
type ItemItemActionsRunsItemJobsRequestBuilderGetQueryParameters struct {
    // Filters jobs by their `completed_at` timestamp. `latest` returns jobs from the most recent execution of the workflow run. `all` returns all jobs for a workflow run, including from old executions of the workflow run.
    // Deprecated: This property is deprecated, use filterAsGetFilterQueryParameterType instead
    Filter *string `uriparametername:"filter"`
    // Filters jobs by their `completed_at` timestamp. `latest` returns jobs from the most recent execution of the workflow run. `all` returns all jobs for a workflow run, including from old executions of the workflow run.
    FilterAsGetFilterQueryParameterType *i08b4bf52b4de47221f213f980c03d04a2d7557c63c7158cc63f99ebaac398e57.GetFilterQueryParameterType `uriparametername:"filter"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemActionsRunsItemJobsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemJobsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunsItemJobsRequestBuilderGetQueryParameters
}
// NewItemItemActionsRunsItemJobsRequestBuilderInternal instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemJobsRequestBuilder) {
    m := &ItemItemActionsRunsItemJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/runs/{run_id}/jobs{?filter*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemJobsRequestBuilder instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists jobs for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).
// Deprecated: This method is obsolete. Use GetAsJobsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-jobs#list-jobs-for-a-workflow-run
func (m *ItemItemActionsRunsItemJobsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemJobsRequestBuilderGetRequestConfiguration)(ItemItemActionsRunsItemJobsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunsItemJobsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunsItemJobsResponseable), nil
}
// GetAsJobsGetResponse lists jobs for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-jobs#list-jobs-for-a-workflow-run
func (m *ItemItemActionsRunsItemJobsRequestBuilder) GetAsJobsGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemJobsRequestBuilderGetRequestConfiguration)(ItemItemActionsRunsItemJobsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunsItemJobsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunsItemJobsGetResponseable), nil
}
// ToGetRequestInformation lists jobs for a workflow run. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint. You can use parameters to narrow the list of results. For more information about using parameters, see [Parameters](https://docs.github.com/rest/overview/resources-in-the-rest-api#parameters).
func (m *ItemItemActionsRunsItemJobsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunsItemJobsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunsItemJobsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemJobsRequestBuilder) {
    return NewItemItemActionsRunsItemJobsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
