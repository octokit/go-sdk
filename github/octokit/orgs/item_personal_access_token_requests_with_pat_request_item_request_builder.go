package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\personal-access-token-requests\{pat_request_id}
type ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal instantiates a new WithPat_request_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    m := &ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/personal-access-token-requests/{pat_request_id}", pathParameters),
    }
    return m
}
// NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder instantiates a new WithPat_request_ItemRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approves or denies a pending request to access organization resources via a fine-grained personal access token. Only GitHub Apps can call this API,using the `organization_personal_access_token_requests: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#review-a-request-to-access-organization-resources-with-a-fine-grained-personal-access-token
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) Post(ctx context.Context, body ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyable, requestConfiguration *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
        "500": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Repositories the repositories property
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) Repositories()(*ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToPostRequestInformation approves or denies a pending request to access organization resources via a fine-grained personal access token. Only GitHub Apps can call this API,using the `organization_personal_access_token_requests: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPersonalAccessTokenRequestsItemWithPat_request_PostRequestBodyable, requestConfiguration *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsWithPat_request_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
