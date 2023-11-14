package user

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesItemMachinesGetResponse 
type CodespacesItemMachinesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The machines property
    machines []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable
    // The total_count property
    total_count *int32
}
// NewCodespacesItemMachinesGetResponse instantiates a new CodespacesItemMachinesGetResponse and sets the default values.
func NewCodespacesItemMachinesGetResponse()(*CodespacesItemMachinesGetResponse) {
    m := &CodespacesItemMachinesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespacesItemMachinesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesItemMachinesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemMachinesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodespacesItemMachinesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CodespacesItemMachinesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["machines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateCodespaceMachineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable)
                }
            }
            m.SetMachines(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetMachines gets the machines property value. The machines property
func (m *CodespacesItemMachinesGetResponse) GetMachines()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable) {
    return m.machines
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *CodespacesItemMachinesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *CodespacesItemMachinesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMachines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMachines()))
        for i, v := range m.GetMachines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("machines", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *CodespacesItemMachinesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMachines sets the machines property value. The machines property
func (m *CodespacesItemMachinesGetResponse) SetMachines(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable)() {
    m.machines = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *CodespacesItemMachinesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// CodespacesItemMachinesGetResponseable 
type CodespacesItemMachinesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMachines()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable)
    GetTotalCount()(*int32)
    SetMachines(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CodespaceMachineable)()
    SetTotalCount(value *int32)()
}
