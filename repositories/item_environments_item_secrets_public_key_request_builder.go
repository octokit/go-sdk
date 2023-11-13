package repositories

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsItemSecretsPublicKeyRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\secrets\public-key
type ItemEnvironmentsItemSecretsPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEnvironmentsItemSecretsPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemSecretsPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    m := &ItemEnvironmentsItemSecretsPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/secrets/public-key", pathParameters),
    }
    return m
}
// NewItemEnvironmentsItemSecretsPublicKeyRequestBuilder instantiates a new PublicKeyRequestBuilder and sets the default values.
func NewItemEnvironmentsItemSecretsPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemSecretsPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the public key for an environment, which you need to encrypt environmentsecrets. You need to encrypt a secret before you can create or update secrets.Anyone with read access to the repository can use this endpoint.If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/secrets#get-an-environment-public-key
func (m *ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEnvironmentsItemSecretsPublicKeyRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsPublicKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateActionsPublicKeyFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsPublicKeyable), nil
}
// ToGetRequestInformation get the public key for an environment, which you need to encrypt environmentsecrets. You need to encrypt a secret before you can create or update secrets.Anyone with read access to the repository can use this endpoint.If the repository is private you must use an access token with the `repo` scope.GitHub Apps must have the `secrets` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read secrets.
func (m *ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemSecretsPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) WithUrl(rawUrl string)(*ItemEnvironmentsItemSecretsPublicKeyRequestBuilder) {
    return NewItemEnvironmentsItemSecretsPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
