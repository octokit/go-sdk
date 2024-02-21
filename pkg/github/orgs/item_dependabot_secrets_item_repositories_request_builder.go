package orgs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemDependabotSecretsItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\dependabot\secrets\{secret_name}\repositories
type ItemDependabotSecretsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDependabotSecretsItemRepositoriesRequestBuilderGetQueryParameters lists all repositories that have been selected when the `visibility`for repository access to a secret is set to `selected`.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
type ItemDependabotSecretsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// ByRepository_id gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.dependabot.secrets.item.repositories.item collection
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) ByRepository_id(repository_id int32)(*ItemDependabotSecretsItemRepositoriesWithRepository_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["repository_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(repository_id), 10)
    return NewItemDependabotSecretsItemRepositoriesWithRepository_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDependabotSecretsItemRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemDependabotSecretsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsItemRepositoriesRequestBuilder) {
    m := &ItemDependabotSecretsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/dependabot/secrets/{secret_name}/repositories{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemDependabotSecretsItemRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemDependabotSecretsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDependabotSecretsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDependabotSecretsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all repositories that have been selected when the `visibility`for repository access to a secret is set to `selected`.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#list-selected-repositories-for-an-organization-secret
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotSecretsItemRepositoriesRequestBuilderGetQueryParameters])(ItemDependabotSecretsItemRepositoriesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDependabotSecretsItemRepositoriesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDependabotSecretsItemRepositoriesGetResponseable), nil
}
// Put replaces all repositories for an organization secret when the `visibility`for repository access is set to `selected`. The visibility is set when you [Createor update an organization secret](https://docs.github.com/rest/dependabot/secrets#create-or-update-an-organization-secret).OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/dependabot/secrets#set-selected-repositories-for-an-organization-secret
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) Put(ctx context.Context, body ItemDependabotSecretsItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToGetRequestInformation lists all repositories that have been selected when the `visibility`for repository access to a secret is set to `selected`.OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemDependabotSecretsItemRepositoriesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation replaces all repositories for an organization secret when the `visibility`for repository access is set to `selected`. The visibility is set when you [Createor update an organization secret](https://docs.github.com/rest/dependabot/secrets#create-or-update-an-organization-secret).OAuth app tokens and personal access tokens (classic) need the `admin:org` scope to use this endpoint.
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemDependabotSecretsItemRepositoriesPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemDependabotSecretsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*ItemDependabotSecretsItemRepositoriesRequestBuilder) {
    return NewItemDependabotSecretsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
