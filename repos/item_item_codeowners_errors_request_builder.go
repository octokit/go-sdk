package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeownersErrorsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\codeowners\errors
type ItemItemCodeownersErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeownersErrorsRequestBuilderGetQueryParameters list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
type ItemItemCodeownersErrorsRequestBuilderGetQueryParameters struct {
    // A branch, tag or commit name used to determine which version of the CODEOWNERS file to use. Default: the repository's default branch (e.g. `main`)
    Ref *string `uriparametername:"ref"`
}
// ItemItemCodeownersErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeownersErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemCodeownersErrorsRequestBuilderGetQueryParameters
}
// NewItemItemCodeownersErrorsRequestBuilderInternal instantiates a new ErrorsRequestBuilder and sets the default values.
func NewItemItemCodeownersErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersErrorsRequestBuilder) {
    m := &ItemItemCodeownersErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/codeowners/errors{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemCodeownersErrorsRequestBuilder instantiates a new ErrorsRequestBuilder and sets the default values.
func NewItemItemCodeownersErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeownersErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeownersErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/repos#list-codeowners-errors
func (m *ItemItemCodeownersErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeownersErrorsRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeownersErrorsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodeownersErrorsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodeownersErrorsable), nil
}
// ToGetRequestInformation list any syntax errors that are detected in the CODEOWNERSfile.For more information about the correct CODEOWNERS syntax,see "[About code owners](https://docs.github.com/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)."
func (m *ItemItemCodeownersErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeownersErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodeownersErrorsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeownersErrorsRequestBuilder) {
    return NewItemItemCodeownersErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
