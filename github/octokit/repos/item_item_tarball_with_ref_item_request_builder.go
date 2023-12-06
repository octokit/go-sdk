package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemTarballWithRefItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\tarball\{ref}
type ItemItemTarballWithRefItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemTarballWithRefItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemTarballWithRefItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemTarballWithRefItemRequestBuilderInternal instantiates a new WithRefItemRequestBuilder and sets the default values.
func NewItemItemTarballWithRefItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballWithRefItemRequestBuilder) {
    m := &ItemItemTarballWithRefItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/tarball/{ref}", pathParameters),
    }
    return m
}
// NewItemItemTarballWithRefItemRequestBuilder instantiates a new WithRefItemRequestBuilder and sets the default values.
func NewItemItemTarballWithRefItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemTarballWithRefItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemTarballWithRefItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a redirect URL to download a tar archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#download-a-repository-archive-tar
func (m *ItemItemTarballWithRefItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemTarballWithRefItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation gets a redirect URL to download a tar archive for a repository. If you omit `:ref`, the repository’s default branch (usually`main`) will be used. Please make sure your HTTP framework is configured to follow redirects or you will need to usethe `Location` header to make a second `GET` request.**Note**: For private repositories, these links are temporary and expire after five minutes.
func (m *ItemItemTarballWithRefItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemTarballWithRefItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemTarballWithRefItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemTarballWithRefItemRequestBuilder) {
    return NewItemItemTarballWithRefItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
