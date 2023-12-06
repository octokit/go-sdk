package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\tags\protection\{tag_protection_id}
type ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTagsProtectionWithTag_protection_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTagsProtectionWithTag_protection_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal instantiates a new WithTag_protection_ItemRequestBuilder and sets the default values.
func NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    m := &ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/tags/protection/{tag_protection_id}", pathParameters),
    }
    return m
}
// NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder instantiates a new WithTag_protection_ItemRequestBuilder and sets the default values.
func NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete this deletes a tag protection state for a repository.This endpoint is only available to repository administrators.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/tags#delete-a-tag-protection-state-for-a-repository
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation this deletes a tag protection state for a repository.This endpoint is only available to repository administrators.
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    return NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
