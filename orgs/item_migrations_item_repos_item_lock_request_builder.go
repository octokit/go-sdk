package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMigrationsItemReposItemLockRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}\repos\{repo_name}\lock
type ItemMigrationsItemReposItemLockRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMigrationsItemReposItemLockRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMigrationsItemReposItemLockRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMigrationsItemReposItemLockRequestBuilderInternal instantiates a new LockRequestBuilder and sets the default values.
func NewItemMigrationsItemReposItemLockRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposItemLockRequestBuilder) {
    m := &ItemMigrationsItemReposItemLockRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock", pathParameters),
    }
    return m
}
// NewItemMigrationsItemReposItemLockRequestBuilder instantiates a new LockRequestBuilder and sets the default values.
func NewItemMigrationsItemReposItemLockRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsItemReposItemLockRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsItemReposItemLockRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete unlocks a repository that was locked for migration. You should unlock each migrated repository and [delete them](https://docs.github.com/rest/repos/repos#delete-a-repository) when the migration is complete and you no longer need the source data.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/orgs#unlock-an-organization-repository
func (m *ItemMigrationsItemReposItemLockRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMigrationsItemReposItemLockRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation unlocks a repository that was locked for migration. You should unlock each migrated repository and [delete them](https://docs.github.com/rest/repos/repos#delete-a-repository) when the migration is complete and you no longer need the source data.
func (m *ItemMigrationsItemReposItemLockRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMigrationsItemReposItemLockRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemMigrationsItemReposItemLockRequestBuilder) WithUrl(rawUrl string)(*ItemMigrationsItemReposItemLockRequestBuilder) {
    return NewItemMigrationsItemReposItemLockRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
