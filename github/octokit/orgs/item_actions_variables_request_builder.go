package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemActionsVariablesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\variables
type ItemActionsVariablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsVariablesRequestBuilderGetQueryParameters lists all organization variables.You must authenticate using an access token with the `admin:org` scope to use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `organization_actions_variables:read` organization permission to use this endpoint. Authenticated users must have collaborator access to a repository to create, update, or read variables.
type ItemActionsVariablesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 30).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemActionsVariablesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsVariablesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActionsVariablesRequestBuilderGetQueryParameters
}
// ItemActionsVariablesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActionsVariablesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByName gets an item from the github.com/octokit/go-sdk/github/octokit/.orgs.item.actions.variables.item collection
func (m *ItemActionsVariablesRequestBuilder) ByName(name string)(*ItemActionsVariablesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewItemActionsVariablesWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsVariablesRequestBuilderInternal instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemActionsVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsVariablesRequestBuilder) {
    m := &ItemActionsVariablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/variables{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemActionsVariablesRequestBuilder instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemActionsVariablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsVariablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsVariablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all organization variables.You must authenticate using an access token with the `admin:org` scope to use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `organization_actions_variables:read` organization permission to use this endpoint. Authenticated users must have collaborator access to a repository to create, update, or read variables.
// Deprecated: This method is obsolete. Use GetAsVariablesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-organization-variables
func (m *ItemActionsVariablesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActionsVariablesRequestBuilderGetRequestConfiguration)(ItemActionsVariablesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsVariablesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsVariablesResponseable), nil
}
// GetAsVariablesGetResponse lists all organization variables.You must authenticate using an access token with the `admin:org` scope to use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `organization_actions_variables:read` organization permission to use this endpoint. Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-organization-variables
func (m *ItemActionsVariablesRequestBuilder) GetAsVariablesGetResponse(ctx context.Context, requestConfiguration *ItemActionsVariablesRequestBuilderGetRequestConfiguration)(ItemActionsVariablesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsVariablesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsVariablesGetResponseable), nil
}
// Post creates an organization variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `organization_actions_variables:write` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#create-an-organization-variable
func (m *ItemActionsVariablesRequestBuilder) Post(ctx context.Context, body ItemActionsVariablesPostRequestBodyable, requestConfiguration *ItemActionsVariablesRequestBuilderPostRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.EmptyObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateEmptyObjectFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.EmptyObjectable), nil
}
// ToGetRequestInformation lists all organization variables.You must authenticate using an access token with the `admin:org` scope to use this endpoint. If the repository is private, you must use an access token with the `repo` scope. GitHub Apps must have the `organization_actions_variables:read` organization permission to use this endpoint. Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemActionsVariablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActionsVariablesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates an organization variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `admin:org` scope to use this endpoint.If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have the `organization_actions_variables:write` organization permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemActionsVariablesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemActionsVariablesPostRequestBodyable, requestConfiguration *ItemActionsVariablesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemActionsVariablesRequestBuilder) WithUrl(rawUrl string)(*ItemActionsVariablesRequestBuilder) {
    return NewItemActionsVariablesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
