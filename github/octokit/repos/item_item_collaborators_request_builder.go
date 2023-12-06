package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    ie8934dc2f1b24917e110bacbca0d06ccb93e9a2b121d984beb1453db51fbd5ac "github.com/octokit/go-sdk/github/octokit/repos/item/item/collaborators"
)

// ItemItemCollaboratorsRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\collaborators
type ItemItemCollaboratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCollaboratorsRequestBuilderGetQueryParameters for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
type ItemItemCollaboratorsRequestBuilderGetQueryParameters struct {
    // Filter collaborators returned by their affiliation. `outside` means all outside collaborators of an organization-owned repository. `direct` means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. `all` means all collaborators the authenticated user can see.
    // Deprecated: This property is deprecated, use affiliationAsGetAffiliationQueryParameterType instead
    Affiliation *string `uriparametername:"affiliation"`
    // Filter collaborators returned by their affiliation. `outside` means all outside collaborators of an organization-owned repository. `direct` means all collaborators with permissions to an organization-owned repository, regardless of organization membership status. `all` means all collaborators the authenticated user can see.
    AffiliationAsGetAffiliationQueryParameterType *ie8934dc2f1b24917e110bacbca0d06ccb93e9a2b121d984beb1453db51fbd5ac.GetAffiliationQueryParameterType `uriparametername:"affiliation"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Filter collaborators by the permissions they have on the repository. If not specified, all collaborators will be returned.
    // Deprecated: This property is deprecated, use permissionAsGetPermissionQueryParameterType instead
    Permission *string `uriparametername:"permission"`
    // Filter collaborators by the permissions they have on the repository. If not specified, all collaborators will be returned.
    PermissionAsGetPermissionQueryParameterType *ie8934dc2f1b24917e110bacbca0d06ccb93e9a2b121d984beb1453db51fbd5ac.GetPermissionQueryParameterType `uriparametername:"permission"`
}
// ItemItemCollaboratorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCollaboratorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCollaboratorsRequestBuilderGetQueryParameters
}
// ByUsername gets an item from the github.com/octokit/go-sdk/github/octokit/.repos.item.item.collaborators.item collection
func (m *ItemItemCollaboratorsRequestBuilder) ByUsername(username string)(*ItemItemCollaboratorsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemItemCollaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCollaboratorsRequestBuilderInternal instantiates a new CollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    m := &ItemItemCollaboratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/collaborators{?affiliation*,permission*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemCollaboratorsRequestBuilder instantiates a new CollaboratorsRequestBuilder and sets the default values.
func NewItemItemCollaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCollaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCollaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/collaborators#list-repository-collaborators
func (m *ItemItemCollaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCollaboratorsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Collaboratorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateCollaboratorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Collaboratorable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Collaboratorable)
        }
    }
    return val, nil
}
// ToGetRequestInformation for organization-owned repositories, the list of collaborators includes outside collaborators, organization members that are direct collaborators, organization members with access through team memberships, organization members with access through default organization permissions, and organization owners.Organization members with write, maintain, or admin privileges on the organization-owned repository can use this endpoint.Team members will include the members of child teams.You must authenticate using an access token with the `read:org` and `repo` scopes with push access to use thisendpoint. GitHub Apps must have the `members` organization permission and `metadata` repository permission to use thisendpoint.
func (m *ItemItemCollaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCollaboratorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCollaboratorsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCollaboratorsRequestBuilder) {
    return NewItemItemCollaboratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
