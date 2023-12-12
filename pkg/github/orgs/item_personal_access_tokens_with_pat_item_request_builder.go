package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemPersonalAccessTokensWithPat_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\personal-access-tokens\{pat_id}
type ItemPersonalAccessTokensWithPat_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPersonalAccessTokensWithPat_ItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPersonalAccessTokensWithPat_ItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPersonalAccessTokensWithPat_ItemRequestBuilderInternal instantiates a new WithPat_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokensWithPat_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokensWithPat_ItemRequestBuilder) {
    m := &ItemPersonalAccessTokensWithPat_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/personal-access-tokens/{pat_id}", pathParameters),
    }
    return m
}
// NewItemPersonalAccessTokensWithPat_ItemRequestBuilder instantiates a new WithPat_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokensWithPat_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokensWithPat_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPersonalAccessTokensWithPat_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates the access an organization member has to organization resources via a fine-grained personal access token. Limited to revoking the token's existing access. Limited to revoking a token's existing access. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#update-the-access-a-fine-grained-personal-access-token-has-to-organization-resources
func (m *ItemPersonalAccessTokensWithPat_ItemRequestBuilder) Post(ctx context.Context, body ItemPersonalAccessTokensItemWithPat_PostRequestBodyable, requestConfiguration *ItemPersonalAccessTokensWithPat_ItemRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Repositories the repositories property
func (m *ItemPersonalAccessTokensWithPat_ItemRequestBuilder) Repositories()(*ItemPersonalAccessTokensItemRepositoriesRequestBuilder) {
    return NewItemPersonalAccessTokensItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPostRequestInformation updates the access an organization member has to organization resources via a fine-grained personal access token. Limited to revoking the token's existing access. Limited to revoking a token's existing access. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
func (m *ItemPersonalAccessTokensWithPat_ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPersonalAccessTokensItemWithPat_PostRequestBodyable, requestConfiguration *ItemPersonalAccessTokensWithPat_ItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPersonalAccessTokensWithPat_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemPersonalAccessTokensWithPat_ItemRequestBuilder) {
    return NewItemPersonalAccessTokensWithPat_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
