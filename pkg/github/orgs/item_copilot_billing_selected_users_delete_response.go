package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_usersDeleteResponse the total number of seat assignments cancelled.
type ItemCopilotBillingSelected_usersDeleteResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seats_cancelled property
    seats_cancelled *int32
}
// NewItemCopilotBillingSelected_usersDeleteResponse instantiates a new ItemCopilotBillingSelected_usersDeleteResponse and sets the default values.
func NewItemCopilotBillingSelected_usersDeleteResponse()(*ItemCopilotBillingSelected_usersDeleteResponse) {
    m := &ItemCopilotBillingSelected_usersDeleteResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_usersDeleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSelected_usersDeleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_usersDeleteResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCopilotBillingSelected_usersDeleteResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCopilotBillingSelected_usersDeleteResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["seats_cancelled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatsCancelled(val)
        }
        return nil
    }
    return res
}
// GetSeatsCancelled gets the seats_cancelled property value. The seats_cancelled property
func (m *ItemCopilotBillingSelected_usersDeleteResponse) GetSeatsCancelled()(*int32) {
    return m.seats_cancelled
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_usersDeleteResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("seats_cancelled", m.GetSeatsCancelled())
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
func (m *ItemCopilotBillingSelected_usersDeleteResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeatsCancelled sets the seats_cancelled property value. The seats_cancelled property
func (m *ItemCopilotBillingSelected_usersDeleteResponse) SetSeatsCancelled(value *int32)() {
    m.seats_cancelled = value
}
// ItemCopilotBillingSelected_usersDeleteResponseable 
type ItemCopilotBillingSelected_usersDeleteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSeatsCancelled()(*int32)
    SetSeatsCancelled(value *int32)()
}
