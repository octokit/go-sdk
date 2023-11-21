package marketplace_listing

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2bc2ea0d9551b2a66882b3fff8fc93c10c238901010a22ba83c4e973264f4d48 "github.com/octokit/go-sdk/github/octokit/marketplace_listing/stubbed/plans/item/accounts"
)

// StubbedPlansItemAccountsRequestBuilder builds and executes requests for operations under \marketplace_listing\stubbed\plans\{plan_id}\accounts
type StubbedPlansItemAccountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StubbedPlansItemAccountsRequestBuilderGetQueryParameters returns repository and organization accounts associated with the specified plan, including free plans. For per-seat pricing, you see the list of accounts that have purchased the plan, including the number of seats purchased. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.GitHub Apps must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth apps must use [basic authentication](https://docs.github.com/rest/overview/other-authentication-methods#basic-authentication) with their client ID and client secret to access this endpoint.
type StubbedPlansItemAccountsRequestBuilderGetQueryParameters struct {
    // To return the oldest accounts first, set to `asc`. Ignored without the `sort` parameter.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // To return the oldest accounts first, set to `asc`. Ignored without the `sort` parameter.
    DirectionAsGetDirectionQueryParameterType *i2bc2ea0d9551b2a66882b3fff8fc93c10c238901010a22ba83c4e973264f4d48.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *i2bc2ea0d9551b2a66882b3fff8fc93c10c238901010a22ba83c4e973264f4d48.GetSortQueryParameterType `uriparametername:"sort"`
}
// StubbedPlansItemAccountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StubbedPlansItemAccountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *StubbedPlansItemAccountsRequestBuilderGetQueryParameters
}
// NewStubbedPlansItemAccountsRequestBuilderInternal instantiates a new AccountsRequestBuilder and sets the default values.
func NewStubbedPlansItemAccountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansItemAccountsRequestBuilder) {
    m := &StubbedPlansItemAccountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/marketplace_listing/stubbed/plans/{plan_id}/accounts{?sort*,direction*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewStubbedPlansItemAccountsRequestBuilder instantiates a new AccountsRequestBuilder and sets the default values.
func NewStubbedPlansItemAccountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StubbedPlansItemAccountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStubbedPlansItemAccountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns repository and organization accounts associated with the specified plan, including free plans. For per-seat pricing, you see the list of accounts that have purchased the plan, including the number of seats purchased. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.GitHub Apps must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth apps must use [basic authentication](https://docs.github.com/rest/overview/other-authentication-methods#basic-authentication) with their client ID and client secret to access this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/marketplace#list-accounts-for-a-plan-stubbed
func (m *StubbedPlansItemAccountsRequestBuilder) Get(ctx context.Context, requestConfiguration *StubbedPlansItemAccountsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MarketplacePurchaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateMarketplacePurchaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MarketplacePurchaseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.MarketplacePurchaseable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns repository and organization accounts associated with the specified plan, including free plans. For per-seat pricing, you see the list of accounts that have purchased the plan, including the number of seats purchased. When someone submits a plan change that won't be processed until the end of their billing cycle, you will also see the upcoming pending change.GitHub Apps must use a [JWT](https://docs.github.com/apps/building-github-apps/authenticating-with-github-apps/#authenticating-as-a-github-app) to access this endpoint. OAuth apps must use [basic authentication](https://docs.github.com/rest/overview/other-authentication-methods#basic-authentication) with their client ID and client secret to access this endpoint.
func (m *StubbedPlansItemAccountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *StubbedPlansItemAccountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *StubbedPlansItemAccountsRequestBuilder) WithUrl(rawUrl string)(*StubbedPlansItemAccountsRequestBuilder) {
    return NewStubbedPlansItemAccountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
