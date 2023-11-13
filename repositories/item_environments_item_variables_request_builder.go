package repositories

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEnvironmentsItemVariablesRequestBuilder builds and executes requests for operations under \repositories\{repository_id}\environments\{environment_name}\variables
type ItemEnvironmentsItemVariablesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters lists all environment variables.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environments:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
type ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 30).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEnvironmentsItemVariablesRequestBuilderGetQueryParameters
}
// ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByName gets an item from the go-sdk.repositories.item.environments.item.variables.item collection
func (m *ItemEnvironmentsItemVariablesRequestBuilder) ByName(name string)(*ItemEnvironmentsItemVariablesWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewItemEnvironmentsItemVariablesWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEnvironmentsItemVariablesRequestBuilderInternal instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesRequestBuilder) {
    m := &ItemEnvironmentsItemVariablesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repositories/{repository_id}/environments/{environment_name}/variables{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemEnvironmentsItemVariablesRequestBuilder instantiates a new VariablesRequestBuilder and sets the default values.
func NewItemEnvironmentsItemVariablesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEnvironmentsItemVariablesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEnvironmentsItemVariablesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all environment variables.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environments:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// Deprecated: This method is obsolete. Use GetAsVariablesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-environment-variables
func (m *ItemEnvironmentsItemVariablesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration)(ItemEnvironmentsItemVariablesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemVariablesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemVariablesResponseable), nil
}
// GetAsVariablesGetResponse lists all environment variables.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environments:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#list-environment-variables
func (m *ItemEnvironmentsItemVariablesRequestBuilder) GetAsVariablesGetResponse(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration)(ItemEnvironmentsItemVariablesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemEnvironmentsItemVariablesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemEnvironmentsItemVariablesGetResponseable), nil
}
// Post create an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#create-an-environment-variable
func (m *ItemEnvironmentsItemVariablesRequestBuilder) Post(ctx context.Context, body ItemEnvironmentsItemVariablesPostRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateEmptyObjectFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable), nil
}
// ToGetRequestInformation lists all environment variables.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environments:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemEnvironmentsItemVariablesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an environment variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `environment:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemEnvironmentsItemVariablesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemEnvironmentsItemVariablesPostRequestBodyable, requestConfiguration *ItemEnvironmentsItemVariablesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemEnvironmentsItemVariablesRequestBuilder) WithUrl(rawUrl string)(*ItemEnvironmentsItemVariablesRequestBuilder) {
    return NewItemEnvironmentsItemVariablesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
