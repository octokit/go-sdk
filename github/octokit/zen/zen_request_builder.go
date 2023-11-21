package zen

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ZenRequestBuilder builds and executes requests for operations under \zen
type ZenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZenRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZenRequestBuilderInternal instantiates a new ZenRequestBuilder and sets the default values.
func NewZenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZenRequestBuilder) {
    m := &ZenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/zen", pathParameters),
    }
    return m
}
// NewZenRequestBuilder instantiates a new ZenRequestBuilder and sets the default values.
func NewZenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZenRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a random sentence from the Zen of GitHub
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#get-the-zen-of-github
func (m *ZenRequestBuilder) Get(ctx context.Context, requestConfiguration *ZenRequestBuilderGetRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// ToGetRequestInformation get a random sentence from the Zen of GitHub
func (m *ZenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ZenRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
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
func (m *ZenRequestBuilder) WithUrl(rawUrl string)(*ZenRequestBuilder) {
    return NewZenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
