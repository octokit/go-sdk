package issues

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// IssuesRequestBuilder builds and executes requests for operations under \issues
type IssuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IssuesRequestBuilderGetQueryParameters list issues assigned to the authenticated user across all visible repositories including owned repositories, memberrepositories, and organization repositories. You can use the `filter` query parameter to fetch issues that are notnecessarily assigned to you.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/pulls/pulls#list-pull-requests)" endpoint.
type IssuesRequestBuilderGetQueryParameters struct {
    // 
    Collab *bool `uriparametername:"collab"`
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *GetDirectionQueryParameterType `uriparametername:"direction"`
    // Indicates which sorts of issues to return. `assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or `repos` means all issues you can see, regardless of participation or creation.
    // Deprecated: This property is deprecated, use filterAsGetFilterQueryParameterType instead
    Filter *string `uriparametername:"filter"`
    // Indicates which sorts of issues to return. `assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or `repos` means all issues you can see, regardless of participation or creation.
    FilterAsGetFilterQueryParameterType *GetFilterQueryParameterType `uriparametername:"filter"`
    // A list of comma separated label names. Example: `bug,ui,@high`
    Labels *string `uriparametername:"labels"`
    // 
    Orgs *bool `uriparametername:"orgs"`
    // 
    Owned *bool `uriparametername:"owned"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // 
    Pulls *bool `uriparametername:"pulls"`
    // Only show results that were last updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Since *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"since"`
    // What to sort results by.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // What to sort results by.
    SortAsGetSortQueryParameterType *GetSortQueryParameterType `uriparametername:"sort"`
    // Indicates the state of the issues to return.
    // Deprecated: This property is deprecated, use stateAsGetStateQueryParameterType instead
    State *string `uriparametername:"state"`
    // Indicates the state of the issues to return.
    StateAsGetStateQueryParameterType *GetStateQueryParameterType `uriparametername:"state"`
}
// IssuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IssuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IssuesRequestBuilderGetQueryParameters
}
// NewIssuesRequestBuilderInternal instantiates a new IssuesRequestBuilder and sets the default values.
func NewIssuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssuesRequestBuilder) {
    m := &IssuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/issues{?filter*,state*,labels*,sort*,direction*,since*,collab*,orgs*,owned*,pulls*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewIssuesRequestBuilder instantiates a new IssuesRequestBuilder and sets the default values.
func NewIssuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIssuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list issues assigned to the authenticated user across all visible repositories including owned repositories, memberrepositories, and organization repositories. You can use the `filter` query parameter to fetch issues that are notnecessarily assigned to you.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/pulls/pulls#list-pull-requests)" endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/issues#list-issues-assigned-to-the-authenticated-user
func (m *IssuesRequestBuilder) Get(ctx context.Context, requestConfiguration *IssuesRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Issueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Issueable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Issueable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list issues assigned to the authenticated user across all visible repositories including owned repositories, memberrepositories, and organization repositories. You can use the `filter` query parameter to fetch issues that are notnecessarily assigned to you.**Note**: GitHub's REST API considers every pull request an issue, but not every issue is a pull request. For thisreason, "Issues" endpoints may return both issues and pull requests in the response. You can identify pull requests bythe `pull_request` key. Be aware that the `id` of a pull request returned from "Issues" endpoints will be an _issue id_. To find out the pullrequest id, use the "[List pull requests](https://docs.github.com/rest/pulls/pulls#list-pull-requests)" endpoint.
func (m *IssuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IssuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IssuesRequestBuilder) WithUrl(rawUrl string)(*IssuesRequestBuilder) {
    return NewIssuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
