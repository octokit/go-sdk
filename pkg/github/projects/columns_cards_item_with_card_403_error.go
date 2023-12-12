package projects

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemWithCard_403Error 
type ColumnsCardsItemWithCard_403Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The documentation_url property
    documentation_url *string
    // The errors property
    errors []string
    // The message property
    message *string
}
// NewColumnsCardsItemWithCard_403Error instantiates a new ColumnsCardsItemWithCard_403Error and sets the default values.
func NewColumnsCardsItemWithCard_403Error()(*ColumnsCardsItemWithCard_403Error) {
    m := &ColumnsCardsItemWithCard_403Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnsCardsItemWithCard_403ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnsCardsItemWithCard_403ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemWithCard_403Error(), nil
}
// Error the primary error message.
func (m *ColumnsCardsItemWithCard_403Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnsCardsItemWithCard_403Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDocumentationUrl gets the documentation_url property value. The documentation_url property
func (m *ColumnsCardsItemWithCard_403Error) GetDocumentationUrl()(*string) {
    return m.documentation_url
}
// GetErrors gets the errors property value. The errors property
func (m *ColumnsCardsItemWithCard_403Error) GetErrors()([]string) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnsCardsItemWithCard_403Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["documentation_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
func (m *ColumnsCardsItemWithCard_403Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *ColumnsCardsItemWithCard_403Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("documentation_url", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        err := writer.WriteCollectionOfStringValues("errors", m.GetErrors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnsCardsItemWithCard_403Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDocumentationUrl sets the documentation_url property value. The documentation_url property
func (m *ColumnsCardsItemWithCard_403Error) SetDocumentationUrl(value *string)() {
    m.documentation_url = value
}
// SetErrors sets the errors property value. The errors property
func (m *ColumnsCardsItemWithCard_403Error) SetErrors(value []string)() {
    m.errors = value
}
// SetMessage sets the message property value. The message property
func (m *ColumnsCardsItemWithCard_403Error) SetMessage(value *string)() {
    m.message = value
}
// ColumnsCardsItemWithCard_403Errorable 
type ColumnsCardsItemWithCard_403Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDocumentationUrl()(*string)
    GetErrors()([]string)
    GetMessage()(*string)
    SetDocumentationUrl(value *string)()
    SetErrors(value []string)()
    SetMessage(value *string)()
}
