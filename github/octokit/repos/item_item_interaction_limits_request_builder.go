package repos

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemInteractionLimitsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\interaction-limits
type ItemItemInteractionLimitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InteractionLimitsGetResponse composed type wrapper for classes interactionLimitResponse, ItemItemInteractionLimitsGetResponseMember1
type InteractionLimitsGetResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable
    // Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
    itemItemInteractionLimitsGetResponseMember1 ItemItemInteractionLimitsGetResponseMember1able
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
        if val, err := parseNode.GetObjectValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateInteractionLimitResponseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable); ok {
                result.SetInteractionLimitResponse(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemInteractionLimitsGetResponseMember1able); ok {
                result.SetItemItemInteractionLimitsGetResponseMember1(cast)
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
func (m *InteractionLimitsGetResponse) GetInteractionLimitResponse()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable) {
    return m.interactionLimitResponse
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemItemInteractionLimitsGetResponseMember1 gets the ItemItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) GetItemItemInteractionLimitsGetResponseMember1()(ItemItemInteractionLimitsGetResponseMember1able) {
    return m.itemItemInteractionLimitsGetResponseMember1
}
// Serialize serializes information the current object
func (m *InteractionLimitsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetItemItemInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemInteractionLimitsGetResponseMember1())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteractionLimitResponse sets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsGetResponse) SetInteractionLimitResponse(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)() {
    m.interactionLimitResponse = value
}
// SetItemItemInteractionLimitsGetResponseMember1 sets the ItemItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsGetResponse) SetItemItemInteractionLimitsGetResponseMember1(value ItemItemInteractionLimitsGetResponseMember1able)() {
    m.itemItemInteractionLimitsGetResponseMember1 = value
}
// InteractionLimitsResponse composed type wrapper for classes interactionLimitResponse, ItemItemInteractionLimitsGetResponseMember1
type InteractionLimitsResponse struct {
    // Composed type representation for type interactionLimitResponse
    interactionLimitResponse i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable
    // Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
    itemItemInteractionLimitsGetResponseMember1 ItemItemInteractionLimitsGetResponseMember1able
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
        if val, err := parseNode.GetObjectValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateInteractionLimitResponseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable); ok {
                result.SetInteractionLimitResponse(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemInteractionLimitsGetResponseMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemInteractionLimitsGetResponseMember1able); ok {
                result.SetItemItemInteractionLimitsGetResponseMember1(cast)
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
func (m *InteractionLimitsResponse) GetInteractionLimitResponse()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable) {
    return m.interactionLimitResponse
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *InteractionLimitsResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemItemInteractionLimitsGetResponseMember1 gets the ItemItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) GetItemItemInteractionLimitsGetResponseMember1()(ItemItemInteractionLimitsGetResponseMember1able) {
    return m.itemItemInteractionLimitsGetResponseMember1
}
// Serialize serializes information the current object
func (m *InteractionLimitsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInteractionLimitResponse() != nil {
        err := writer.WriteObjectValue("", m.GetInteractionLimitResponse())
        if err != nil {
            return err
        }
    } else if m.GetItemItemInteractionLimitsGetResponseMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemInteractionLimitsGetResponseMember1())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteractionLimitResponse sets the interactionLimitResponse property value. Composed type representation for type interactionLimitResponse
func (m *InteractionLimitsResponse) SetInteractionLimitResponse(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)() {
    m.interactionLimitResponse = value
}
// SetItemItemInteractionLimitsGetResponseMember1 sets the ItemItemInteractionLimitsGetResponseMember1 property value. Composed type representation for type ItemItemInteractionLimitsGetResponseMember1
func (m *InteractionLimitsResponse) SetItemItemInteractionLimitsGetResponseMember1(value ItemItemInteractionLimitsGetResponseMember1able)() {
    m.itemItemInteractionLimitsGetResponseMember1 = value
}
// ItemItemInteractionLimitsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemInteractionLimitsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemInteractionLimitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemInteractionLimitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemInteractionLimitsRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemInteractionLimitsRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InteractionLimitsGetResponseable 
type InteractionLimitsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)
    GetItemItemInteractionLimitsGetResponseMember1()(ItemItemInteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)()
    SetItemItemInteractionLimitsGetResponseMember1(value ItemItemInteractionLimitsGetResponseMember1able)()
}
// InteractionLimitsResponseable 
type InteractionLimitsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractionLimitResponse()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)
    GetItemItemInteractionLimitsGetResponseMember1()(ItemItemInteractionLimitsGetResponseMember1able)
    SetInteractionLimitResponse(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable)()
    SetItemItemInteractionLimitsGetResponseMember1(value ItemItemInteractionLimitsGetResponseMember1able)()
}
// NewItemItemInteractionLimitsRequestBuilderInternal instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewItemItemInteractionLimitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemInteractionLimitsRequestBuilder) {
    m := &ItemItemInteractionLimitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/interaction-limits", pathParameters),
    }
    return m
}
// NewItemItemInteractionLimitsRequestBuilder instantiates a new InteractionLimitsRequestBuilder and sets the default values.
func NewItemItemInteractionLimitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemInteractionLimitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemInteractionLimitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes all interaction restrictions from the given repository. You must have owner or admin access to remove restrictions. If the interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/repos#remove-interaction-restrictions-for-a-repository
func (m *ItemItemInteractionLimitsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemInteractionLimitsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get shows which type of GitHub user can interact with this repository and when the restriction expires. If there are no restrictions, you will see an empty response.
// Deprecated: This method is obsolete. Use GetAsInteractionLimitsGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/repos#get-interaction-restrictions-for-a-repository
func (m *ItemItemInteractionLimitsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemInteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsResponseable, error) {
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
// GetAsInteractionLimitsGetResponse shows which type of GitHub user can interact with this repository and when the restriction expires. If there are no restrictions, you will see an empty response.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/repos#get-interaction-restrictions-for-a-repository
func (m *ItemItemInteractionLimitsRequestBuilder) GetAsInteractionLimitsGetResponse(ctx context.Context, requestConfiguration *ItemItemInteractionLimitsRequestBuilderGetRequestConfiguration)(InteractionLimitsGetResponseable, error) {
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
// Put temporarily restricts interactions to a certain type of GitHub user within the given repository. You must have owner or admin access to set these restrictions. If an interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/interactions/repos#set-interaction-restrictions-for-a-repository
func (m *ItemItemInteractionLimitsRequestBuilder) Put(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitable, requestConfiguration *ItemItemInteractionLimitsRequestBuilderPutRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateInteractionLimitResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitResponseable), nil
}
// ToDeleteRequestInformation removes all interaction restrictions from the given repository. You must have owner or admin access to remove restrictions. If the interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.
func (m *ItemItemInteractionLimitsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemInteractionLimitsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// ToGetRequestInformation shows which type of GitHub user can interact with this repository and when the restriction expires. If there are no restrictions, you will see an empty response.
func (m *ItemItemInteractionLimitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemInteractionLimitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation temporarily restricts interactions to a certain type of GitHub user within the given repository. You must have owner or admin access to set these restrictions. If an interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.
func (m *ItemItemInteractionLimitsRequestBuilder) ToPutRequestInformation(ctx context.Context, body i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.InteractionLimitable, requestConfiguration *ItemItemInteractionLimitsRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemInteractionLimitsRequestBuilder) WithUrl(rawUrl string)(*ItemItemInteractionLimitsRequestBuilder) {
    return NewItemItemInteractionLimitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
