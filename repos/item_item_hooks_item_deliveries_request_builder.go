package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemHooksItemDeliveriesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\hooks\{hook_id}\deliveries
type ItemItemHooksItemDeliveriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemHooksItemDeliveriesRequestBuilderGetQueryParameters returns a list of webhook deliveries for a webhook configured in a repository.
type ItemItemHooksItemDeliveriesRequestBuilderGetQueryParameters struct {
    // Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the `link` header for the next and previous page cursors.
    Cursor *string `uriparametername:"cursor"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // 
    Redelivery *bool `uriparametername:"redelivery"`
}
// ItemItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemHooksItemDeliveriesRequestBuilderGetQueryParameters
}
// ByDelivery_id gets an item from the go-sdk.repos.item.item.hooks.item.deliveries.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemHooksItemDeliveriesRequestBuilder) ByDelivery_id(delivery_id string)(*ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if delivery_id != "" {
        urlTplParams["delivery_id"] = delivery_id
    }
    return NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDelivery_idInteger gets an item from the go-sdk.repos.item.item.hooks.item.deliveries.item collection
func (m *ItemItemHooksItemDeliveriesRequestBuilder) ByDelivery_idInteger(delivery_id int32)(*ItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["delivery_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(delivery_id), 10)
    return NewItemItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemHooksItemDeliveriesRequestBuilderInternal instantiates a new DeliveriesRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesRequestBuilder) {
    m := &ItemItemHooksItemDeliveriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/hooks/{hook_id}/deliveries{?per_page*,cursor*,redelivery*}", pathParameters),
    }
    return m
}
// NewItemItemHooksItemDeliveriesRequestBuilder instantiates a new DeliveriesRequestBuilder and sets the default values.
func NewItemItemHooksItemDeliveriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemHooksItemDeliveriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemHooksItemDeliveriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of webhook deliveries for a webhook configured in a repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/webhooks#list-deliveries-for-a-repository-webhook
func (m *ItemItemHooksItemDeliveriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.HookDeliveryItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateHookDeliveryItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.HookDeliveryItemable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.HookDeliveryItemable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns a list of webhook deliveries for a webhook configured in a repository.
func (m *ItemItemHooksItemDeliveriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemHooksItemDeliveriesRequestBuilder) WithUrl(rawUrl string)(*ItemItemHooksItemDeliveriesRequestBuilder) {
    return NewItemItemHooksItemDeliveriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
