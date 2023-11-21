package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemPropertiesValuesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\properties\values
type ItemItemPropertiesValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemPropertiesValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemPropertiesValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemPropertiesValuesRequestBuilderInternal instantiates a new ValuesRequestBuilder and sets the default values.
func NewItemItemPropertiesValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPropertiesValuesRequestBuilder) {
    m := &ItemItemPropertiesValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/properties/values", pathParameters),
    }
    return m
}
// NewItemItemPropertiesValuesRequestBuilder instantiates a new ValuesRequestBuilder and sets the default values.
func NewItemItemPropertiesValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemPropertiesValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemPropertiesValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets all custom property values that are set for a repository.Users with read access to the repository can use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/custom-properties#get-all-custom-property-values-for-a-repository
func (m *ItemItemPropertiesValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemPropertiesValuesRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CustomPropertyValueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCustomPropertyValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CustomPropertyValueable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CustomPropertyValueable)
        }
    }
    return val, nil
}
// ToGetRequestInformation gets all custom property values that are set for a repository.Users with read access to the repository can use this endpoint.
func (m *ItemItemPropertiesValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemPropertiesValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemPropertiesValuesRequestBuilder) WithUrl(rawUrl string)(*ItemItemPropertiesValuesRequestBuilder) {
    return NewItemItemPropertiesValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
