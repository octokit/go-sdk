package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// InteractionLimitsRequestBuilder builds and executes requests for operations under \user\interaction-limits
type InteractionLimitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InteractionLimitsGetResponse composed type wrapper for classes interactionLimitResponse, InteractionLimitsGetResponseMember1
type InteractionLimitsGetResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable
    // Composed type representation for type InteractionLimitsGetResponseMember1
    interactionLimitsGetResponseMember1 InteractionLimitsGetResponseMember1able
}
// NewInteractionLimitsGetResponse instantiates a new interactionLimitsGetResponse and sets the default values.
func NewInteractionLimitsGetResponse()(*InteractionLimitsGetResponse) {
    m := &InteractionLimitsGetResponse{
    }
    return m
}
// CreateInteractionLimitsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInteractionLimitsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewInteractionLimitsGetResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateInteractionLimitResponseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable); ok {
                result.SetInteractionLimitResponse(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(InteractionLimitsGetResponseMember1able); ok {
                result.SetInteractionLimitsGetResponseMember1(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InteractionLimitsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteractionLimitResponse gets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsGetResponse) GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable) {
    return m.interactionLimitResponse
}
// GetInteractionLimitsGetResponseMember1 gets the InteractionLimitsGetResponseMember1 property value. Composed type representation for type InteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) GetInteractionLimitsGetResponseMember1()(InteractionLimitsGetResponseMember1able) {
    return m.interactionLimitsGetResponseMember1
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsGetResponse) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *InteractionLimitsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitsGetResponseMember1())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteractionLimitResponse sets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsGetResponse) SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)() {
    m.interactionLimitResponse = value
}
// SetInteractionLimitsGetResponseMember1 sets the InteractionLimitsGetResponseMember1 property value. Composed type representation for type InteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) SetInteractionLimitsGetResponseMember1(value InteractionLimitsGetResponseMember1able)() {
    m.interactionLimitsGetResponseMember1 = value
}
// InteractionLimitsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InteractionLimitsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsResponse composed type wrapper for classes interactionLimitResponse, InteractionLimitsGetResponseMember1
type InteractionLimitsResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable
    // Composed type representation for type InteractionLimitsGetResponseMember1
    interactionLimitsGetResponseMember1 InteractionLimitsGetResponseMember1able
}
// NewInteractionLimitsResponse instantiates a new interactionLimitsResponse and sets the default values.
func NewInteractionLimitsResponse()(*InteractionLimitsResponse) {
    m := &InteractionLimitsResponse{
    }
    return m
}
// CreateInteractionLimitsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInteractionLimitsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewInteractionLimitsResponse()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateInteractionLimitResponseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable); ok {
                result.SetInteractionLimitResponse(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(InteractionLimitsGetResponseMember1able); ok {
                result.SetInteractionLimitsGetResponseMember1(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InteractionLimitsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetInteractionLimitResponse gets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsResponse) GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable) {
    return m.interactionLimitResponse
}
// GetInteractionLimitsGetResponseMember1 gets the InteractionLimitsGetResponseMember1 property value. Composed type representation for type InteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) GetInteractionLimitsGetResponseMember1()(InteractionLimitsGetResponseMember1able) {
    return m.interactionLimitsGetResponseMember1
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsResponse) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *InteractionLimitsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitsGetResponseMember1())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteractionLimitResponse sets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsResponse) SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)() {
    m.interactionLimitResponse = value
}
// SetInteractionLimitsGetResponseMember1 sets the InteractionLimitsGetResponseMember1 property value. Composed type representation for type InteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) SetInteractionLimitsGetResponseMember1(value InteractionLimitsGetResponseMember1able)() {
    m.interactionLimitsGetResponseMember1 = value
}
// InteractionLimitsGetResponseable 
type InteractionLimitsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)
    GetInteractionLimitsGetResponseMember1()(InteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)()
    SetInteractionLimitsGetResponseMember1(value InteractionLimitsGetResponseMember1able)()
}
// InteractionLimitsResponseable 
type InteractionLimitsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)
    GetInteractionLimitsGetResponseMember1()(InteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)()
    SetInteractionLimitsGetResponseMember1(value InteractionLimitsGetResponseMember1able)()
}
// NewInteractionLimitsRequestBuilderInternal instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewInteractionLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InteractionLimitsRequestBuilder) {
    m := &InteractionLimitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/interaction-limits", pathParameters),
    }
    return m
}
// NewInteractionLimitsRequestBuilder instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewInteractionLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InteractionLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInteractionLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes any interaction restrictions from your public repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/user#remove-interaction-restrictions-from-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Delete(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get shows which type of GitHub user can interact with your public repositories and when the restriction expires.
// Deprecated: This method is obsolete. Use GetAsInteractionLimitsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/user#get-interaction-restrictions-for-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Get(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInteractionLimitsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InteractionLimitsResponseable), nil
}
// GetAsInteractionLimitsGetResponse shows which type of GitHub user can interact with your public repositories and when the restriction expires.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/user#get-interaction-restrictions-for-your-public-repositories
func (m *InteractionLimitsRequestBuilder) GetAsInteractionLimitsGetResponse(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInteractionLimitsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InteractionLimitsGetResponseable), nil
}
// Put temporarily restricts which type of GitHub user can interact with your public repositories. Setting the interaction limit at the user level will overwrite any interaction limits that are set for individual repositories owned by the user.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/user#set-interaction-restrictions-for-your-public-repositories
func (m *InteractionLimitsRequestBuilder) Put(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitable, requestConfiguration *InteractionLimitsRequestBuilderPutRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateInteractionLimitResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable), nil
}
// ToDeleteRequestInformation removes any interaction restrictions from your public repositories.
func (m *InteractionLimitsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation shows which type of GitHub user can interact with your public repositories and when the restriction expires.
func (m *InteractionLimitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InteractionLimitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation temporarily restricts which type of GitHub user can interact with your public repositories. Setting the interaction limit at the user level will overwrite any interaction limits that are set for individual repositories owned by the user.
func (m *InteractionLimitsRequestBuilder) ToPutRequestInformation(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitable, requestConfiguration *InteractionLimitsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *InteractionLimitsRequestBuilder) WithUrl(rawUrl string)(*InteractionLimitsRequestBuilder) {
    return NewInteractionLimitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
