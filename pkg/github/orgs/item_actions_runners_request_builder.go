package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnersRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runners
type ItemActionsRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnersRequestBuilderGetQueryParameters lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
type ItemActionsRunnersRequestBuilderGetQueryParameters struct {
    // The name of a self-hosted runner.
    Name *string `uriparametername:"name"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemActionsRunnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsRunnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsRunnersRequestBuilderGetQueryParameters
}
// ByRunner_id gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.actions.runners.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemActionsRunnersRequestBuilder) ByRunner_id(runner_id string)(*ItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if runner_id != "" {
        urlTplParams["runner_id"] = runner_id
    }
    return NewItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByRunner_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.actions.runners.item collection
func (m *ItemActionsRunnersRequestBuilder) ByRunner_idInteger(runner_id int32)(*ItemActionsRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["runner_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(runner_id), 10)
    return NewItemActionsRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRunnersRequestBuilderInternal instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemActionsRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRequestBuilder) {
    m := &ItemActionsRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runners{?name*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnersRequestBuilder instantiates a new RunnersRequestBuilder and sets the default values.
func NewItemActionsRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Downloads the downloads property
func (m *ItemActionsRunnersRequestBuilder) Downloads()(*ItemActionsRunnersDownloadsRequestBuilder) {
    return NewItemActionsRunnersDownloadsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateJitconfig the generateJitconfig property
func (m *ItemActionsRunnersRequestBuilder) GenerateJitconfig()(*ItemActionsRunnersGenerateJitconfigRequestBuilder) {
    return NewItemActionsRunnersGenerateJitconfigRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// Deprecated: This method is obsolete. Use GetAsRunnersGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#list-self-hosted-runners-for-an-organization
func (m *ItemActionsRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemActionsRunnersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersResponseable), nil
}
// GetAsRunnersGetResponse lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#list-self-hosted-runners-for-an-organization
func (m *ItemActionsRunnersRequestBuilder) GetAsRunnersGetResponse(ctx context.Context, requestConfiguration *ItemActionsRunnersRequestBuilderGetRequestConfiguration)(ItemActionsRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnersGetResponseable), nil
}
// RegistrationToken the registrationToken property
func (m *ItemActionsRunnersRequestBuilder) RegistrationToken()(*ItemActionsRunnersRegistrationTokenRequestBuilder) {
    return NewItemActionsRunnersRegistrationTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveToken the removeToken property
func (m *ItemActionsRunnersRequestBuilder) RemoveToken()(*ItemActionsRunnersRemoveTokenRequestBuilder) {
    return NewItemActionsRunnersRemoveTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists all self-hosted runners configured in an organization.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
func (m *ItemActionsRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsRunnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemActionsRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnersRequestBuilder) {
    return NewItemActionsRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
