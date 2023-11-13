package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4e36cc172de8f3b029128f4479100b6a7803b1481c1d97dd13f03f920709966e "github.com/octokit/go-sdk/orgs/item/repos"
)

// ItemReposRequestBuilder builds and executes requests for operations under \orgs\{org}\repos
type ItemReposRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReposRequestBuilderGetQueryParameters lists repositories for the specified organization.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
type ItemReposRequestBuilderGetQueryParameters struct {
    // The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The order to sort by. Default: `asc` when using `full_name`, otherwise `desc`.
    DirectionAsGetDirectionQueryParameterType *i4e36cc172de8f3b029128f4479100b6a7803b1481c1d97dd13f03f920709966e.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *i4e36cc172de8f3b029128f4479100b6a7803b1481c1d97dd13f03f920709966e.GetSortQueryParameterType `uriparametername:"sort"`
    // Specifies the types of repositories you want returned.
    // Deprecated: This property is deprecated, use typeAsGetTypeQueryParameterType instead
    Type *string `uriparametername:"type"`
    // Specifies the types of repositories you want returned.
    TypeAsGetTypeQueryParameterType *i4e36cc172de8f3b029128f4479100b6a7803b1481c1d97dd13f03f920709966e.GetTypeQueryParameterType `uriparametername:"type"`
}
// ItemReposRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReposRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemReposRequestBuilderGetQueryParameters
}
// ItemReposRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReposRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemReposRequestBuilderInternal instantiates a new ReposRequestBuilder and sets the default values.
func NewItemReposRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposRequestBuilder) {
    m := &ItemReposRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/repos{?type*,sort*,direction*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemReposRequestBuilder instantiates a new ReposRequestBuilder and sets the default values.
func NewItemReposRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReposRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReposRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repositories for the specified organization.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#list-organization-repositories
func (m *ItemReposRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemReposRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMinimalRepositoryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)
        }
    }
    return val, nil
}
// Post creates a new repository in the specified organization. The authenticated user must be a member of the organization.**OAuth scope requirements**When using [OAuth](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), authorizations must include:*   `public_repo` scope or `repo` scope to create a public repository. Note: For GitHub AE, use `repo` scope to create an internal repository.*   `repo` scope to create a private repository
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#create-an-organization-repository
func (m *ItemReposRequestBuilder) Post(ctx context.Context, body ItemReposPostRequestBodyable, requestConfiguration *ItemReposRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Repositoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Repositoryable), nil
}
// ToGetRequestInformation lists repositories for the specified organization.**Note:** In order to see the `security_and_analysis` block for a repository you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."
func (m *ItemReposRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemReposRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a new repository in the specified organization. The authenticated user must be a member of the organization.**OAuth scope requirements**When using [OAuth](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/), authorizations must include:*   `public_repo` scope or `repo` scope to create a public repository. Note: For GitHub AE, use `repo` scope to create an internal repository.*   `repo` scope to create a private repository
func (m *ItemReposRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemReposPostRequestBodyable, requestConfiguration *ItemReposRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemReposRequestBuilder) WithUrl(rawUrl string)(*ItemReposRequestBuilder) {
    return NewItemReposRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
