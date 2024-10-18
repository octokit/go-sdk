package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runner-groups\{runner_group_id}\runners\{runner_id}
type ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) {
    m := &ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder instantiates a new ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes a self-hosted runner from a group configured in an organization. The runner is then returned to the default group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#remove-a-self-hosted-runner-from-a-group-for-an-organization
func (m *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// Put adds a self-hosted runner to a runner group configured in an organization.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#add-a-self-hosted-runner-to-a-group-for-an-organization
func (m *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) Put(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation removes a self-hosted runner from a group configured in an organization. The runner is then returned to the default group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToPutRequestInformation adds a self-hosted runner to a runner group configured in an organization.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) {
    return NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
