package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic047fde0792c6a665e244fff56a235f70f2297ba9dea40a69e5d95a7f73e9b2b "github.com/octokit/go-sdk/repos/item/item/securityadvisories"
)

// ItemItemSecurityAdvisoriesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\security-advisories
type ItemItemSecurityAdvisoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecurityAdvisoriesRequestBuilderGetQueryParameters lists security advisories in a repository.You must authenticate using an access token with the `repo` scope or `repository_advisories:read` permissionin order to get published security advisories in a private repository, or any unpublished security advisories that you have access to.You can access unpublished security advisories from a repository if you are a security manager or administrator of that repository, or if you are a collaborator on any security advisory.
type ItemItemSecurityAdvisoriesRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor.
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor.
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *ic047fde0792c6a665e244fff56a235f70f2297ba9dea40a69e5d95a7f73e9b2b.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Number of advisories to return per page.
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *ic047fde0792c6a665e244fff56a235f70f2297ba9dea40a69e5d95a7f73e9b2b.GetSortQueryParameterType `uriparametername:"sort"`
    // Filter by state of the repository advisories. Only advisories of this state will be returned.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Filter by state of the repository advisories. Only advisories of this state will be returned.
    StateAsGetStateQueryParameterType *ic047fde0792c6a665e244fff56a235f70f2297ba9dea40a69e5d95a7f73e9b2b.GetStateQueryParameterType `uriparametername:"state"`
}
// ItemItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemSecurityAdvisoriesRequestBuilderGetQueryParameters
}
// ItemItemSecurityAdvisoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGhsa_id gets an item from the go-sdk.repos.item.item.securityAdvisories.item collection
func (m *ItemItemSecurityAdvisoriesRequestBuilder) ByGhsa_id(ghsa_id string)(*ItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ghsa_id != "" {
        urlTplParams["ghsa_id"] = ghsa_id
    }
    return NewItemItemSecurityAdvisoriesWithGhsa_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemSecurityAdvisoriesRequestBuilderInternal instantiates a new SecurityAdvisoriesRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesRequestBuilder) {
    m := &ItemItemSecurityAdvisoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/security-advisories{?direction*,sort*,before*,after*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewItemItemSecurityAdvisoriesRequestBuilder instantiates a new SecurityAdvisoriesRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecurityAdvisoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists security advisories in a repository.You must authenticate using an access token with the `repo` scope or `repository_advisories:read` permissionin order to get published security advisories in a private repository, or any unpublished security advisories that you have access to.You can access unpublished security advisories from a repository if you are a security manager or administrator of that repository, or if you are a collaborator on any security advisory.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#list-repository-security-advisories
func (m *ItemItemSecurityAdvisoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryable)
        }
    }
    return val, nil
}
// Post creates a new repository security advisory.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to create a draft repository security advisory, you must be a security manager or administrator of that repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#create-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesRequestBuilder) Post(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryCreateable, requestConfiguration *ItemItemSecurityAdvisoriesRequestBuilderPostRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryable), nil
}
// Reports the reports property
func (m *ItemItemSecurityAdvisoriesRequestBuilder) Reports()(*ItemItemSecurityAdvisoriesReportsRequestBuilder) {
    return NewItemItemSecurityAdvisoriesReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists security advisories in a repository.You must authenticate using an access token with the `repo` scope or `repository_advisories:read` permissionin order to get published security advisories in a private repository, or any unpublished security advisories that you have access to.You can access unpublished security advisories from a repository if you are a security manager or administrator of that repository, or if you are a collaborator on any security advisory.
func (m *ItemItemSecurityAdvisoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation creates a new repository security advisory.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to create a draft repository security advisory, you must be a security manager or administrator of that repository.
func (m *ItemItemSecurityAdvisoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryAdvisoryCreateable, requestConfiguration *ItemItemSecurityAdvisoriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemSecurityAdvisoriesRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecurityAdvisoriesRequestBuilder) {
    return NewItemItemSecurityAdvisoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
