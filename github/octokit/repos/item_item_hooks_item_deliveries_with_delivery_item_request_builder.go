package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\hooks\{hook_id}\deliveries\{delivery_id}
type ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attempts the attempts property
func (m *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Attempts()(*ItemItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    m := &ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}", pathParameters),
    }
    return m
}
// NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder instantiates a new WithDelivery_ItemRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a delivery for a webhook configured in a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#get-a-delivery-for-a-repository-webhook
func (m *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.HookDeliveryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateHookDeliveryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.HookDeliveryable), nil
}
// ToGetRequestInformation returns a delivery for a webhook configured in a repository.
func (m *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    return NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
