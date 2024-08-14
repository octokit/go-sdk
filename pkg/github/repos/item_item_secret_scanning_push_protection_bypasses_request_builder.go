package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemSecretScanningPushProtectionBypassesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\secret-scanning\push-protection-bypasses
type ItemItemSecretScanningPushProtectionBypassesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemSecretScanningPushProtectionBypassesRequestBuilderInternal instantiates a new ItemItemSecretScanningPushProtectionBypassesRequestBuilder and sets the default values.
func NewItemItemSecretScanningPushProtectionBypassesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningPushProtectionBypassesRequestBuilder) {
    m := &ItemItemSecretScanningPushProtectionBypassesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/secret-scanning/push-protection-bypasses", pathParameters),
    }
    return m
}
// NewItemItemSecretScanningPushProtectionBypassesRequestBuilder instantiates a new ItemItemSecretScanningPushProtectionBypassesRequestBuilder and sets the default values.
func NewItemItemSecretScanningPushProtectionBypassesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecretScanningPushProtectionBypassesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecretScanningPushProtectionBypassesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a bypass for a previously push protected secret.The authenticated user must be the original author of the committed secret.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a SecretScanningPushProtectionBypassable when successful
// returns a SecretScanningPushProtectionBypass503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/secret-scanning/secret-scanning#create-a-push-protection-bypass
func (m *ItemItemSecretScanningPushProtectionBypassesRequestBuilder) Post(ctx context.Context, body ItemItemSecretScanningPushProtectionBypassesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningPushProtectionBypass503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSecretScanningPushProtectionBypassFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassable), nil
}
// ToPostRequestInformation creates a bypass for a previously push protected secret.The authenticated user must be the original author of the committed secret.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemSecretScanningPushProtectionBypassesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemSecretScanningPushProtectionBypassesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemSecretScanningPushProtectionBypassesRequestBuilder when successful
func (m *ItemItemSecretScanningPushProtectionBypassesRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecretScanningPushProtectionBypassesRequestBuilder) {
    return NewItemItemSecretScanningPushProtectionBypassesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
