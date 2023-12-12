package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusCheckPolicy_checks 
type StatusCheckPolicy_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The app_id property
    app_id *int32
    // The context property
    context *string
}
// NewStatusCheckPolicy_checks instantiates a new statusCheckPolicy_checks and sets the default values.
func NewStatusCheckPolicy_checks()(*StatusCheckPolicy_checks) {
    m := &StatusCheckPolicy_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStatusCheckPolicy_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStatusCheckPolicy_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusCheckPolicy_checks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StatusCheckPolicy_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the app_id property value. The app_id property
func (m *StatusCheckPolicy_checks) GetAppId()(*int32) {
    return m.app_id
}
// GetContext gets the context property value. The context property
func (m *StatusCheckPolicy_checks) GetContext()(*string) {
    return m.context
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StatusCheckPolicy_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *StatusCheckPolicy_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *StatusCheckPolicy_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the app_id property value. The app_id property
func (m *StatusCheckPolicy_checks) SetAppId(value *int32)() {
    m.app_id = value
}
// SetContext sets the context property value. The context property
func (m *StatusCheckPolicy_checks) SetContext(value *string)() {
    m.context = value
}
// StatusCheckPolicy_checksable 
type StatusCheckPolicy_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*int32)
    GetContext()(*string)
    SetAppId(value *int32)()
    SetContext(value *string)()
}
