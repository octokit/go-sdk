package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemInteractionLimitsRequestBuilder builds and executes requests for operations under \orgs\{org}\interaction-limits
type ItemInteractionLimitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InteractionLimitsGetResponse composed type wrapper for classes interactionLimitResponse, ItemInteractionLimitsGetResponseMember1
type InteractionLimitsGetResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable
    // Composed type representation for type ItemInteractionLimitsGetResponseMember1
    itemInteractionLimitsGetResponseMember1 ItemInteractionLimitsGetResponseMember1able
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
        } else if val, err := parseNode.GetObjectValue(CreateItemInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemInteractionLimitsGetResponseMember1able); ok {
                result.SetItemInteractionLimitsGetResponseMember1(cast)
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
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemInteractionLimitsGetResponseMember1 gets the ItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) GetItemInteractionLimitsGetResponseMember1()(ItemInteractionLimitsGetResponseMember1able) {
    return m.itemInteractionLimitsGetResponseMember1
}
// Serialize serializes information the current object
func (m *InteractionLimitsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetItemInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemInteractionLimitsGetResponseMember1())
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
// SetItemInteractionLimitsGetResponseMember1 sets the ItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) SetItemInteractionLimitsGetResponseMember1(value ItemInteractionLimitsGetResponseMember1able)() {
    m.itemInteractionLimitsGetResponseMember1 = value
}
// InteractionLimitsResponse composed type wrapper for classes interactionLimitResponse, ItemInteractionLimitsGetResponseMember1
type InteractionLimitsResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable
    // Composed type representation for type ItemInteractionLimitsGetResponseMember1
    itemInteractionLimitsGetResponseMember1 ItemInteractionLimitsGetResponseMember1able
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
        } else if val, err := parseNode.GetObjectValue(CreateItemInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemInteractionLimitsGetResponseMember1able); ok {
                result.SetItemInteractionLimitsGetResponseMember1(cast)
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
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemInteractionLimitsGetResponseMember1 gets the ItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) GetItemInteractionLimitsGetResponseMember1()(ItemInteractionLimitsGetResponseMember1able) {
    return m.itemInteractionLimitsGetResponseMember1
}
// Serialize serializes information the current object
func (m *InteractionLimitsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetItemInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemInteractionLimitsGetResponseMember1())
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
// SetItemInteractionLimitsGetResponseMember1 sets the ItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) SetItemInteractionLimitsGetResponseMember1(value ItemInteractionLimitsGetResponseMember1able)() {
    m.itemInteractionLimitsGetResponseMember1 = value
}
// ItemInteractionLimitsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInteractionLimitsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInteractionLimitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInteractionLimitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInteractionLimitsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInteractionLimitsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsGetResponseable 
type InteractionLimitsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)
    GetItemInteractionLimitsGetResponseMember1()(ItemInteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)()
    SetItemInteractionLimitsGetResponseMember1(value ItemInteractionLimitsGetResponseMember1able)()
}
// InteractionLimitsResponseable 
type InteractionLimitsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)
    GetItemInteractionLimitsGetResponseMember1()(ItemInteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable)()
    SetItemInteractionLimitsGetResponseMember1(value ItemInteractionLimitsGetResponseMember1able)()
}
// NewItemInteractionLimitsRequestBuilderInternal instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewItemInteractionLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInteractionLimitsRequestBuilder) {
    m := &ItemInteractionLimitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/interaction-limits", pathParameters),
    }
    return m
}
// NewItemInteractionLimitsRequestBuilder instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewItemInteractionLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInteractionLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInteractionLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes all interaction restrictions from public repositories in the given organization. You must be an organization owner to remove restrictions.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/orgs#remove-interaction-restrictions-for-an-organization
func (m *ItemInteractionLimitsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInteractionLimitsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get shows which type of GitHub user can interact with this organization and when the restriction expires. If there is no restrictions, you will see an empty response.
// Deprecated: This method is obsolete. Use GetAsInteractionLimitsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/orgs#get-interaction-restrictions-for-an-organization
func (m *ItemInteractionLimitsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsResponseable, error) {
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
// GetAsInteractionLimitsGetResponse shows which type of GitHub user can interact with this organization and when the restriction expires. If there is no restrictions, you will see an empty response.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/orgs#get-interaction-restrictions-for-an-organization
func (m *ItemInteractionLimitsRequestBuilder) GetAsInteractionLimitsGetResponse(ctx context.Context, requestConfiguration *ItemInteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsGetResponseable, error) {
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
// Put temporarily restricts interactions to a certain type of GitHub user in any public repository in the given organization. You must be an organization owner to set these restrictions. Setting the interaction limit at the organization level will overwrite any interaction limits that are set for individual repositories owned by the organization.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/orgs#set-interaction-restrictions-for-an-organization
func (m *ItemInteractionLimitsRequestBuilder) Put(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitable, requestConfiguration *ItemInteractionLimitsRequestBuilderPutRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitResponseable, error) {
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
// ToDeleteRequestInformation removes all interaction restrictions from public repositories in the given organization. You must be an organization owner to remove restrictions.
func (m *ItemInteractionLimitsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInteractionLimitsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation shows which type of GitHub user can interact with this organization and when the restriction expires. If there is no restrictions, you will see an empty response.
func (m *ItemInteractionLimitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInteractionLimitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation temporarily restricts interactions to a certain type of GitHub user in any public repository in the given organization. You must be an organization owner to set these restrictions. Setting the interaction limit at the organization level will overwrite any interaction limits that are set for individual repositories owned by the organization.
func (m *ItemInteractionLimitsRequestBuilder) ToPutRequestInformation(ctx context.Context, body i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.InteractionLimitable, requestConfiguration *ItemInteractionLimitsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemInteractionLimitsRequestBuilder) WithUrl(rawUrl string)(*ItemInteractionLimitsRequestBuilder) {
    return NewItemInteractionLimitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
