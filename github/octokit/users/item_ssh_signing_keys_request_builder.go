package users

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSsh_signing_keysRequestBuilder builds and executes requests for operations under \users\{username}\ssh_signing_keys
type ItemSsh_signing_keysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSsh_signing_keysRequestBuilderGetQueryParameters lists the SSH signing keys for a user. This operation is accessible by anyone.
type ItemSsh_signing_keysRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemSsh_signing_keysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSsh_signing_keysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSsh_signing_keysRequestBuilderGetQueryParameters
}
// NewItemSsh_signing_keysRequestBuilderInternal instantiates a new Ssh_signing_keysRequestBuilder and sets the default values.
func NewItemSsh_signing_keysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSsh_signing_keysRequestBuilder) {
    m := &ItemSsh_signing_keysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/ssh_signing_keys{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemSsh_signing_keysRequestBuilder instantiates a new Ssh_signing_keysRequestBuilder and sets the default values.
func NewItemSsh_signing_keysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSsh_signing_keysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSsh_signing_keysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the SSH signing keys for a user. This operation is accessible by anyone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/ssh-signing-keys#list-ssh-signing-keys-for-a-user
func (m *ItemSsh_signing_keysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSsh_signing_keysRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SshSigningKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateSshSigningKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SshSigningKeyable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.SshSigningKeyable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the SSH signing keys for a user. This operation is accessible by anyone.
func (m *ItemSsh_signing_keysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSsh_signing_keysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSsh_signing_keysRequestBuilder) WithUrl(rawUrl string)(*ItemSsh_signing_keysRequestBuilder) {
    return NewItemSsh_signing_keysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
