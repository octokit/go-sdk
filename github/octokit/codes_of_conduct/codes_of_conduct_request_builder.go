package codes_of_conduct

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Codes_of_conductRequestBuilder builds and executes requests for operations under \codes_of_conduct
type Codes_of_conductRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Codes_of_conductRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Codes_of_conductRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByKey gets an item from the octokit.codes_of_conduct.item collection
func (m *Codes_of_conductRequestBuilder) ByKey(key string)(*WithKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if key != "" {
        urlTplParams["key"] = key
    }
    return NewWithKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodes_of_conductRequestBuilderInternal instantiates a new Codes_of_conductRequestBuilder and sets the default values.
func NewCodes_of_conductRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Codes_of_conductRequestBuilder) {
    m := &Codes_of_conductRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/codes_of_conduct", pathParameters),
    }
    return m
}
// NewCodes_of_conductRequestBuilder instantiates a new Codes_of_conductRequestBuilder and sets the default values.
func NewCodes_of_conductRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Codes_of_conductRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodes_of_conductRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns array of all GitHub's codes of conduct.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codes-of-conduct/codes-of-conduct#get-all-codes-of-conduct
func (m *Codes_of_conductRequestBuilder) Get(ctx context.Context, requestConfiguration *Codes_of_conductRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeOfConductable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodeOfConductFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeOfConductable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeOfConductable)
        }
    }
    return val, nil
}
// ToGetRequestInformation returns array of all GitHub's codes of conduct.
func (m *Codes_of_conductRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Codes_of_conductRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *Codes_of_conductRequestBuilder) WithUrl(rawUrl string)(*Codes_of_conductRequestBuilder) {
    return NewCodes_of_conductRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
