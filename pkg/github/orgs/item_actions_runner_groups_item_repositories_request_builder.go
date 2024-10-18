package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnerGroupsItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runner-groups\{runner_group_id}\repositories
type ItemActionsRunnerGroupsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnerGroupsItemRepositoriesRequestBuilderGetQueryParameters lists the repositories with access to a self-hosted runner group configured in an organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemActionsRunnerGroupsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByRepository_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.runnerGroups.item.repositories.item collection
// returns a *ItemActionsRunnerGroupsItemRepositoriesWithRepository_ItemRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) ByRepository_id(repository_id int32)(*ItemActionsRunnerGroupsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["repository_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(repository_id), 10)
    return NewItemActionsRunnerGroupsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRunnerGroupsItemRepositoriesRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsItemRepositoriesRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) {
    m := &ItemActionsRunnerGroupsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runner-groups/{runner_group_id}/repositories{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsItemRepositoriesRequestBuilder instantiates a new ItemActionsRunnerGroupsItemRepositoriesRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the repositories with access to a self-hosted runner group configured in an organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemActionsRunnerGroupsItemRepositoriesGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#list-repository-access-to-a-self-hosted-runner-group-in-an-organization
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemRepositoriesRequestBuilderGetQueryParameters])(ItemActionsRunnerGroupsItemRepositoriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnerGroupsItemRepositoriesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnerGroupsItemRepositoriesGetResponseable), nil
}
// Put replaces the list of repositories that have access to a self-hosted runner group configured in an organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#set-repository-access-for-a-self-hosted-runner-group-in-an-organization
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) Put(ctx context.Context, body ItemActionsRunnerGroupsItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToGetRequestInformation lists the repositories with access to a self-hosted runner group configured in an organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemRepositoriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation replaces the list of repositories that have access to a self-hosted runner group configured in an organization.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemActionsRunnerGroupsItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsItemRepositoriesRequestBuilder) {
    return NewItemActionsRunnerGroupsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
