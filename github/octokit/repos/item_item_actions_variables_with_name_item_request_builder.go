package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemActionsVariablesWithNameItemRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\variables\{name}
type ItemItemActionsVariablesWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsVariablesWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsVariablesWithNameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsVariablesWithNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsVariablesWithNameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemActionsVariablesWithNameItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsVariablesWithNameItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemActionsVariablesWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsVariablesWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsVariablesWithNameItemRequestBuilder) {
    m := &ItemItemActionsVariablesWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/variables/{name}", pathParameters),
    }
    return m
}
// NewItemItemActionsVariablesWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewItemItemActionsVariablesWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsVariablesWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsVariablesWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a repository variable using the variable name.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#delete-a-repository-variable
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get gets a specific variable in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#get-a-repository-variable
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateActionsVariableFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ActionsVariableable), nil
}
// Patch updates a repository variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/variables#update-a-repository-variable
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) Patch(ctx context.Context, body ItemItemActionsVariablesItemWithNamePatchRequestBodyable, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation deletes a repository variable using the variable name.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToGetRequestInformation gets a specific variable in a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:read` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a repository variable that you can reference in a GitHub Actions workflow.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions_variables:write` repository permission to use this endpoint.Authenticated users must have collaborator access to a repository to create, update, or read variables.
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemActionsVariablesItemWithNamePatchRequestBodyable, requestConfiguration *ItemItemActionsVariablesWithNameItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsVariablesWithNameItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsVariablesWithNameItemRequestBuilder) {
    return NewItemItemActionsVariablesWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
