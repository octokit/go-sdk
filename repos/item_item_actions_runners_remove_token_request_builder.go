package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersRemoveTokenRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\remove-token
type ItemItemActionsRunnersRemoveTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    m := &ItemItemActionsRunnersRemoveTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runners/remove-token", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersRemoveTokenRequestBuilder instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRemoveTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to remove a self-hosted runner froma repository. The token expires after one hour.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using remove token:To remove your self-hosted runner from a repository, replace TOKEN withthe remove token provided by this endpoint.```config.sh remove --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-a-remove-token-for-a-repository
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AuthenticationTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateAuthenticationTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AuthenticationTokenable), nil
}
// ToPostRequestInformation returns a token that you can pass to remove a self-hosted runner froma repository. The token expires after one hour.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using remove token:To remove your self-hosted runner from a repository, replace TOKEN withthe remove token provided by this endpoint.```config.sh remove --token TOKEN```
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunnersRemoveTokenRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemItemActionsRunnersRemoveTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
