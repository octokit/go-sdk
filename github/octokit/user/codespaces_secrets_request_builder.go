package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesSecretsRequestBuilder builds and executes requests for operations under \user\codespaces\secrets
type CodespacesSecretsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesSecretsRequestBuilderGetQueryParameters lists all development environment secrets available for a user's codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
type CodespacesSecretsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// CodespacesSecretsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesSecretsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CodespacesSecretsRequestBuilderGetQueryParameters
}
// BySecret_name gets an item from the octokit.user.codespaces.secrets.item collection
func (m *CodespacesSecretsRequestBuilder) BySecret_name(secret_name string)(*CodespacesSecretsWithSecret_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if secret_name != "" {
        urlTplParams["secret_name"] = secret_name
    }
    return NewCodespacesSecretsWithSecret_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodespacesSecretsRequestBuilderInternal instantiates a new SecretsRequestBuilder and sets the default values.
func NewCodespacesSecretsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsRequestBuilder) {
    m := &CodespacesSecretsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/secrets{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewCodespacesSecretsRequestBuilder instantiates a new SecretsRequestBuilder and sets the default values.
func NewCodespacesSecretsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesSecretsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesSecretsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all development environment secrets available for a user's codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsSecretsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/secrets#list-secrets-for-the-authenticated-user
func (m *CodespacesSecretsRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesSecretsRequestBuilderGetRequestConfiguration)(CodespacesSecretsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesSecretsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesSecretsResponseable), nil
}
// GetAsSecretsGetResponse lists all development environment secrets available for a user's codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/secrets#list-secrets-for-the-authenticated-user
func (m *CodespacesSecretsRequestBuilder) GetAsSecretsGetResponse(ctx context.Context, requestConfiguration *CodespacesSecretsRequestBuilderGetRequestConfiguration)(CodespacesSecretsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesSecretsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesSecretsGetResponseable), nil
}
// PublicKey the publicKey property
func (m *CodespacesSecretsRequestBuilder) PublicKey()(*CodespacesSecretsPublicKeyRequestBuilder) {
    return NewCodespacesSecretsPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all development environment secrets available for a user's codespaces without revealing theirencrypted values.You must authenticate using an access token with the `codespace` or `codespace:secrets` scope to use this endpoint. User must have Codespaces access to use this endpoint.GitHub Apps must have read access to the `codespaces_user_secrets` user permission to use this endpoint.
func (m *CodespacesSecretsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesSecretsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CodespacesSecretsRequestBuilder) WithUrl(rawUrl string)(*CodespacesSecretsRequestBuilder) {
    return NewCodespacesSecretsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
