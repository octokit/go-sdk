package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemStatusesWithShaItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\statuses\{sha}
type ItemItemStatusesWithShaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStatusesWithShaItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemStatusesWithShaItemRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemStatusesWithShaItemRequestBuilderInternal instantiates a new WithShaItemRequestBuilder and sets the default values.
func NewItemItemStatusesWithShaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatusesWithShaItemRequestBuilder) {
    m := &ItemItemStatusesWithShaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/statuses/{sha}", pathParameters),
    }
    return m
}
// NewItemItemStatusesWithShaItemRequestBuilder instantiates a new WithShaItemRequestBuilder and sets the default values.
func NewItemItemStatusesWithShaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStatusesWithShaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStatusesWithShaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Post users with push access in a repository can create commit statuses for a given SHA.Note: there is a limit of 1000 statuses per `sha` and `context` within a repository. Attempts to create more than 1000 statuses will result in a validation error.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/commits/statuses#create-a-commit-status
func (m *ItemItemStatusesWithShaItemRequestBuilder) Post(ctx context.Context, body ItemItemStatusesItemWithShaPostRequestBodyable, requestConfiguration *ItemItemStatusesWithShaItemRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Statusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateStatusFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Statusable), nil
}
// ToPostRequestInformation users with push access in a repository can create commit statuses for a given SHA.Note: there is a limit of 1000 statuses per `sha` and `context` within a repository. Attempts to create more than 1000 statuses will result in a validation error.
func (m *ItemItemStatusesWithShaItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemStatusesItemWithShaPostRequestBodyable, requestConfiguration *ItemItemStatusesWithShaItemRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemStatusesWithShaItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemStatusesWithShaItemRequestBuilder) {
    return NewItemItemStatusesWithShaItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
