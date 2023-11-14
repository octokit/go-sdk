package emojis

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmojisRequestBuilder builds and executes requests for operations under \emojis
type EmojisRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmojisRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmojisRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmojisRequestBuilderInternal instantiates a new EmojisRequestBuilder and sets the default values.
func NewEmojisRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmojisRequestBuilder) {
    m := &EmojisRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/emojis", pathParameters),
    }
    return m
}
// NewEmojisRequestBuilder instantiates a new EmojisRequestBuilder and sets the default values.
func NewEmojisRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmojisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmojisRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all the emojis available to use on GitHub.
// Deprecated: This method is obsolete. Use GetAsEmojisGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/emojis/emojis#get-emojis
func (m *EmojisRequestBuilder) Get(ctx context.Context, requestConfiguration *EmojisRequestBuilderGetRequestConfiguration)(EmojisResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmojisResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmojisResponseable), nil
}
// GetAsEmojisGetResponse lists all the emojis available to use on GitHub.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/emojis/emojis#get-emojis
func (m *EmojisRequestBuilder) GetAsEmojisGetResponse(ctx context.Context, requestConfiguration *EmojisRequestBuilderGetRequestConfiguration)(EmojisGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmojisGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmojisGetResponseable), nil
}
// ToGetRequestInformation lists all the emojis available to use on GitHub.
func (m *EmojisRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmojisRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *EmojisRequestBuilder) WithUrl(rawUrl string)(*EmojisRequestBuilder) {
    return NewEmojisRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
