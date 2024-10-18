package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActionsRunnerGroupsItemRunnersPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of runner IDs to add to the runner group.
    runners []int32
}
// NewItemActionsRunnerGroupsItemRunnersPutRequestBody instantiates a new ItemActionsRunnerGroupsItemRunnersPutRequestBody and sets the default values.
func NewItemActionsRunnerGroupsItemRunnersPutRequestBody()(*ItemActionsRunnerGroupsItemRunnersPutRequestBody) {
    m := &ItemActionsRunnerGroupsItemRunnersPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemActionsRunnerGroupsItemRunnersPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActionsRunnerGroupsItemRunnersPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActionsRunnerGroupsItemRunnersPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["runners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetRunners(res)
        }
        return nil
    }
    return res
}
// GetRunners gets the runners property value. List of runner IDs to add to the runner group.
// returns a []int32 when successful
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) GetRunners()([]int32) {
    return m.runners
}
// Serialize serializes information the current object
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRunners() != nil {
        err := writer.WriteCollectionOfInt32Values("runners", m.GetRunners())
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
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRunners sets the runners property value. List of runner IDs to add to the runner group.
func (m *ItemActionsRunnerGroupsItemRunnersPutRequestBody) SetRunners(value []int32)() {
    m.runners = value
}
type ItemActionsRunnerGroupsItemRunnersPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRunners()([]int32)
    SetRunners(value []int32)()
}
