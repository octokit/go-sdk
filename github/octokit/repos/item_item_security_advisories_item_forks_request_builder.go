package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemSecurityAdvisoriesItemForksRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\security-advisories\{ghsa_id}\forks
type ItemItemSecurityAdvisoriesItemForksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecurityAdvisoriesItemForksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesItemForksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemSecurityAdvisoriesItemForksRequestBuilderInternal instantiates a new ForksRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesItemForksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesItemForksRequestBuilder) {
    m := &ItemItemSecurityAdvisoriesItemForksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/security-advisories/{ghsa_id}/forks", pathParameters),
    }
    return m
}
// NewItemItemSecurityAdvisoriesItemForksRequestBuilder instantiates a new ForksRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesItemForksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesItemForksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecurityAdvisoriesItemForksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a temporary private fork to collaborate on fixing a security vulnerability in your repository.**Note**: Forking a repository happens asynchronously. You may have to wait up to 5 minutes before you can access the fork.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#create-a-temporary-private-fork
func (m *ItemItemSecurityAdvisoriesItemForksRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesItemForksRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.FullRepositoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateFullRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.FullRepositoryable), nil
}
// ToPostRequestInformation create a temporary private fork to collaborate on fixing a security vulnerability in your repository.**Note**: Forking a repository happens asynchronously. You may have to wait up to 5 minutes before you can access the fork.
func (m *ItemItemSecurityAdvisoriesItemForksRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesItemForksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemSecurityAdvisoriesItemForksRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecurityAdvisoriesItemForksRequestBuilder) {
    return NewItemItemSecurityAdvisoriesItemForksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
