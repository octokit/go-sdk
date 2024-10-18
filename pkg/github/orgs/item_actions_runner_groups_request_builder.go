package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemActionsRunnerGroupsRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runner-groups
type ItemActionsRunnerGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnerGroupsRequestBuilderGetQueryParameters lists all self-hosted runner groups configured in an organization and inherited from an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemActionsRunnerGroupsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // Only return runner groups that are allowed to be used by this repository.
    Visible_to_repository *string `uriparametername:"visible_to_repository"`
}
// ByRunner_group_id gets an item from the github.com/octokit/go-sdk/pkg/github.orgs.item.actions.runnerGroups.item collection
// returns a *ItemActionsRunnerGroupsWithRunner_group_ItemRequestBuilder when successful
func (m *ItemActionsRunnerGroupsRequestBuilder) ByRunner_group_id(runner_group_id int32)(*ItemActionsRunnerGroupsWithRunner_group_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["runner_group_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(runner_group_id), 10)
    return NewItemActionsRunnerGroupsWithRunner_group_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActionsRunnerGroupsRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsRequestBuilder) {
    m := &ItemActionsRunnerGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runner-groups{?page*,per_page*,visible_to_repository*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsRequestBuilder instantiates a new ItemActionsRunnerGroupsRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all self-hosted runner groups configured in an organization and inherited from an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemActionsRunnerGroupsGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#list-self-hosted-runner-groups-for-an-organization
func (m *ItemActionsRunnerGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsRequestBuilderGetQueryParameters])(ItemActionsRunnerGroupsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnerGroupsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnerGroupsGetResponseable), nil
}
// Post creates a new self-hosted runner group for an organization.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a RunnerGroupsOrgable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#create-a-self-hosted-runner-group-for-an-organization
func (m *ItemActionsRunnerGroupsRequestBuilder) Post(ctx context.Context, body ItemActionsRunnerGroupsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerGroupsOrgable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRunnerGroupsOrgFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.RunnerGroupsOrgable), nil
}
// ToGetRequestInformation lists all self-hosted runner groups configured in an organization and inherited from an enterprise.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a new self-hosted runner group for an organization.OAuth tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemActionsRunnerGroupsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsRequestBuilder when successful
func (m *ItemActionsRunnerGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsRequestBuilder) {
    return NewItemActionsRunnerGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
