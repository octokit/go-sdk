package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCheckSuitesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites
type ItemItemCheckSuitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckSuitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCheck_suite_id gets an item from the go-sdk.repos.item.item.checkSuites.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemCheckSuitesRequestBuilder) ByCheck_suite_id(check_suite_id string)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if check_suite_id != "" {
        urlTplParams["check_suite_id"] = check_suite_id
    }
    return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCheck_suite_idInteger gets an item from the go-sdk.repos.item.item.checkSuites.item collection
func (m *ItemItemCheckSuitesRequestBuilder) ByCheck_suite_idInteger(check_suite_id int32)(*ItemItemCheckSuitesWithCheck_suite_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["check_suite_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(check_suite_id), 10)
    return NewItemItemCheckSuitesWithCheck_suite_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCheckSuitesRequestBuilderInternal instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCheckSuitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesRequestBuilder) {
    m := &ItemItemCheckSuitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/check-suites", pathParameters),
    }
    return m
}
// NewItemItemCheckSuitesRequestBuilder instantiates a new CheckSuitesRequestBuilder and sets the default values.
func NewItemItemCheckSuitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.By default, check suites are automatically created when you create a [check run](https://docs.github.com/rest/checks/runs). You only need to use this endpoint for manually creating check suites when you've disabled automatic creation using "[Update repository preferences for check suites](https://docs.github.com/rest/checks/suites#update-repository-preferences-for-check-suites)". Your GitHub App must have the `checks:write` permission to create check suites.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#create-a-check-suite
func (m *ItemItemCheckSuitesRequestBuilder) Post(ctx context.Context, body ItemItemCheckSuitesPostRequestBodyable, requestConfiguration *ItemItemCheckSuitesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CheckSuiteable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCheckSuiteFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CheckSuiteable), nil
}
// Preferences the preferences property
func (m *ItemItemCheckSuitesRequestBuilder) Preferences()(*ItemItemCheckSuitesPreferencesRequestBuilder) {
    return NewItemItemCheckSuitesPreferencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPostRequestInformation **Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array and a `null` value for `head_branch`.By default, check suites are automatically created when you create a [check run](https://docs.github.com/rest/checks/runs). You only need to use this endpoint for manually creating check suites when you've disabled automatic creation using "[Update repository preferences for check suites](https://docs.github.com/rest/checks/suites#update-repository-preferences-for-check-suites)". Your GitHub App must have the `checks:write` permission to create check suites.
func (m *ItemItemCheckSuitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemCheckSuitesPostRequestBodyable, requestConfiguration *ItemItemCheckSuitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckSuitesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckSuitesRequestBuilder) {
    return NewItemItemCheckSuitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
