package app

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HookDeliveriesItemAttemptsRequestBuilder builds and executes requests for operations under \app\hook\deliveries\{delivery_id}\attempts
type HookDeliveriesItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HookDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HookDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHookDeliveriesItemAttemptsRequestBuilderInternal instantiates a new AttemptsRequestBuilder and sets the default values.
func NewHookDeliveriesItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookDeliveriesItemAttemptsRequestBuilder) {
    m := &HookDeliveriesItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/hook/deliveries/{delivery_id}/attempts", pathParameters),
    }
    return m
}
// NewHookDeliveriesItemAttemptsRequestBuilder instantiates a new AttemptsRequestBuilder and sets the default values.
func NewHookDeliveriesItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HookDeliveriesItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHookDeliveriesItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post redeliver a delivery for the webhook configured for a GitHub App.You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
// Deprecated: This method is obsolete. Use PostAsAttemptsPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/webhooks#redeliver-a-delivery-for-an-app-webhook
func (m *HookDeliveriesItemAttemptsRequestBuilder) Post(ctx context.Context, requestConfiguration *HookDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(HookDeliveriesItemAttemptsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHookDeliveriesItemAttemptsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HookDeliveriesItemAttemptsResponseable), nil
}
// PostAsAttemptsPostResponse redeliver a delivery for the webhook configured for a GitHub App.You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/webhooks#redeliver-a-delivery-for-an-app-webhook
func (m *HookDeliveriesItemAttemptsRequestBuilder) PostAsAttemptsPostResponse(ctx context.Context, requestConfiguration *HookDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(HookDeliveriesItemAttemptsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHookDeliveriesItemAttemptsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HookDeliveriesItemAttemptsPostResponseable), nil
}
// ToPostRequestInformation redeliver a delivery for the webhook configured for a GitHub App.You must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint.
func (m *HookDeliveriesItemAttemptsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *HookDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *HookDeliveriesItemAttemptsRequestBuilder) WithUrl(rawUrl string)(*HookDeliveriesItemAttemptsRequestBuilder) {
    return NewHookDeliveriesItemAttemptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
