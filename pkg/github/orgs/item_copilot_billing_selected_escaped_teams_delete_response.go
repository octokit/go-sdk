package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_teamsDeleteResponse the total number of seats set to "pending cancellation" for members of the specified team(s).
type ItemCopilotBillingSelected_teamsDeleteResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seats_cancelled property
    seats_cancelled *int32
}
// NewItemCopilotBillingSelected_teamsDeleteResponse instantiates a new ItemCopilotBillingSelected_teamsDeleteResponse and sets the default values.
func NewItemCopilotBillingSelected_teamsDeleteResponse()(*ItemCopilotBillingSelected_teamsDeleteResponse) {
    m := &ItemCopilotBillingSelected_teamsDeleteResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_teamsDeleteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCopilotBillingSelected_teamsDeleteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_teamsDeleteResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// returns a *int32 when successful
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) GetSeatsCancelled()(*int32) {
    return m.seats_cancelled
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeatsCancelled sets the seats_cancelled property value. The seats_cancelled property
func (m *ItemCopilotBillingSelected_teamsDeleteResponse) SetSeatsCancelled(value *int32)() {
    m.seats_cancelled = value
}
type ItemCopilotBillingSelected_teamsDeleteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSeatsCancelled()(*int32)
    SetSeatsCancelled(value *int32)()
}
