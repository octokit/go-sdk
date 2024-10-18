package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnerGroupsItemRunnersRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runner-groups\{runner_group_id}\runners
type ItemActionsRunnerGroupsItemRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnerGroupsItemRunnersRequestBuilderGetQueryParameters lists self-hosted runners that are in a specific organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemActionsRunnerGroupsItemRunnersRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByRunner_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.runnerGroups.item.runners.item collection
// returns a *ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) ByRunner_id(runner_id int32)(*ItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["runner_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(runner_id), 10)
    return NewItemActionsRunnerGroupsItemRunnersWithRunner_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRunnerGroupsItemRunnersRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsItemRunnersRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRunnersRequestBuilder) {
    m := &ItemActionsRunnerGroupsItemRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runner-groups/{runner_group_id}/runners{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsItemRunnersRequestBuilder instantiates a new ItemActionsRunnerGroupsItemRunnersRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsItemRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists self-hosted runners that are in a specific organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemActionsRunnerGroupsItemRunnersGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#list-self-hosted-runners-in-a-group-for-an-organization
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemRunnersRequestBuilderGetQueryParameters])(ItemActionsRunnerGroupsItemRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnerGroupsItemRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnerGroupsItemRunnersGetResponseable), nil
}
// Put replaces the list of self-hosted runners that are part of an organization runner group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#set-self-hosted-runners-in-a-group-for-an-organization
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) Put(ctx context.Context, body ItemActionsRunnerGroupsItemRunnersPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation lists self-hosted runners that are in a specific organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemRunnersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation replaces the list of self-hosted runners that are part of an organization runner group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsRunnerGroupsItemRunnersPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsItemRunnersRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsItemRunnersRequestBuilder) {
    return NewItemActionsRunnerGroupsItemRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
