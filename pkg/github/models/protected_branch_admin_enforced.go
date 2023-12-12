package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchAdminEnforced protected Branch Admin Enforced
type ProtectedBranchAdminEnforced struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enabled property
    enabled *bool
    // The url property
    url *string
}
// NewProtectedBranchAdminEnforced instantiates a new protectedBranchAdminEnforced and sets the default values.
func NewProtectedBranchAdminEnforced()(*ProtectedBranchAdminEnforced) {
    m := &ProtectedBranchAdminEnforced{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchAdminEnforcedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectedBranchAdminEnforcedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranchAdminEnforced(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectedBranchAdminEnforced) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
func (m *ProtectedBranchAdminEnforced) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectedBranchAdminEnforced) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The url property
func (m *ProtectedBranchAdminEnforced) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ProtectedBranchAdminEnforced) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ProtectedBranchAdminEnforced) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ProtectedBranchAdminEnforced) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetUrl sets the url property value. The url property
func (m *ProtectedBranchAdminEnforced) SetUrl(value *string)() {
    m.url = value
}
// ProtectedBranchAdminEnforcedable 
type ProtectedBranchAdminEnforcedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetUrl()(*string)
    SetEnabled(value *bool)()
    SetUrl(value *string)()
}
