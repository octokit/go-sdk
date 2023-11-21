package gitignore

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TemplatesWithNameItemRequestBuilder builds and executes requests for operations under \gitignore\templates\{name}
type TemplatesWithNameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesWithNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesWithNameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesWithNameItemRequestBuilderInternal instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewTemplatesWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesWithNameItemRequestBuilder) {
    m := &TemplatesWithNameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gitignore/templates/{name}", pathParameters),
    }
    return m
}
// NewTemplatesWithNameItemRequestBuilder instantiates a new WithNameItemRequestBuilder and sets the default values.
func NewTemplatesWithNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesWithNameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the API also allows fetching the source of a single template.Use the raw [media type](https://docs.github.com/rest/overview/media-types/) to get the raw contents.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/gitignore/gitignore#get-a-gitignore-template
func (m *TemplatesWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesWithNameItemRequestBuilderGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GitignoreTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateGitignoreTemplateFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.GitignoreTemplateable), nil
}
// ToGetRequestInformation the API also allows fetching the source of a single template.Use the raw [media type](https://docs.github.com/rest/overview/media-types/) to get the raw contents.
func (m *TemplatesWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesWithNameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *TemplatesWithNameItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesWithNameItemRequestBuilder) {
    return NewTemplatesWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
