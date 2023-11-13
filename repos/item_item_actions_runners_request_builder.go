package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners
type ItemItemActionsRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersRequestBuilderGetQueryParameters lists all self-hosted runners configured in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
type ItemItemActionsRunnersRequestBuilderGetQueryParameters struct {
    // The name of a self-hosted runner.
    Name *string `uriparametername:"name"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemActionsRunnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsRunnersRequestBuilderGetQueryParameters
}
// ByRunner_id gets an item from the go-sdk.repos.item.item.actions.runners.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsRunnersRequestBuilder) ByRunner_id(runner_id string)(*ItemItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if runner_id != "" {
        urlTplParams["runner_id"] = runner_id
    }
    return NewItemItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRunner_idInteger gets an item from the go-sdk.repos.item.item.actions.runners.item collection
func (m *ItemItemActionsRunnersRequestBuilder) ByRunner_idInteger(runner_id int32)(*ItemItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["runner_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(runner_id), 10)
    return NewItemItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsRunnersRequestBuilderInternal instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    m := &ItemItemActionsRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runners{?name*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersRequestBuilder instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemItemActionsRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Downloads the downloads property
func (m *ItemItemActionsRunnersRequestBuilder) Downloads()(*ItemItemActionsRunnersDownloadsRequestBuilder) {
    return NewItemItemActionsRunnersDownloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateJitconfig the generateJitconfig property
func (m *ItemItemActionsRunnersRequestBuilder) GenerateJitconfig()(*ItemItemActionsRunnersGenerateJitconfigRequestBuilder) {
    return NewItemItemActionsRunnersGenerateJitconfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get lists all self-hosted runners configured in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// Deprecated: This method is obsolete. Use GetAsRunnersGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#list-self-hosted-runners-for-a-repository
func (m *ItemItemActionsRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemItemActionsRunnersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersResponseable), nil
}
// GetAsRunnersGetResponse lists all self-hosted runners configured in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#list-self-hosted-runners-for-a-repository
func (m *ItemItemActionsRunnersRequestBuilder) GetAsRunnersGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemItemActionsRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersGetResponseable), nil
}
// RegistrationToken the registrationToken property
func (m *ItemItemActionsRunnersRequestBuilder) RegistrationToken()(*ItemItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemItemActionsRunnersRegistrationTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveToken the removeToken property
func (m *ItemItemActionsRunnersRequestBuilder) RemoveToken()(*ItemItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemItemActionsRunnersRemoveTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all self-hosted runners configured in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
func (m *ItemItemActionsRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersRequestBuilder) {
    return NewItemItemActionsRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
