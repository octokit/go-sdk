package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsSecretsPublicKeyRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\secrets\public-key
type ItemActionsSecretsPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsSecretsPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsSecretsPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsSecretsPublicKeyRequestBuilderInternal instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemActionsSecretsPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsPublicKeyRequestBuilder) {
    m := &ItemActionsSecretsPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/secrets/public-key", pathParameters),
    }
    return m
}
// NewItemActionsSecretsPublicKeyRequestBuilder instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemActionsSecretsPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsSecretsPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsSecretsPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets your public key, which you need to encrypt secrets. You need toencrypt a secret before you can create or update secrets.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#get-an-organization-public-key
func (m *ItemActionsSecretsPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsSecretsPublicKeyRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsPublicKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateActionsPublicKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsPublicKeyable), nil
}
// ToGetRequestInformation gets your public key, which you need to encrypt secrets. You need toencrypt a secret before you can create or update secrets.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemActionsSecretsPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsSecretsPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemActionsSecretsPublicKeyRequestBuilder) WithUrl(rawUrl string)(*ItemActionsSecretsPublicKeyRequestBuilder) {
    return NewItemActionsSecretsPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
