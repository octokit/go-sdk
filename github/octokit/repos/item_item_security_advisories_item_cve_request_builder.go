package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemSecurityAdvisoriesItemCveRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\security-advisories\{ghsa_id}\cve
type ItemItemSecurityAdvisoriesItemCveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemSecurityAdvisoriesItemCveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemSecurityAdvisoriesItemCveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemSecurityAdvisoriesItemCveRequestBuilderInternal instantiates a new CveRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesItemCveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesItemCveRequestBuilder) {
    m := &ItemItemSecurityAdvisoriesItemCveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/security-advisories/{ghsa_id}/cve", pathParameters),
    }
    return m
}
// NewItemItemSecurityAdvisoriesItemCveRequestBuilder instantiates a new CveRequestBuilder and sets the default values.
func NewItemItemSecurityAdvisoriesItemCveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemSecurityAdvisoriesItemCveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemSecurityAdvisoriesItemCveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post if you want a CVE identification number for the security vulnerability in your project, and don't already have one, you can request a CVE identification number from GitHub. For more information see "[Requesting a CVE identification number](https://docs.github.com/code-security/security-advisories/repository-security-advisories/publishing-a-repository-security-advisory#requesting-a-cve-identification-number-optional)."You may request a CVE for public repositories, but cannot do so for private repositories.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to request a CVE for a repository security advisory, you must be a security manager or administrator of that repository.
// Deprecated: This method is obsolete. Use PostAsCvePostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#request-a-cve-for-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesItemCveRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesItemCveRequestBuilderPostRequestConfiguration)(ItemItemSecurityAdvisoriesItemCveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemSecurityAdvisoriesItemCveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemSecurityAdvisoriesItemCveResponseable), nil
}
// PostAsCvePostResponse if you want a CVE identification number for the security vulnerability in your project, and don't already have one, you can request a CVE identification number from GitHub. For more information see "[Requesting a CVE identification number](https://docs.github.com/code-security/security-advisories/repository-security-advisories/publishing-a-repository-security-advisory#requesting-a-cve-identification-number-optional)."You may request a CVE for public repositories, but cannot do so for private repositories.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to request a CVE for a repository security advisory, you must be a security manager or administrator of that repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/security-advisories/repository-advisories#request-a-cve-for-a-repository-security-advisory
func (m *ItemItemSecurityAdvisoriesItemCveRequestBuilder) PostAsCvePostResponse(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesItemCveRequestBuilderPostRequestConfiguration)(ItemItemSecurityAdvisoriesItemCvePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemSecurityAdvisoriesItemCvePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemSecurityAdvisoriesItemCvePostResponseable), nil
}
// ToPostRequestInformation if you want a CVE identification number for the security vulnerability in your project, and don't already have one, you can request a CVE identification number from GitHub. For more information see "[Requesting a CVE identification number](https://docs.github.com/code-security/security-advisories/repository-security-advisories/publishing-a-repository-security-advisory#requesting-a-cve-identification-number-optional)."You may request a CVE for public repositories, but cannot do so for private repositories.You must authenticate using an access token with the `repo` scope or `repository_advisories:write` permission to use this endpoint.In order to request a CVE for a repository security advisory, you must be a security manager or administrator of that repository.
func (m *ItemItemSecurityAdvisoriesItemCveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemSecurityAdvisoriesItemCveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemSecurityAdvisoriesItemCveRequestBuilder) WithUrl(rawUrl string)(*ItemItemSecurityAdvisoriesItemCveRequestBuilder) {
    return NewItemItemSecurityAdvisoriesItemCveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
