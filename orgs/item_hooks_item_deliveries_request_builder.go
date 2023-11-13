package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemHooksItemDeliveriesRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\deliveries
type ItemHooksItemDeliveriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHooksItemDeliveriesRequestBuilderGetQueryParameters returns a list of webhook deliveries for a webhook configured in an organization.
type ItemHooksItemDeliveriesRequestBuilderGetQueryParameters struct {
    // Used for pagination: the starting delivery from which the page of deliveries is fetched. Refer to the `link` header for the next and previous page cursors.
    Cursor *string `uriparametername:"cursor"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // 
    Redelivery *bool `uriparametername:"redelivery"`
}
// ItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemHooksItemDeliveriesRequestBuilderGetQueryParameters
}
// ByDelivery_id gets an item from the go-sdk.orgs.item.hooks.item.deliveries.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemHooksItemDeliveriesRequestBuilder) ByDelivery_id(delivery_id string)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if delivery_id != "" {
        urlTplParams["delivery_id"] = delivery_id
    }
    return NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDelivery_idInteger gets an item from the go-sdk.orgs.item.hooks.item.deliveries.item collection
func (m *ItemHooksItemDeliveriesRequestBuilder) ByDelivery_idInteger(delivery_id int32)(*ItemHooksItemDeliveriesWithDelivery_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["delivery_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(delivery_id), 10)
    return NewItemHooksItemDeliveriesWithDelivery_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemHooksItemDeliveriesRequestBuilderInternal instantiates a new DeliveriesRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesRequestBuilder) {
    m := &ItemHooksItemDeliveriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/hooks/{hook_id}/deliveries{?per_page*,cursor*,redelivery*}", pathParameters),
    }
    return m
}
// NewItemHooksItemDeliveriesRequestBuilder instantiates a new DeliveriesRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemDeliveriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a list of webhook deliveries for a webhook configured in an organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/webhooks#list-deliveries-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.HookDeliveryItemable, error) {
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
// ToGetRequestInformation returns a list of webhook deliveries for a webhook configured in an organization.
func (m *ItemHooksItemDeliveriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemHooksItemDeliveriesRequestBuilder) WithUrl(rawUrl string)(*ItemHooksItemDeliveriesRequestBuilder) {
    return NewItemHooksItemDeliveriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
