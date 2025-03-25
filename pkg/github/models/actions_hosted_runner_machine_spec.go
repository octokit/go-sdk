package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsHostedRunnerMachineSpec provides details of a particular machine spec.
type ActionsHostedRunnerMachineSpec struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of cores.
    cpu_cores *int32
    // The ID used for the `size` parameter when creating a new runner.
    id *string
    // The available RAM for the machine spec.
    memory_gb *int32
    // The available SSD storage for the machine spec.
    storage_gb *int32
}
// NewActionsHostedRunnerMachineSpec instantiates a new ActionsHostedRunnerMachineSpec and sets the default values.
func NewActionsHostedRunnerMachineSpec()(*ActionsHostedRunnerMachineSpec) {
    m := &ActionsHostedRunnerMachineSpec{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsHostedRunnerMachineSpecFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsHostedRunnerMachineSpecFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsHostedRunnerMachineSpec(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsHostedRunnerMachineSpec) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuCores gets the cpu_cores property value. The number of cores.
// returns a *int32 when successful
func (m *ActionsHostedRunnerMachineSpec) GetCpuCores()(*int32) {
    return m.cpu_cores
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsHostedRunnerMachineSpec) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpu_cores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuCores(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["memory_gb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemoryGb(val)
        }
        return nil
    }
    res["storage_gb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageGb(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID used for the `size` parameter when creating a new runner.
// returns a *string when successful
func (m *ActionsHostedRunnerMachineSpec) GetId()(*string) {
    return m.id
}
// GetMemoryGb gets the memory_gb property value. The available RAM for the machine spec.
// returns a *int32 when successful
func (m *ActionsHostedRunnerMachineSpec) GetMemoryGb()(*int32) {
    return m.memory_gb
}
// GetStorageGb gets the storage_gb property value. The available SSD storage for the machine spec.
// returns a *int32 when successful
func (m *ActionsHostedRunnerMachineSpec) GetStorageGb()(*int32) {
    return m.storage_gb
}
// Serialize serializes information the current object
func (m *ActionsHostedRunnerMachineSpec) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("cpu_cores", m.GetCpuCores())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("memory_gb", m.GetMemoryGb())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("storage_gb", m.GetStorageGb())
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
func (m *ActionsHostedRunnerMachineSpec) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuCores sets the cpu_cores property value. The number of cores.
func (m *ActionsHostedRunnerMachineSpec) SetCpuCores(value *int32)() {
    m.cpu_cores = value
}
// SetId sets the id property value. The ID used for the `size` parameter when creating a new runner.
func (m *ActionsHostedRunnerMachineSpec) SetId(value *string)() {
    m.id = value
}
// SetMemoryGb sets the memory_gb property value. The available RAM for the machine spec.
func (m *ActionsHostedRunnerMachineSpec) SetMemoryGb(value *int32)() {
    m.memory_gb = value
}
// SetStorageGb sets the storage_gb property value. The available SSD storage for the machine spec.
func (m *ActionsHostedRunnerMachineSpec) SetStorageGb(value *int32)() {
    m.storage_gb = value
}
type ActionsHostedRunnerMachineSpecable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuCores()(*int32)
    GetId()(*string)
    GetMemoryGb()(*int32)
    GetStorageGb()(*int32)
    SetCpuCores(value *int32)()
    SetId(value *string)()
    SetMemoryGb(value *int32)()
    SetStorageGb(value *int32)()
}
