package users

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemEventsOrgsWithOrgItemRequestBuilder builds and executes requests for operations under \users\{username}\events\orgs\{org}
type ItemEventsOrgsWithOrgItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters this is the user's organization dashboard. You must be authenticated as the user to view this.
type ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsOrgsWithOrgItemRequestBuilderGetQueryParameters
}
// NewItemEventsOrgsWithOrgItemRequestBuilderInternal instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    m := &ItemEventsOrgsWithOrgItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/events/orgs/{org}{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemEventsOrgsWithOrgItemRequestBuilder instantiates a new WithOrgItemRequestBuilder and sets the default values.
func NewItemEventsOrgsWithOrgItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsOrgsWithOrgItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get this is the user's organization dashboard. You must be authenticated as the user to view this.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/events#list-organization-events-for-the-authenticated-user
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Eventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Eventable)
        }
    }
    return val, nil
}
// ToGetRequestInformation this is the user's organization dashboard. You must be authenticated as the user to view this.
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsOrgsWithOrgItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemEventsOrgsWithOrgItemRequestBuilder) WithUrl(rawUrl string)(*ItemEventsOrgsWithOrgItemRequestBuilder) {
    return NewItemEventsOrgsWithOrgItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
