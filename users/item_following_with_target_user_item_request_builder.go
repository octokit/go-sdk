package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemFollowingWithTarget_userItemRequestBuilder builds and executes requests for operations under \users\{username}\following\{target_user}
type ItemFollowingWithTarget_userItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemFollowingWithTarget_userItemRequestBuilderInternal instantiates a new WithTarget_userItemRequestBuilder and sets the default values.
func NewItemFollowingWithTarget_userItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    m := &ItemFollowingWithTarget_userItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/following/{target_user}", pathParameters),
    }
    return m
}
// NewItemFollowingWithTarget_userItemRequestBuilder instantiates a new WithTarget_userItemRequestBuilder and sets the default values.
func NewItemFollowingWithTarget_userItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFollowingWithTarget_userItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check if a user follows another user
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/followers#check-if-a-user-follows-another-user
func (m *ItemFollowingWithTarget_userItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration)(error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ItemFollowingWithTarget_userItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFollowingWithTarget_userItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemFollowingWithTarget_userItemRequestBuilder) WithUrl(rawUrl string)(*ItemFollowingWithTarget_userItemRequestBuilder) {
    return NewItemFollowingWithTarget_userItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
