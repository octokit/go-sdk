package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Codespace_runtime_constraints 
type Codespace_runtime_constraints struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The privacy settings a user can select from when forwarding a port.
    allowed_port_privacy_settings []string
}
// NewCodespace_runtime_constraints instantiates a new codespace_runtime_constraints and sets the default values.
func NewCodespace_runtime_constraints()(*Codespace_runtime_constraints) {
    m := &Codespace_runtime_constraints{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodespace_runtime_constraintsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespace_runtime_constraintsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespace_runtime_constraints(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Codespace_runtime_constraints) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedPortPrivacySettings gets the allowed_port_privacy_settings property value. The privacy settings a user can select from when forwarding a port.
func (m *Codespace_runtime_constraints) GetAllowedPortPrivacySettings()([]string) {
    return m.allowed_port_privacy_settings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Codespace_runtime_constraints) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_port_privacy_settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedPortPrivacySettings(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Codespace_runtime_constraints) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedPortPrivacySettings() != nil {
        err := writer.WriteCollectionOfStringValues("allowed_port_privacy_settings", m.GetAllowedPortPrivacySettings())
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
func (m *Codespace_runtime_constraints) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedPortPrivacySettings sets the allowed_port_privacy_settings property value. The privacy settings a user can select from when forwarding a port.
func (m *Codespace_runtime_constraints) SetAllowedPortPrivacySettings(value []string)() {
    m.allowed_port_privacy_settings = value
}
// Codespace_runtime_constraintsable 
type Codespace_runtime_constraintsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedPortPrivacySettings()([]string)
    SetAllowedPortPrivacySettings(value []string)()
}
