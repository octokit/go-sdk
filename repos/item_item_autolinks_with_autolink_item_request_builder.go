package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemAutolinksWithAutolink_ItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\autolinks\{autolink_id}
type ItemItemAutolinksWithAutolink_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemAutolinksWithAutolink_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemAutolinksWithAutolink_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemAutolinksWithAutolink_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemAutolinksWithAutolink_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemAutolinksWithAutolink_ItemRequestBuilderInternal instantiates a new WithAutolink_ItemRequestBuilder and sets the default values.
func NewItemItemAutolinksWithAutolink_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAutolinksWithAutolink_ItemRequestBuilder) {
    m := &ItemItemAutolinksWithAutolink_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/autolinks/{autolink_id}", pathParameters),
    }
    return m
}
// NewItemItemAutolinksWithAutolink_ItemRequestBuilder instantiates a new WithAutolink_ItemRequestBuilder and sets the default values.
func NewItemItemAutolinksWithAutolink_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemAutolinksWithAutolink_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemAutolinksWithAutolink_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete this deletes a single autolink reference by ID that was configured for the given repository.Information about autolinks are only available to repository administrators.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/autolinks#delete-an-autolink-reference-from-a-repository
func (m *ItemItemAutolinksWithAutolink_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemAutolinksWithAutolink_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get this returns a single autolink reference by ID that was configured for the given repository.Information about autolinks are only available to repository administrators.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/autolinks#get-an-autolink-reference-of-a-repository
func (m *ItemItemAutolinksWithAutolink_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemAutolinksWithAutolink_ItemRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Autolinkable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateAutolinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Autolinkable), nil
}
// ToDeleteRequestInformation this deletes a single autolink reference by ID that was configured for the given repository.Information about autolinks are only available to repository administrators.
func (m *ItemItemAutolinksWithAutolink_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemAutolinksWithAutolink_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation this returns a single autolink reference by ID that was configured for the given repository.Information about autolinks are only available to repository administrators.
func (m *ItemItemAutolinksWithAutolink_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemAutolinksWithAutolink_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemAutolinksWithAutolink_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemAutolinksWithAutolink_ItemRequestBuilder) {
    return NewItemItemAutolinksWithAutolink_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
