package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemDockerConflictsRequestBuilder builds and executes requests for operations under \users\{username}\docker\conflicts
type ItemDockerConflictsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDockerConflictsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDockerConflictsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDockerConflictsRequestBuilderInternal instantiates a new ConflictsRequestBuilder and sets the default values.
func NewItemDockerConflictsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDockerConflictsRequestBuilder) {
    m := &ItemDockerConflictsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/docker/conflicts", pathParameters),
    }
    return m
}
// NewItemDockerConflictsRequestBuilder instantiates a new ConflictsRequestBuilder and sets the default values.
func NewItemDockerConflictsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDockerConflictsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDockerConflictsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all packages that are in a specific user's namespace, that the requesting user has access to, and that encountered a conflict during Docker migration.To use this endpoint, you must authenticate using an access token with the `read:packages` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/packages/packages#get-list-of-conflicting-packages-during-docker-migration-for-user
func (m *ItemDockerConflictsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDockerConflictsRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageEscapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreatePackageEscapedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageEscapedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PackageEscapedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all packages that are in a specific user's namespace, that the requesting user has access to, and that encountered a conflict during Docker migration.To use this endpoint, you must authenticate using an access token with the `read:packages` scope.
func (m *ItemDockerConflictsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDockerConflictsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemDockerConflictsRequestBuilder) WithUrl(rawUrl string)(*ItemDockerConflictsRequestBuilder) {
    return NewItemDockerConflictsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
