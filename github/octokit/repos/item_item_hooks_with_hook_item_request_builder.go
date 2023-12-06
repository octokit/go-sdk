package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemHooksWithHook_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\hooks\{hook_id}
type ItemItemHooksWithHook_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemHooksWithHook_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksWithHook_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemHooksWithHook_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksWithHook_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemHooksWithHook_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksWithHook_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Config the config property
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Config()(*ItemItemHooksItemConfigRequestBuilder) {
    return NewItemItemHooksItemConfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemHooksWithHook_ItemRequestBuilderInternal instantiates a new WithHook_ItemRequestBuilder and sets the default values.
func NewItemItemHooksWithHook_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksWithHook_ItemRequestBuilder) {
    m := &ItemItemHooksWithHook_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/hooks/{hook_id}", pathParameters),
    }
    return m
}
// NewItemItemHooksWithHook_ItemRequestBuilder instantiates a new WithHook_ItemRequestBuilder and sets the default values.
func NewItemItemHooksWithHook_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksWithHook_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemHooksWithHook_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a repository webhook
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#delete-a-repository-webhook
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Deliveries the deliveries property
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Deliveries()(*ItemItemHooksItemDeliveriesRequestBuilder) {
    return NewItemItemHooksItemDeliveriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns a webhook configured in a repository. To get only the webhook `config` properties, see "[Get a webhook configuration for a repository](/rest/webhooks/repo-config#get-a-webhook-configuration-for-a-repository)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#get-a-repository-webhook
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Hookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateHookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Hookable), nil
}
// Patch updates a webhook configured in a repository. If you previously had a `secret` set, you must provide the same `secret` or set a new `secret` or the secret will be removed. If you are only updating individual webhook `config` properties, use "[Update a webhook configuration for a repository](/rest/webhooks/repo-config#update-a-webhook-configuration-for-a-repository)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#update-a-repository-webhook
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Patch(ctx context.Context, body ItemItemHooksItemWithHook_PatchRequestBodyable, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderPatchRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Hookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateHookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Hookable), nil
}
// Pings the pings property
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Pings()(*ItemItemHooksItemPingsRequestBuilder) {
    return NewItemItemHooksItemPingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tests the tests property
func (m *ItemItemHooksWithHook_ItemRequestBuilder) Tests()(*ItemItemHooksItemTestsRequestBuilder) {
    return NewItemItemHooksItemTestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *ItemItemHooksWithHook_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns a webhook configured in a repository. To get only the webhook `config` properties, see "[Get a webhook configuration for a repository](/rest/webhooks/repo-config#get-a-webhook-configuration-for-a-repository)."
func (m *ItemItemHooksWithHook_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a webhook configured in a repository. If you previously had a `secret` set, you must provide the same `secret` or set a new `secret` or the secret will be removed. If you are only updating individual webhook `config` properties, use "[Update a webhook configuration for a repository](/rest/webhooks/repo-config#update-a-webhook-configuration-for-a-repository)."
func (m *ItemItemHooksWithHook_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemHooksItemWithHook_PatchRequestBodyable, requestConfiguration *ItemItemHooksWithHook_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemHooksWithHook_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemHooksWithHook_ItemRequestBuilder) {
    return NewItemItemHooksWithHook_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
