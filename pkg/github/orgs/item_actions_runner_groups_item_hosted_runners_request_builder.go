package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder builds and executes requests for operations under \orgs\{org}\actions\runner-groups\{runner_group_id}\hosted-runners
type ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActionsRunnerGroupsItemHostedRunnersRequestBuilderGetQueryParameters lists the GitHub-hosted runners in an organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemActionsRunnerGroupsItemHostedRunnersRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilderInternal instantiates a new ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) {
    m := &ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/actions/runner-groups/{runner_group_id}/hosted-runners{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilder instantiates a new ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder and sets the default values.
func NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the GitHub-hosted runners in an organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a ItemActionsRunnerGroupsItemHostedRunnersGetResponseable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/self-hosted-runner-groups#list-github-hosted-runners-in-a-group-for-an-organization
func (m *ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemHostedRunnersRequestBuilderGetQueryParameters])(ItemActionsRunnerGroupsItemHostedRunnersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemActionsRunnerGroupsItemHostedRunnersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemActionsRunnerGroupsItemHostedRunnersGetResponseable), nil
}
// ToGetRequestInformation lists the GitHub-hosted runners in an organization group.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemActionsRunnerGroupsItemHostedRunnersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder when successful
func (m *ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) WithUrl(rawUrl string)(*ItemActionsRunnerGroupsItemHostedRunnersRequestBuilder) {
    return NewItemActionsRunnerGroupsItemHostedRunnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
