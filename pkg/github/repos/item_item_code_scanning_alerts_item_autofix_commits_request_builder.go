package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\alerts\{alert_number}\autofix\commits
type ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilderInternal instantiates a new ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) {
    m := &ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/alerts/{alert_number}/autofix/commits", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder instantiates a new ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post commits an autofix for a code scanning alert.If an autofix is commited as a result of this request, then this endpoint will return a 201 Created response.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningAutofixCommitsResponseable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningAutofixCommitsResponse503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#commit-an-autofix-for-a-code-scanning-alert
func (m *ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) Post(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixCommitsable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixCommitsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofixCommitsResponse503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofixCommitsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixCommitsResponseable), nil
}
// ToPostRequestInformation commits an autofix for a code scanning alert.If an autofix is commited as a result of this request, then this endpoint will return a 201 Created response.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixCommitsable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder when successful
func (m *ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) {
    return NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
