package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemHooksItemDeliveriesItemAttemptsRequestBuilder builds and executes requests for operations under \orgs\{org}\hooks\{hook_id}\deliveries\{delivery_id}\attempts
type ItemHooksItemDeliveriesItemAttemptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    m := &ItemHooksItemDeliveriesItemAttemptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}/attempts", pathParameters),
    }
    return m
}
// NewItemHooksItemDeliveriesItemAttemptsRequestBuilder instantiates a new AttemptsRequestBuilder and sets the default values.
func NewItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post redeliver a delivery for a webhook configured in an organization.
// Deprecated: This method is obsolete. Use PostAsAttemptsPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/webhooks#redeliver-a-delivery-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(ItemHooksItemDeliveriesItemAttemptsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemHooksItemDeliveriesItemAttemptsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemHooksItemDeliveriesItemAttemptsResponseable), nil
}
// PostAsAttemptsPostResponse redeliver a delivery for a webhook configured in an organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/webhooks#redeliver-a-delivery-for-an-organization-webhook
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) PostAsAttemptsPostResponse(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(ItemHooksItemDeliveriesItemAttemptsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemHooksItemDeliveriesItemAttemptsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemHooksItemDeliveriesItemAttemptsPostResponseable), nil
}
// ToPostRequestInformation redeliver a delivery for a webhook configured in an organization.
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemHooksItemDeliveriesItemAttemptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemHooksItemDeliveriesItemAttemptsRequestBuilder) WithUrl(rawUrl string)(*ItemHooksItemDeliveriesItemAttemptsRequestBuilder) {
    return NewItemHooksItemDeliveriesItemAttemptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
