package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTeamsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\teams
type ItemItemTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTeamsRequestBuilderGetQueryParameters lists the teams that have access to the specified repository and that are also visible to the authenticated user.For a public repository, a team is listed only if that team added the public repository explicitly.Personal access tokens require the following scopes:* `public_repo` to call this endpoint on a public repository* `repo` to call this endpoint on a private repository (this scope also includes public repositories)This endpoint is not compatible with fine-grained personal access tokens.
type ItemItemTeamsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemTeamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTeamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemTeamsRequestBuilderGetQueryParameters
}
// NewItemItemTeamsRequestBuilderInternal instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemItemTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTeamsRequestBuilder) {
    m := &ItemItemTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/teams{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemTeamsRequestBuilder instantiates a new TeamsRequestBuilder and sets the default values.
func NewItemItemTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the teams that have access to the specified repository and that are also visible to the authenticated user.For a public repository, a team is listed only if that team added the public repository explicitly.Personal access tokens require the following scopes:* `public_repo` to call this endpoint on a public repository* `repo` to call this endpoint on a private repository (this scope also includes public repositories)This endpoint is not compatible with fine-grained personal access tokens.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#list-repository-teams
func (m *ItemItemTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTeamsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Teamable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the teams that have access to the specified repository and that are also visible to the authenticated user.For a public repository, a team is listed only if that team added the public repository explicitly.Personal access tokens require the following scopes:* `public_repo` to call this endpoint on a public repository* `repo` to call this endpoint on a private repository (this scope also includes public repositories)This endpoint is not compatible with fine-grained personal access tokens.
func (m *ItemItemTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTeamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemTeamsRequestBuilder) WithUrl(rawUrl string)(*ItemItemTeamsRequestBuilder) {
    return NewItemItemTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
