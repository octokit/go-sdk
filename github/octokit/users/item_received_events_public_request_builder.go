package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemReceived_eventsPublicRequestBuilder builds and executes requests for operations under \users\{username}\received_events\public
type ItemReceived_eventsPublicRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemReceived_eventsPublicRequestBuilderGetQueryParameters list public events received by a user
type ItemReceived_eventsPublicRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemReceived_eventsPublicRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemReceived_eventsPublicRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemReceived_eventsPublicRequestBuilderGetQueryParameters
}
// NewItemReceived_eventsPublicRequestBuilderInternal instantiates a new PublicRequestBuilder and sets the default values.
func NewItemReceived_eventsPublicRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReceived_eventsPublicRequestBuilder) {
    m := &ItemReceived_eventsPublicRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/received_events/public{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemReceived_eventsPublicRequestBuilder instantiates a new PublicRequestBuilder and sets the default values.
func NewItemReceived_eventsPublicRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemReceived_eventsPublicRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemReceived_eventsPublicRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list public events received by a user
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/events#list-public-events-received-by-a-user
func (m *ItemReceived_eventsPublicRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemReceived_eventsPublicRequestBuilderGetRequestConfiguration)([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateEventFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Eventable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Eventable)
        }
    }
    return val, nil
}
func (m *ItemReceived_eventsPublicRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemReceived_eventsPublicRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemReceived_eventsPublicRequestBuilder) WithUrl(rawUrl string)(*ItemReceived_eventsPublicRequestBuilder) {
    return NewItemReceived_eventsPublicRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
