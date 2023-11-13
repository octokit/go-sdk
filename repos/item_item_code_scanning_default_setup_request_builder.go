package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningDefaultSetupRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\default-setup
type ItemItemCodeScanningDefaultSetupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningDefaultSetupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningDefaultSetupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemCodeScanningDefaultSetupRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningDefaultSetupRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCodeScanningDefaultSetupRequestBuilderInternal instantiates a new DefaultSetupRequestBuilder and sets the default values.
func NewItemItemCodeScanningDefaultSetupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningDefaultSetupRequestBuilder) {
    m := &ItemItemCodeScanningDefaultSetupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/code-scanning/default-setup", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningDefaultSetupRequestBuilder instantiates a new DefaultSetupRequestBuilder and sets the default values.
func NewItemItemCodeScanningDefaultSetupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningDefaultSetupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningDefaultSetupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a code scanning default setup configuration.You must use an access token with the `repo` scope to use this endpoint with private repos or the `public_repo`scope for public repos.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-a-code-scanning-default-setup-configuration
func (m *ItemItemCodeScanningDefaultSetupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningDefaultSetupRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningDefaultSetupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "503": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodeScanningDefaultSetup503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodeScanningDefaultSetupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningDefaultSetupable), nil
}
// Patch updates a code scanning default setup configuration.You must use an access token with the `repo` scope to use this endpoint with private repos or the `public_repo`scope for public repos.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#update-a-code-scanning-default-setup-configuration
func (m *ItemItemCodeScanningDefaultSetupRequestBuilder) Patch(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningDefaultSetupUpdateable, requestConfiguration *ItemItemCodeScanningDefaultSetupRequestBuilderPatchRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "409": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "503": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateEmptyObject503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateEmptyObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.EmptyObjectable), nil
}
// ToGetRequestInformation gets a code scanning default setup configuration.You must use an access token with the `repo` scope to use this endpoint with private repos or the `public_repo`scope for public repos.
func (m *ItemItemCodeScanningDefaultSetupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningDefaultSetupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation updates a code scanning default setup configuration.You must use an access token with the `repo` scope to use this endpoint with private repos or the `public_repo`scope for public repos.
func (m *ItemItemCodeScanningDefaultSetupRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeScanningDefaultSetupUpdateable, requestConfiguration *ItemItemCodeScanningDefaultSetupRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCodeScanningDefaultSetupRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningDefaultSetupRequestBuilder) {
    return NewItemItemCodeScanningDefaultSetupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
