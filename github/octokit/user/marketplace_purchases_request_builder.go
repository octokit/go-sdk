package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Marketplace_purchasesRequestBuilder builds and executes requests for operations under \user\marketplace_purchases
type Marketplace_purchasesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Marketplace_purchasesRequestBuilderGetQueryParameters lists the active subscriptions for the authenticated user. GitHub Apps must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint. OAuth apps must authenticate using an [OAuth token](https://docs.github.com/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps).
type Marketplace_purchasesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// Marketplace_purchasesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Marketplace_purchasesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *Marketplace_purchasesRequestBuilderGetQueryParameters
}
// NewMarketplace_purchasesRequestBuilderInternal instantiates a new Marketplace_purchasesRequestBuilder and sets the default values.
func NewMarketplace_purchasesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_purchasesRequestBuilder) {
    m := &Marketplace_purchasesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/marketplace_purchases{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewMarketplace_purchasesRequestBuilder instantiates a new Marketplace_purchasesRequestBuilder and sets the default values.
func NewMarketplace_purchasesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Marketplace_purchasesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMarketplace_purchasesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the active subscriptions for the authenticated user. GitHub Apps must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint. OAuth apps must authenticate using an [OAuth token](https://docs.github.com/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/apps/marketplace#list-subscriptions-for-the-authenticated-user
func (m *Marketplace_purchasesRequestBuilder) Get(ctx context.Context, requestConfiguration *Marketplace_purchasesRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.UserMarketplacePurchaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateUserMarketplacePurchaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.UserMarketplacePurchaseable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.UserMarketplacePurchaseable)
        }
    }
    return val, nil
}
// Stubbed the stubbed property
func (m *Marketplace_purchasesRequestBuilder) Stubbed()(*Marketplace_purchasesStubbedRequestBuilder) {
    return NewMarketplace_purchasesStubbedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists the active subscriptions for the authenticated user. GitHub Apps must use a [user access token](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app), created for a user who has authorized your GitHub App, to access this endpoint. OAuth apps must authenticate using an [OAuth token](https://docs.github.com/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps).
func (m *Marketplace_purchasesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Marketplace_purchasesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *Marketplace_purchasesRequestBuilder) WithUrl(rawUrl string)(*Marketplace_purchasesRequestBuilder) {
    return NewMarketplace_purchasesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
