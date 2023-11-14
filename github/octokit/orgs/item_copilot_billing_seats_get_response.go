package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSeatsGetResponse 
type ItemCopilotBillingSeatsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The seats property
    seats []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable
    // Total number of Copilot For Business seats for the organization currently being billed.
    total_seats *int32
}
// NewItemCopilotBillingSeatsGetResponse instantiates a new ItemCopilotBillingSeatsGetResponse and sets the default values.
func NewItemCopilotBillingSeatsGetResponse()(*ItemCopilotBillingSeatsGetResponse) {
    m := &ItemCopilotBillingSeatsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSeatsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSeatsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSeatsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCopilotBillingSeatsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCopilotBillingSeatsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCopilotSeatDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable)
                }
            }
            m.SetSeats(res)
        }
        return nil
    }
    res["total_seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSeats(val)
        }
        return nil
    }
    return res
}
// GetSeats gets the seats property value. The seats property
func (m *ItemCopilotBillingSeatsGetResponse) GetSeats()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable) {
    return m.seats
}
// GetTotalSeats gets the total_seats property value. Total number of Copilot For Business seats for the organization currently being billed.
func (m *ItemCopilotBillingSeatsGetResponse) GetTotalSeats()(*int32) {
    return m.total_seats
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSeatsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSeats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSeats()))
        for i, v := range m.GetSeats() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("seats", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_seats", m.GetTotalSeats())
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
func (m *ItemCopilotBillingSeatsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSeats sets the seats property value. The seats property
func (m *ItemCopilotBillingSeatsGetResponse) SetSeats(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable)() {
    m.seats = value
}
// SetTotalSeats sets the total_seats property value. Total number of Copilot For Business seats for the organization currently being billed.
func (m *ItemCopilotBillingSeatsGetResponse) SetTotalSeats(value *int32)() {
    m.total_seats = value
}
// ItemCopilotBillingSeatsGetResponseable 
type ItemCopilotBillingSeatsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSeats()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable)
    GetTotalSeats()(*int32)
    SetSeats(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CopilotSeatDetailsable)()
    SetTotalSeats(value *int32)()
}
