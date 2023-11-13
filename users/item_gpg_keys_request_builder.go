package users

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemGpg_keysRequestBuilder builds and executes requests for operations under \users\{username}\gpg_keys
type ItemGpg_keysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGpg_keysRequestBuilderGetQueryParameters lists the GPG keys for a user. This information is accessible by anyone.
type ItemGpg_keysRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemGpg_keysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGpg_keysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGpg_keysRequestBuilderGetQueryParameters
}
// NewItemGpg_keysRequestBuilderInternal instantiates a new Gpg_keysRequestBuilder and sets the default values.
func NewItemGpg_keysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGpg_keysRequestBuilder) {
    m := &ItemGpg_keysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/gpg_keys{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemGpg_keysRequestBuilder instantiates a new Gpg_keysRequestBuilder and sets the default values.
func NewItemGpg_keysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGpg_keysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGpg_keysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the GPG keys for a user. This information is accessible by anyone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/gpg-keys#list-gpg-keys-for-a-user
func (m *ItemGpg_keysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGpg_keysRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.GpgKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateGpgKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.GpgKeyable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.GpgKeyable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the GPG keys for a user. This information is accessible by anyone.
func (m *ItemGpg_keysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGpg_keysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemGpg_keysRequestBuilder) WithUrl(rawUrl string)(*ItemGpg_keysRequestBuilder) {
    return NewItemGpg_keysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
