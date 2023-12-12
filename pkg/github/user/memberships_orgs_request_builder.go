package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i6248a706a395d319f3e66611fac0b2654fe48b61e0e7152b109cc83e1d984972 "github.com/octokit/go-sdk/pkg/github/user/memberships/orgs"
)

// MembershipsOrgsRequestBuilder builds and executes requests for operations under \user\memberships\orgs
type MembershipsOrgsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MembershipsOrgsRequestBuilderGetQueryParameters lists all of the authenticated user's organization memberships.
type MembershipsOrgsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Indicates the state of the memberships to return. If not specified, the API returns both active and pending memberships.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Indicates the state of the memberships to return. If not specified, the API returns both active and pending memberships.
    StateAsGetStateQueryParameterType *i6248a706a395d319f3e66611fac0b2654fe48b61e0e7152b109cc83e1d984972.GetStateQueryParameterType `uriparametername:"state"`
}
// MembershipsOrgsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembershipsOrgsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MembershipsOrgsRequestBuilderGetQueryParameters
}
// ByOrg gets an item from the github.com/octokit/go-sdk/pkg/github/.user.memberships.orgs.item collection
func (m *MembershipsOrgsRequestBuilder) ByOrg(org string)(*MembershipsOrgsWithOrgItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if org != "" {
        urlTplParams["org"] = org
    }
    return NewMembershipsOrgsWithOrgItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMembershipsOrgsRequestBuilderInternal instantiates a new OrgsRequestBuilder and sets the default values.
func NewMembershipsOrgsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsOrgsRequestBuilder) {
    m := &MembershipsOrgsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/memberships/orgs{?state*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewMembershipsOrgsRequestBuilder instantiates a new OrgsRequestBuilder and sets the default values.
func NewMembershipsOrgsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembershipsOrgsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembershipsOrgsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all of the authenticated user's organization memberships.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/members#list-organization-memberships-for-the-authenticated-user
func (m *MembershipsOrgsRequestBuilder) Get(ctx context.Context, requestConfiguration *MembershipsOrgsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrgMembershipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrgMembershipable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all of the authenticated user's organization memberships.
func (m *MembershipsOrgsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MembershipsOrgsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MembershipsOrgsRequestBuilder) WithUrl(rawUrl string)(*MembershipsOrgsRequestBuilder) {
    return NewMembershipsOrgsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
