package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchRequiredStatusCheck_checks 
type ProtectedBranchRequiredStatusCheck_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The app_id property
    app_id *int32
    // The context property
    context *string
}
// NewProtectedBranchRequiredStatusCheck_checks instantiates a new protectedBranchRequiredStatusCheck_checks and sets the default values.
func NewProtectedBranchRequiredStatusCheck_checks()(*ProtectedBranchRequiredStatusCheck_checks) {
    m := &ProtectedBranchRequiredStatusCheck_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchRequiredStatusCheck_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectedBranchRequiredStatusCheck_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranchRequiredStatusCheck_checks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProtectedBranchRequiredStatusCheck_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the app_id property value. The app_id property
func (m *ProtectedBranchRequiredStatusCheck_checks) GetAppId()(*int32) {
    return m.app_id
}
// GetContext gets the context property value. The context property
func (m *ProtectedBranchRequiredStatusCheck_checks) GetContext()(*string) {
    return m.context
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectedBranchRequiredStatusCheck_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["app_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["context"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContext(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtectedBranchRequiredStatusCheck_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("app_id", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("context", m.GetContext())
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
func (m *ProtectedBranchRequiredStatusCheck_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the app_id property value. The app_id property
func (m *ProtectedBranchRequiredStatusCheck_checks) SetAppId(value *int32)() {
    m.app_id = value
}
// SetContext sets the context property value. The context property
func (m *ProtectedBranchRequiredStatusCheck_checks) SetContext(value *string)() {
    m.context = value
}
// ProtectedBranchRequiredStatusCheck_checksable 
type ProtectedBranchRequiredStatusCheck_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*int32)
    GetContext()(*string)
    SetAppId(value *int32)()
    SetContext(value *string)()
}
