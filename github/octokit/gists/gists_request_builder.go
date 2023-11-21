package gists

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GistsRequestBuilder builds and executes requests for operations under \gists
type GistsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GistsRequestBuilderGetQueryParameters lists the authenticated user's gists or if called anonymously, this endpoint returns all public gists:
type GistsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // Only show results that were last updated after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Since *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"since"`
}
// GistsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GistsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GistsRequestBuilderGetQueryParameters
}
// GistsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GistsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGist_id gets an item from the octokit.gists.item collection
func (m *GistsRequestBuilder) ByGist_id(gist_id string)(*WithGist_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if gist_id != "" {
        urlTplParams["gist_id"] = gist_id
    }
    return NewWithGist_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGistsRequestBuilderInternal instantiates a new GistsRequestBuilder and sets the default values.
func NewGistsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GistsRequestBuilder) {
    m := &GistsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gists{?since*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewGistsRequestBuilder instantiates a new GistsRequestBuilder and sets the default values.
func NewGistsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GistsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGistsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the authenticated user's gists or if called anonymously, this endpoint returns all public gists:
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/gists/gists#list-gists-for-the-authenticated-user
func (m *GistsRequestBuilder) Get(ctx context.Context, requestConfiguration *GistsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.BaseGistable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBaseGistFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.BaseGistable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.BaseGistable)
        }
    }
    return val, nil
}
// Post allows you to add a new gist with one or more files.**Note:** Don't name your files "gistfile" with a numerical suffix. This is the format of the automatic naming scheme that Gist uses internally.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/gists/gists#create-a-gist
func (m *GistsRequestBuilder) Post(ctx context.Context, body GistsPostRequestBodyable, requestConfiguration *GistsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GistSimpleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateGistSimpleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GistSimpleable), nil
}
// Public the public property
func (m *GistsRequestBuilder) Public()(*PublicRequestBuilder) {
    return NewPublicRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Starred the starred property
func (m *GistsRequestBuilder) Starred()(*StarredRequestBuilder) {
    return NewStarredRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists the authenticated user's gists or if called anonymously, this endpoint returns all public gists:
func (m *GistsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GistsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation allows you to add a new gist with one or more files.**Note:** Don't name your files "gistfile" with a numerical suffix. This is the format of the automatic naming scheme that Gist uses internally.
func (m *GistsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GistsPostRequestBodyable, requestConfiguration *GistsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GistsRequestBuilder) WithUrl(rawUrl string)(*GistsRequestBuilder) {
    return NewGistsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
