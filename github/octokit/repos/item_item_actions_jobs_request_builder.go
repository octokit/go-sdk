package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsJobsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\actions\jobs
type ItemItemActionsJobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByJob_id gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.actions.jobs.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsJobsRequestBuilder) ByJob_id(job_id string)(*ItemItemActionsJobsWithJob_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if job_id != "" {
        urlTplParams["job_id"] = job_id
    }
    return NewItemItemActionsJobsWithJob_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByJob_idInteger gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.actions.jobs.item collection
func (m *ItemItemActionsJobsRequestBuilder) ByJob_idInteger(job_id int32)(*ItemItemActionsJobsWithJob_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["job_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(job_id), 10)
    return NewItemItemActionsJobsWithJob_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsJobsRequestBuilderInternal instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    m := &ItemItemActionsJobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/actions/jobs", pathParameters),
    }
    return m
}
// NewItemItemActionsJobsRequestBuilder instantiates a new JobsRequestBuilder and sets the default values.
func NewItemItemActionsJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsJobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsJobsRequestBuilderInternal(urlParams, requestAdapter)
}
