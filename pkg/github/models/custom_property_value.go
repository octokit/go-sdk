package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomPropertyValue custom property name and associated value
type CustomPropertyValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the property
    property_name *string
    // The value assigned to the property
    value *string
}
// NewCustomPropertyValue instantiates a new customPropertyValue and sets the default values.
func NewCustomPropertyValue()(*CustomPropertyValue) {
    m := &CustomPropertyValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomPropertyValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomPropertyValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomPropertyValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomPropertyValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomPropertyValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["property_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetPropertyName gets the property_name property value. The name of the property
func (m *CustomPropertyValue) GetPropertyName()(*string) {
    return m.property_name
}
// GetValue gets the value property value. The value assigned to the property
func (m *CustomPropertyValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomPropertyValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("property_name", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *CustomPropertyValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPropertyName sets the property_name property value. The name of the property
func (m *CustomPropertyValue) SetPropertyName(value *string)() {
    m.property_name = value
}
// SetValue sets the value property value. The value assigned to the property
func (m *CustomPropertyValue) SetValue(value *string)() {
    m.value = value
}
// CustomPropertyValueable 
type CustomPropertyValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPropertyName()(*string)
    GetValue()(*string)
    SetPropertyName(value *string)()
    SetValue(value *string)()
}
