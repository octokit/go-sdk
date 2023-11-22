package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\tags\protection\{tag_protection_id}
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/tags/protection/{tag_protection_id}", pathParameters),
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
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation this deletes a tag protection state for a repository.This endpoint is only available to repository administrators.
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemTagsProtectionWithTag_protection_ItemRequestBuilder) {
    return NewItemItemTagsProtectionWithTag_protection_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
