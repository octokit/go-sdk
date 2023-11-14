package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning\codeql\databases\{language}
type ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal instantiates a new WithLanguageItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    m := &ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/code-scanning/codeql/databases/{language}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder instantiates a new WithLanguageItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a CodeQL database for a language in a repository.By default this endpoint returns JSON metadata about the CodeQL database. Todownload the CodeQL database binary content, set the `Accept` header of the requestto [`application/zip`](https://docs.github.com/rest/overview/media-types), and make sureyour HTTP client is configured to follow redirects or use the `Location` headerto make a second request to get the redirect URL.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-a-codeql-database-for-a-repository
func (m *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningCodeqlDatabaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "503": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodeScanningCodeqlDatabase503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodeScanningCodeqlDatabaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodeScanningCodeqlDatabaseable), nil
}
// ToGetRequestInformation gets a CodeQL database for a language in a repository.By default this endpoint returns JSON metadata about the CodeQL database. Todownload the CodeQL database binary content, set the `Accept` header of the requestto [`application/zip`](https://docs.github.com/rest/overview/media-types), and make sureyour HTTP client is configured to follow redirects or use the `Location` headerto make a second request to get the redirect URL.For private repositories, you must use an access token with the `security_events` scope.For public repositories, you can use tokens with the `security_events` or `public_repo` scope.
func (m *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder) {
    return NewItemItemCodeScanningCodeqlDatabasesWithLanguageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
