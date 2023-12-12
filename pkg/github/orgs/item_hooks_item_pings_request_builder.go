package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemHooksItemPingsRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\pings
type ItemHooksItemPingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHooksItemPingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHooksItemPingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemHooksItemPingsRequestBuilderInternal instantiates a new PingsRequestBuilder and sets the default values.
func NewItemHooksItemPingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemPingsRequestBuilder) {
    m := &ItemHooksItemPingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/hooks/{hook_id}/pings", pathParameters),
    }
    return m
}
// NewItemHooksItemPingsRequestBuilder instantiates a new PingsRequestBuilder and sets the default values.
func NewItemHooksItemPingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemPingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemPingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this will trigger a [ping event](https://docs.github.com/webhooks/#ping-event) to be sent to the hook.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/webhooks#ping-an-organization-webhook
func (m *ItemHooksItemPingsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemHooksItemPingsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation this will trigger a [ping event](https://docs.github.com/webhooks/#ping-event) to be sent to the hook.
func (m *ItemHooksItemPingsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemHooksItemPingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemHooksItemPingsRequestBuilder) WithUrl(rawUrl string)(*ItemHooksItemPingsRequestBuilder) {
    return NewItemHooksItemPingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
