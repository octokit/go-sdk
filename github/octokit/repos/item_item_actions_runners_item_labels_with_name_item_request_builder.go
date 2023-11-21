package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runners\{runner_id}\labels\{name}
type ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    m := &ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runners/{runner_id}/labels/{name}", pathParameters),
    }
    return m
}
// NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a custom label from a self-hosted runner configuredin a repository. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// Deprecated: This method is obsolete. Use DeleteAsWithNameDeleteResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#remove-a-custom-label-from-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration)(ItemItemActionsRunnersItemLabelsItemWithNameResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsItemWithNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsItemWithNameResponseable), nil
}
// DeleteAsWithNameDeleteResponse remove a custom label from a self-hosted runner configuredin a repository. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runners#remove-a-custom-label-from-a-self-hosted-runner-for-a-repository
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) DeleteAsWithNameDeleteResponse(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration)(ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemActionsRunnersItemLabelsItemWithNameDeleteResponseable), nil
}
// ToDeleteRequestInformation remove a custom label from a self-hosted runner configuredin a repository. Returns the remaining labels from the runner.This endpoint returns a `404 Not Found` status if the custom label is notpresent on the runner.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `administration` permission for repositories and the `organization_self_hosted_runners` permission for organizations.Authenticated users must have admin access to repositories or organizations, or the `manage_runners:enterprise` scope for enterprises, to use these endpoints.
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder) {
    return NewItemItemActionsRunnersItemLabelsWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
