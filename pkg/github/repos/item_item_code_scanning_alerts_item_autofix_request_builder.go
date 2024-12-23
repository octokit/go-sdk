package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningAlertsItemAutofixRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\alerts\{alert_number}\autofix
type ItemItemCodeScanningAlertsItemAutofixRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Commits the commits property
// returns a *ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder when successful
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) Commits()(*ItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilder) {
    return NewItemItemCodeScanningAlertsItemAutofixCommitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningAlertsItemAutofixRequestBuilderInternal instantiates a new ItemItemCodeScanningAlertsItemAutofixRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemAutofixRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemAutofixRequestBuilder) {
    m := &ItemItemCodeScanningAlertsItemAutofixRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/alerts/{alert_number}/autofix", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAlertsItemAutofixRequestBuilder instantiates a new ItemItemCodeScanningAlertsItemAutofixRequestBuilder and sets the default values.
func NewItemItemCodeScanningAlertsItemAutofixRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAlertsItemAutofixRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAlertsItemAutofixRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the status and description of an autofix for a code scanning alert.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningAutofixable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningAutofix503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-the-status-of-an-autofix-for-a-code-scanning-alert
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofix503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofixFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixable), nil
}
// Post creates an autofix for a code scanning alert.If a new autofix is to be created as a result of this request or is currently being generated, then this endpoint will return a 202 Accepted response.If an autofix already exists for a given alert, then this endpoint will return a 200 OK response.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningAutofixable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningAutofix503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#create-an-autofix-for-a-code-scanning-alert
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofix503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAutofixFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAutofixable), nil
}
// ToGetRequestInformation gets the status and description of an autofix for a code scanning alert.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates an autofix for a code scanning alert.If a new autofix is to be created as a result of this request or is currently being generated, then this endpoint will return a 202 Accepted response.If an autofix already exists for a given alert, then this endpoint will return a 200 OK response.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningAlertsItemAutofixRequestBuilder when successful
func (m *ItemItemCodeScanningAlertsItemAutofixRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAlertsItemAutofixRequestBuilder) {
    return NewItemItemCodeScanningAlertsItemAutofixRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
