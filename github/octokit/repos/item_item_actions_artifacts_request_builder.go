package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsArtifactsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\artifacts
type ItemItemActionsArtifactsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsArtifactsRequestBuilderGetQueryParameters lists all artifacts for a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemActionsArtifactsRequestBuilderGetQueryParameters struct {
    // The name field of an artifact. When specified, only artifacts with this name will be returned.
    Name *string `uriparametername:"name"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemActionsArtifactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsArtifactsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsArtifactsRequestBuilderGetQueryParameters
}
// ByArtifact_id gets an item from the octokit.repos.item.item.actions.artifacts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsArtifactsRequestBuilder) ByArtifact_id(artifact_id string)(*ItemItemActionsArtifactsWithArtifact_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if artifact_id != "" {
        urlTplParams["artifact_id"] = artifact_id
    }
    return NewItemItemActionsArtifactsWithArtifact_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByArtifact_idInteger gets an item from the octokit.repos.item.item.actions.artifacts.item collection
func (m *ItemItemActionsArtifactsRequestBuilder) ByArtifact_idInteger(artifact_id int32)(*ItemItemActionsArtifactsWithArtifact_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["artifact_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(artifact_id), 10)
    return NewItemItemActionsArtifactsWithArtifact_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsArtifactsRequestBuilderInternal instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsRequestBuilder) {
    m := &ItemItemActionsArtifactsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/artifacts{?per_page*,page*,name*}", pathParameters),
    }
    return m
}
// NewItemItemActionsArtifactsRequestBuilder instantiates a new ArtifactsRequestBuilder and sets the default values.
func NewItemItemActionsArtifactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsArtifactsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsArtifactsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all artifacts for a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// Deprecated: This method is obsolete. Use GetAsArtifactsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/artifacts#list-artifacts-for-a-repository
func (m *ItemItemActionsArtifactsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsRequestBuilderGetRequestConfiguration)(ItemItemActionsArtifactsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsArtifactsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsArtifactsResponseable), nil
}
// GetAsArtifactsGetResponse lists all artifacts for a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/artifacts#list-artifacts-for-a-repository
func (m *ItemItemActionsArtifactsRequestBuilder) GetAsArtifactsGetResponse(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsRequestBuilderGetRequestConfiguration)(ItemItemActionsArtifactsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsArtifactsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsArtifactsGetResponseable), nil
}
// ToGetRequestInformation lists all artifacts for a repository. Anyone with read access to the repository can use this endpoint. If the repository is private you must use an access token with the `repo` scope. GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsArtifactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsArtifactsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsArtifactsRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsArtifactsRequestBuilder) {
    return NewItemItemActionsArtifactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
