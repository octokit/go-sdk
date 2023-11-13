package repos

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunsItemAttemptsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}\attempts
type ItemItemActionsRunsItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByAttempt_number gets an item from the go-sdk.repos.item.item.actions.runs.item.attempts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsRunsItemAttemptsRequestBuilder) ByAttempt_number(attempt_number string)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attempt_number != "" {
        urlTplParams["attempt_number"] = attempt_number
    }
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByAttempt_numberInteger gets an item from the go-sdk.repos.item.item.actions.runs.item.attempts.item collection
func (m *ItemItemActionsRunsItemAttemptsRequestBuilder) ByAttempt_numberInteger(attempt_number int32)(*ItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["attempt_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(attempt_number), 10)
    return NewItemItemActionsRunsItemAttemptsWithAttempt_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunsItemAttemptsRequestBuilderInternal instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    m := &ItemItemActionsRunsItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}/attempts", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemAttemptsRequestBuilder instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
