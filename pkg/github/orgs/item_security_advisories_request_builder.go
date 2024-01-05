package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i73af498fc508ab66c6d7b9e1987fbbf6febbe066a7dc653ccf219a9638b2bc60 "github.com/octokit/go-sdk/pkg/github/orgs/item/securityadvisories"
)

// ItemSecurityAdvisoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\security-advisories
type ItemSecurityAdvisoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityAdvisoriesRequestBuilderGetQueryParameters lists repository security advisories for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `repository_advisories:write` permission.
type ItemSecurityAdvisoriesRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i73af498fc508ab66c6d7b9e1987fbbf6febbe066a7dc653ccf219a9638b2bc60.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The number of advisories to return per page. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by.
    SortAsGetSortQueryParameterType *i73af498fc508ab66c6d7b9e1987fbbf6febbe066a7dc653ccf219a9638b2bc60.GetSortQueryParameterType `uriparametername:"sort"`
    // Filter by the state of the repository advisories. Only advisories of this state will be returned.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Filter by the state of the repository advisories. Only advisories of this state will be returned.
    StateAsGetStateQueryParameterType *i73af498fc508ab66c6d7b9e1987fbbf6febbe066a7dc653ccf219a9638b2bc60.GetStateQueryParameterType `uriparametername:"state"`
}
// ItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSecurityAdvisoriesRequestBuilderGetQueryParameters
}
// NewItemSecurityAdvisoriesRequestBuilderInternal instantiates a new SecurityAdvisoriesRequestBuilder and sets the default values.
func NewItemSecurityAdvisoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityAdvisoriesRequestBuilder) {
    m := &ItemSecurityAdvisoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/security-advisories{?direction*,sort*,before*,after*,per_page*,state*}", pathParameters),
    }
    return m
}
// NewItemSecurityAdvisoriesRequestBuilder instantiates a new SecurityAdvisoriesRequestBuilder and sets the default values.
func NewItemSecurityAdvisoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityAdvisoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityAdvisoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists repository security advisories for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `repository_advisories:write` permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#list-repository-security-advisories-for-an-organization
func (m *ItemSecurityAdvisoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRepositoryAdvisoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RepositoryAdvisoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists repository security advisories for an organization.To use this endpoint, you must be an owner or security manager for the organization, and you must use an access token with the `repo` scope or `repository_advisories:write` permission.
func (m *ItemSecurityAdvisoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityAdvisoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSecurityAdvisoriesRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityAdvisoriesRequestBuilder) {
    return NewItemSecurityAdvisoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
