package users

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSocial_accountsRequestBuilder builds and executes requests for operations under \users\{username}\social_accounts
type ItemSocial_accountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSocial_accountsRequestBuilderGetQueryParameters lists social media accounts for a user. This endpoint is accessible by anyone.
type ItemSocial_accountsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemSocial_accountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSocial_accountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSocial_accountsRequestBuilderGetQueryParameters
}
// NewItemSocial_accountsRequestBuilderInternal instantiates a new Social_accountsRequestBuilder and sets the default values.
func NewItemSocial_accountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSocial_accountsRequestBuilder) {
    m := &ItemSocial_accountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/social_accounts{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemSocial_accountsRequestBuilder instantiates a new Social_accountsRequestBuilder and sets the default values.
func NewItemSocial_accountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSocial_accountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSocial_accountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists social media accounts for a user. This endpoint is accessible by anyone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/social-accounts#list-social-accounts-for-a-user
func (m *ItemSocial_accountsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSocial_accountsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SocialAccountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSocialAccountFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SocialAccountable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SocialAccountable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists social media accounts for a user. This endpoint is accessible by anyone.
func (m *ItemSocial_accountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSocial_accountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSocial_accountsRequestBuilder) WithUrl(rawUrl string)(*ItemSocial_accountsRequestBuilder) {
    return NewItemSocial_accountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
