package projects

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    id34ad19071c4401276be58ddfcf31ab64ada4bb6069acb7835854cce34f9d366 "github.com/octokit/go-sdk/github/octokit/projects/columns/item/cards"
)

// ColumnsItemCardsRequestBuilder builds and executes requests for operations under \projects\columns\{column_id}\cards
type ColumnsItemCardsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CardsPostRequestBody composed type wrapper for classes ColumnsItemCardsPostRequestBodyMember1, ColumnsItemCardsPostRequestBodyMember2
type CardsPostRequestBody struct {
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
    cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 ColumnsItemCardsPostRequestBodyMember1able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
    cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 ColumnsItemCardsPostRequestBodyMember2able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
    columnsItemCardsPostRequestBodyMember1 ColumnsItemCardsPostRequestBodyMember1able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
    columnsItemCardsPostRequestBodyMember2 ColumnsItemCardsPostRequestBodyMember2able
}
// NewCardsPostRequestBody instantiates a new cardsPostRequestBody and sets the default values.
func NewCardsPostRequestBody()(*CardsPostRequestBody) {
    m := &CardsPostRequestBody{
    }
    return m
}
// CreateCardsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCardsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCardsPostRequestBody()
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
// GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 gets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
func (m *CardsPostRequestBody) GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able) {
    return m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1
}
// GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 gets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
func (m *CardsPostRequestBody) GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able) {
    return m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2
}
// GetColumnsItemCardsPostRequestBodyMember1 gets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
func (m *CardsPostRequestBody) GetColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able) {
    return m.columnsItemCardsPostRequestBodyMember1
}
// GetColumnsItemCardsPostRequestBodyMember2 gets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
func (m *CardsPostRequestBody) GetColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able) {
    return m.columnsItemCardsPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CardsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *CardsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *CardsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetColumnsItemCardsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetColumnsItemCardsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetColumnsItemCardsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetColumnsItemCardsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 sets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
func (m *CardsPostRequestBody) SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)() {
    m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 = value
}
// SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 sets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
func (m *CardsPostRequestBody) SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)() {
    m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 = value
}
// SetColumnsItemCardsPostRequestBodyMember1 sets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1
func (m *CardsPostRequestBody) SetColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)() {
    m.columnsItemCardsPostRequestBodyMember1 = value
}
// SetColumnsItemCardsPostRequestBodyMember2 sets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2
func (m *CardsPostRequestBody) SetColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)() {
    m.columnsItemCardsPostRequestBodyMember2 = value
}
// ColumnsItemCardsRequestBuilderGetQueryParameters lists the project cards in a project.
type ColumnsItemCardsRequestBuilderGetQueryParameters struct {
    // Filters the project cards that are returned by the card's state.
    // Deprecated: This property is deprecated, use archived_stateAsGetArchived_stateQueryParameterType instead
    Archived_state *string `uriparametername:"archived_state"`
    // Filters the project cards that are returned by the card's state.
    Archived_stateAsGetArchived_stateQueryParameterType *id34ad19071c4401276be58ddfcf31ab64ada4bb6069acb7835854cce34f9d366.GetArchived_stateQueryParameterType `uriparametername:"archived_state"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ColumnsItemCardsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsItemCardsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ColumnsItemCardsRequestBuilderGetQueryParameters
}
// ColumnsItemCardsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsItemCardsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ProjectCard422Error composed type wrapper for classes validationError, validationErrorSimple
type ProjectCard422Error struct {
    // Composed type representation for type validationError
    validationError i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable
    // Composed type representation for type validationErrorSimple
    validationErrorSimple i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable
}
// NewProjectCard422Error instantiates a new projectCard422Error and sets the default values.
func NewProjectCard422Error()(*ProjectCard422Error) {
    m := &ProjectCard422Error{
    }
    return m
}
// CreateProjectCard422ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProjectCard422ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewProjectCard422Error()
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ProjectCard422Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *ProjectCard422Error) GetIsComposedType()(bool) {
    return true
}
// GetValidationError gets the validationError property value. Composed type representation for type validationError
func (m *ProjectCard422Error) GetValidationError()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable) {
    return m.validationError
}
// GetValidationErrorSimple gets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *ProjectCard422Error) GetValidationErrorSimple()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable) {
    return m.validationErrorSimple
}
// Serialize serializes information the current object
func (m *ProjectCard422Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValidationError() != nil {
        err := writer.WriteObjectValue("", m.GetValidationError())
        if err != nil {
            return err
        }
    } else if m.GetValidationErrorSimple() != nil {
        err := writer.WriteObjectValue("", m.GetValidationErrorSimple())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValidationError sets the validationError property value. Composed type representation for type validationError
func (m *ProjectCard422Error) SetValidationError(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)() {
    m.validationError = value
}
// SetValidationErrorSimple sets the validationErrorSimple property value. Composed type representation for type validationErrorSimple
func (m *ProjectCard422Error) SetValidationErrorSimple(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)() {
    m.validationErrorSimple = value
}
// CardsPostRequestBodyable 
type CardsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able)
    GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able)
    GetColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able)
    GetColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able)
    SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)()
    SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)()
    SetColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)()
    SetColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)()
}
// ProjectCard422Errorable 
type ProjectCard422Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValidationError()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)
    GetValidationErrorSimple()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)
    SetValidationError(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorable)()
    SetValidationErrorSimple(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ValidationErrorSimpleable)()
}
// NewColumnsItemCardsRequestBuilderInternal instantiates a new CardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsItemCardsRequestBuilder) {
    m := &ColumnsItemCardsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/columns/{column_id}/cards{?archived_state*,per_page*,page*}", pathParameters),
    }
    return m
}
// NewColumnsItemCardsRequestBuilder instantiates a new CardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsItemCardsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsItemCardsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the project cards in a project.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#list-project-cards
func (m *ColumnsItemCardsRequestBuilder) Get(ctx context.Context, requestConfiguration *ColumnsItemCardsRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProjectCardable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateProjectCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProjectCardable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProjectCardable)
        }
    }
    return val, nil
}
// Post create a project card
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#create-a-project-card
func (m *ColumnsItemCardsRequestBuilder) Post(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *ColumnsItemCardsRequestBuilderPostRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProjectCardable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "403": i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateBasicErrorFromDiscriminatorValue,
        "422": CreateProjectCard422ErrorFromDiscriminatorValue,
        "503": CreateColumnsItemCardsProjectCard503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateProjectCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.ProjectCardable), nil
}
// ToGetRequestInformation lists the project cards in a project.
func (m *ColumnsItemCardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ColumnsItemCardsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
func (m *ColumnsItemCardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *ColumnsItemCardsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ColumnsItemCardsRequestBuilder) WithUrl(rawUrl string)(*ColumnsItemCardsRequestBuilder) {
    return NewColumnsItemCardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
