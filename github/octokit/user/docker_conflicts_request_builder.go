package user

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DockerConflictsRequestBuilder builds and executes requests for operations under \user\docker\conflicts
type DockerConflictsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DockerConflictsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DockerConflictsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDockerConflictsRequestBuilderInternal instantiates a new ConflictsRequestBuilder and sets the default values.
func NewDockerConflictsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DockerConflictsRequestBuilder) {
    m := &DockerConflictsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/docker/conflicts", pathParameters),
    }
    return m
}
// NewDockerConflictsRequestBuilder instantiates a new ConflictsRequestBuilder and sets the default values.
func NewDockerConflictsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DockerConflictsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDockerConflictsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists all packages that are owned by the authenticated user within the user's namespace, and that encountered a conflict during a Docker migration.To use this endpoint, you must authenticate using an access token with the `read:packages` scope.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/packages/packages#get-list-of-conflicting-packages-during-docker-migration-for-authenticated-user
func (m *DockerConflictsRequestBuilder) Get(ctx context.Context, requestConfiguration *DockerConflictsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PackageEscapedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreatePackageEscapedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PackageEscapedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.PackageEscapedable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists all packages that are owned by the authenticated user within the user's namespace, and that encountered a conflict during a Docker migration.To use this endpoint, you must authenticate using an access token with the `read:packages` scope.
func (m *DockerConflictsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DockerConflictsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DockerConflictsRequestBuilder) WithUrl(rawUrl string)(*DockerConflictsRequestBuilder) {
    return NewDockerConflictsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
