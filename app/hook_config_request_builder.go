package app

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HookConfigRequestBuilder builds and executes requests for operations under \app\hook\config
type HookConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HookConfigRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HookConfigRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HookConfigRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HookConfigRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHookConfigRequestBuilderInternal instantiates a new ConfigRequestBuilder and sets the default values.
func NewHookConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookConfigRequestBuilder) {
    m := &HookConfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/hook/config", pathParameters),
    }
    return m
}
// NewHookConfigRequestBuilder instantiates a new ConfigRequestBuilder and sets the default values.
func NewHookConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHookConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the webhook configuration for a GitHub App. For more information about configuring a webhook for your app, see "[Creating a GitHub App](/developers/apps/creating-a-github-app)."You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/webhooks#get-a-webhook-configuration-for-an-app
func (m *HookConfigRequestBuilder) Get(ctx context.Context, requestConfiguration *HookConfigRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.WebhookConfigable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateWebhookConfigFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.WebhookConfigable), nil
}
// Patch updates the webhook configuration for a GitHub App. For more information about configuring a webhook for your app, see "[Creating a GitHub App](/developers/apps/creating-a-github-app)."You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/webhooks#update-a-webhook-configuration-for-an-app
func (m *HookConfigRequestBuilder) Patch(ctx context.Context, body HookConfigPatchRequestBodyable, requestConfiguration *HookConfigRequestBuilderPatchRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.WebhookConfigable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateWebhookConfigFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.WebhookConfigable), nil
}
// ToGetRequestInformation returns the webhook configuration for a GitHub App. For more information about configuring a webhook for your app, see "[Creating a GitHub App](/developers/apps/creating-a-github-app)."You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
func (m *HookConfigRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HookConfigRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation updates the webhook configuration for a GitHub App. For more information about configuring a webhook for your app, see "[Creating a GitHub App](/developers/apps/creating-a-github-app)."You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
func (m *HookConfigRequestBuilder) ToPatchRequestInformation(ctx context.Context, body HookConfigPatchRequestBodyable, requestConfiguration *HookConfigRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *HookConfigRequestBuilder) WithUrl(rawUrl string)(*HookConfigRequestBuilder) {
    return NewHookConfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
