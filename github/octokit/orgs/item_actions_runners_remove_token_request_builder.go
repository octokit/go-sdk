package orgs

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnersRemoveTokenRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners\remove-token
type ItemActionsRunnersRemoveTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemActionsRunnersRemoveTokenRequestBuilderInternal instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    m := &ItemActionsRunnersRemoveTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runners/remove-token", pathParameters),
    }
    return m
}
// NewItemActionsRunnersRemoveTokenRequestBuilder instantiates a new RemoveTokenRequestBuilder and sets the default values.
func NewItemActionsRunnersRemoveTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRemoveTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script to remove a self-hosted runner from an organization. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using remove token:To remove your self-hosted runner from an organization, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-a-remove-token-for-an-organization
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.AuthenticationTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateAuthenticationTokenFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.AuthenticationTokenable), nil
}
// ToPostRequestInformation returns a token that you can pass to the `config` script to remove a self-hosted runner from an organization. The token expires after one hour.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using remove token:To remove your self-hosted runner from an organization, replace `TOKEN` with the remove token provided by thisendpoint.```./config.sh remove --token TOKEN```
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersRemoveTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemActionsRunnersRemoveTokenRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemActionsRunnersRemoveTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
