package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersRegistrationTokenRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\registration-token
type ItemItemActionsRunnersRegistrationTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    m := &ItemItemActionsRunnersRegistrationTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runners/registration-token", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersRegistrationTokenRequestBuilder instantiates a new RegistrationTokenRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a token that you can pass to the `config` script. The tokenexpires after one hour.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using registration token:Configure your self-hosted runner, replacing `TOKEN` with the registration token providedby this endpoint.```config.sh --url https://github.com/octo-org/octo-repo-artifacts --token TOKEN```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#create-a-registration-token-for-a-repository
func (m *ItemItemActionsRunnersRegistrationTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.AuthenticationTokenable, error) {
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
// ToPostRequestInformation returns a token that you can pass to the `config` script. The tokenexpires after one hour.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.Example using registration token:Configure your self-hosted runner, replacing `TOKEN` with the registration token providedby this endpoint.```config.sh --url https://github.com/octo-org/octo-repo-artifacts --token TOKEN```
func (m *ItemItemActionsRunnersRegistrationTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRegistrationTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunnersRegistrationTokenRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
