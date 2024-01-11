package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// CodespacesRequestBuilder builds and executes requests for operations under \user\codespaces
type CodespacesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesPostRequestBody composed type wrapper for classes CodespacesPostRequestBodyMember1, CodespacesPostRequestBodyMember2
type CodespacesPostRequestBody struct {
    // Composed type representation for type CodespacesPostRequestBodyMember1
    codespacesPostRequestBodyCodespacesPostRequestBodyMember1 CodespacesPostRequestBodyMember1able
    // Composed type representation for type CodespacesPostRequestBodyMember2
    codespacesPostRequestBodyCodespacesPostRequestBodyMember2 CodespacesPostRequestBodyMember2able
    // Composed type representation for type CodespacesPostRequestBodyMember1
    codespacesPostRequestBodyMember1 CodespacesPostRequestBodyMember1able
    // Composed type representation for type CodespacesPostRequestBodyMember2
    codespacesPostRequestBodyMember2 CodespacesPostRequestBodyMember2able
}
// NewCodespacesPostRequestBody instantiates a new codespacesPostRequestBody and sets the default values.
func NewCodespacesPostRequestBody()(*CodespacesPostRequestBody) {
    m := &CodespacesPostRequestBody{
    }
    return m
}
// CreateCodespacesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCodespacesPostRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1 gets the CodespacesPostRequestBodyMember1 property value. Composed type representation for type CodespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able) {
    return m.codespacesPostRequestBodyCodespacesPostRequestBodyMember1
}
// GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2 gets the CodespacesPostRequestBodyMember2 property value. Composed type representation for type CodespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able) {
    return m.codespacesPostRequestBodyCodespacesPostRequestBodyMember2
}
// GetCodespacesPostRequestBodyMember1 gets the CodespacesPostRequestBodyMember1 property value. Composed type representation for type CodespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able) {
    return m.codespacesPostRequestBodyMember1
}
// GetCodespacesPostRequestBodyMember2 gets the CodespacesPostRequestBodyMember2 property value. Composed type representation for type CodespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) GetCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able) {
    return m.codespacesPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *CodespacesPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *CodespacesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetCodespacesPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCodespacesPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCodespacesPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1 sets the CodespacesPostRequestBodyMember1 property value. Composed type representation for type CodespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)() {
    m.codespacesPostRequestBodyCodespacesPostRequestBodyMember1 = value
}
// SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2 sets the CodespacesPostRequestBodyMember2 property value. Composed type representation for type CodespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)() {
    m.codespacesPostRequestBodyCodespacesPostRequestBodyMember2 = value
}
// SetCodespacesPostRequestBodyMember1 sets the CodespacesPostRequestBodyMember1 property value. Composed type representation for type CodespacesPostRequestBodyMember1
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)() {
    m.codespacesPostRequestBodyMember1 = value
}
// SetCodespacesPostRequestBodyMember2 sets the CodespacesPostRequestBodyMember2 property value. Composed type representation for type CodespacesPostRequestBodyMember2
func (m *CodespacesPostRequestBody) SetCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)() {
    m.codespacesPostRequestBodyMember2 = value
}
// CodespacesRequestBuilderGetQueryParameters lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
type CodespacesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // ID of the Repository to filter on
    Repository_id *int32 `uriparametername:"repository_id"`
}
// CodespacesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CodespacesRequestBuilderGetQueryParameters
}
// CodespacesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CodespacesPostRequestBodyable 
type CodespacesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able)
    GetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able)
    GetCodespacesPostRequestBodyMember1()(CodespacesPostRequestBodyMember1able)
    GetCodespacesPostRequestBodyMember2()(CodespacesPostRequestBodyMember2able)
    SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)()
    SetCodespacesPostRequestBodyCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)()
    SetCodespacesPostRequestBodyMember1(value CodespacesPostRequestBodyMember1able)()
    SetCodespacesPostRequestBodyMember2(value CodespacesPostRequestBodyMember2able)()
}
// ByCodespace_name gets an item from the github.com/octokit/go-sdk/pkg/github/.user.codespaces.item collection
func (m *CodespacesRequestBuilder) ByCodespace_name(codespace_name string)(*CodespacesWithCodespace_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if codespace_name != "" {
        urlTplParams["codespace_name"] = codespace_name
    }
    return NewCodespacesWithCodespace_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCodespacesRequestBuilderInternal instantiates a new CodespacesRequestBuilder and sets the default values.
func NewCodespacesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesRequestBuilder) {
    m := &CodespacesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces{?per_page*,page*,repository_id*}", pathParameters),
    }
    return m
}
// NewCodespacesRequestBuilder instantiates a new CodespacesRequestBuilder and sets the default values.
func NewCodespacesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
// Deprecated: This method is obsolete. Use GetAsCodespacesGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-codespaces-for-the-authenticated-user
func (m *CodespacesRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesRequestBuilderGetRequestConfiguration)(CodespacesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesResponseable), nil
}
// GetAsCodespacesGetResponse lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#list-codespaces-for-the-authenticated-user
func (m *CodespacesRequestBuilder) GetAsCodespacesGetResponse(ctx context.Context, requestConfiguration *CodespacesRequestBuilderGetRequestConfiguration)(CodespacesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCodespacesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CodespacesGetResponseable), nil
}
// Post creates a new codespace, owned by the authenticated user.This endpoint requires either a `repository_id` OR a `pull_request` but not both.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#create-a-codespace-for-the-authenticated-user
func (m *CodespacesRequestBuilder) Post(ctx context.Context, body CodespacesPostRequestBodyable, requestConfiguration *CodespacesRequestBuilderPostRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespace503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Codespaceable), nil
}
// Secrets the secrets property
func (m *CodespacesRequestBuilder) Secrets()(*CodespacesSecretsRequestBuilder) {
    return NewCodespacesSecretsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation lists the authenticated user's codespaces.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have read access to the `codespaces` repository permission.
func (m *CodespacesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a new codespace, owned by the authenticated user.This endpoint requires either a `repository_id` OR a `pull_request` but not both.You must authenticate using an access token with the `codespace` scope to use this endpoint.To use this endpoint with GitHub Apps:- The app must be authenticated on behalf of the user. For more information, see "[Authenticating with a GitHub App on behalf of a user](https://docs.github.com/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-with-a-github-app-on-behalf-of-a-user)."- The app must have write access to the `codespaces` repository permission.
func (m *CodespacesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CodespacesPostRequestBodyable, requestConfiguration *CodespacesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesRequestBuilder) WithUrl(rawUrl string)(*CodespacesRequestBuilder) {
    return NewCodespacesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
