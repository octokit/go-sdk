package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    id112e634f853a7e7cb3af8e61b7bbc279dff305bc623195b7337e2e6543c47d3 "github.com/octokit/go-sdk/github/octokit/orgs/item/outside_collaborators"
)

// ItemOutside_collaboratorsRequestBuilder builds and executes requests for operations under \orgs\{org}\outside_collaborators
type ItemOutside_collaboratorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutside_collaboratorsRequestBuilderGetQueryParameters list all users who are outside collaborators of an organization.
type ItemOutside_collaboratorsRequestBuilderGetQueryParameters struct {
    // Filter the list of outside collaborators. `2fa_disabled` means that only outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned.
    // Deprecated: This property is deprecated, use filterAsGetFilterQueryParameterType instead
    Filter *string `uriparametername:"filter"`
    // Filter the list of outside collaborators. `2fa_disabled` means that only outside collaborators without [two-factor authentication](https://github.com/blog/1614-two-factor-authentication) enabled will be returned.
    FilterAsGetFilterQueryParameterType *id112e634f853a7e7cb3af8e61b7bbc279dff305bc623195b7337e2e6543c47d3.GetFilterQueryParameterType `uriparametername:"filter"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemOutside_collaboratorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutside_collaboratorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutside_collaboratorsRequestBuilderGetQueryParameters
}
// ByUsername gets an item from the github.com/octokit/go-sdk/github/octokit/.orgs.item.outside_collaborators.item collection
func (m *ItemOutside_collaboratorsRequestBuilder) ByUsername(username string)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if username != "" {
        urlTplParams["username"] = username
    }
    return NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOutside_collaboratorsRequestBuilderInternal instantiates a new Outside_collaboratorsRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsRequestBuilder) {
    m := &ItemOutside_collaboratorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/outside_collaborators{?filter*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemOutside_collaboratorsRequestBuilder instantiates a new Outside_collaboratorsRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutside_collaboratorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all users who are outside collaborators of an organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/outside-collaborators#list-outside-collaborators-for-an-organization
func (m *ItemOutside_collaboratorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutside_collaboratorsRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateSimpleUserFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all users who are outside collaborators of an organization.
func (m *ItemOutside_collaboratorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutside_collaboratorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOutside_collaboratorsRequestBuilder) WithUrl(rawUrl string)(*ItemOutside_collaboratorsRequestBuilder) {
    return NewItemOutside_collaboratorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
