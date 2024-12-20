package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomPropertySetPayload custom property set payload
type CustomPropertySetPayload struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
    allowed_values []string
    // Default value of the property
    default_value CustomPropertySetPayload_CustomPropertySetPayload_default_valueable
    // Short description of the property
    description *string
    // Whether the property is required.
    required *bool
    // The type of the value for the property
    value_type *CustomPropertySetPayload_value_type
}
// CustomPropertySetPayload_CustomPropertySetPayload_default_value composed type wrapper for classes string, []string
type CustomPropertySetPayload_CustomPropertySetPayload_default_value struct {
    // Composed type representation for type string
    customPropertySetPayload_default_valueString *string
    // Composed type representation for type []string
    string []string
}
// NewCustomPropertySetPayload_CustomPropertySetPayload_default_value instantiates a new CustomPropertySetPayload_CustomPropertySetPayload_default_value and sets the default values.
func NewCustomPropertySetPayload_CustomPropertySetPayload_default_value()(*CustomPropertySetPayload_CustomPropertySetPayload_default_value) {
    m := &CustomPropertySetPayload_CustomPropertySetPayload_default_value{
    }
    return m
}
// CreateCustomPropertySetPayload_CustomPropertySetPayload_default_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomPropertySetPayload_CustomPropertySetPayload_default_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCustomPropertySetPayload_CustomPropertySetPayload_default_value()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetCustomPropertySetPayloadDefaultValueString(val)
    } else if val, err := parseNode.GetCollectionOfPrimitiveValues("string"); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]string, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = *(v.(*string))
            }
        }
        result.SetString(cast)
    }
    return result, nil
}
// GetCustomPropertySetPayloadDefaultValueString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) GetCustomPropertySetPayloadDefaultValueString()(*string) {
    return m.customPropertySetPayload_default_valueString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type []string
// returns a []string when successful
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) GetString()([]string) {
    return m.string
}
// Serialize serializes information the current object
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustomPropertySetPayloadDefaultValueString() != nil {
        err := writer.WriteStringValue("", m.GetCustomPropertySetPayloadDefaultValueString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteCollectionOfStringValues("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomPropertySetPayloadDefaultValueString sets the string property value. Composed type representation for type string
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) SetCustomPropertySetPayloadDefaultValueString(value *string)() {
    m.customPropertySetPayload_default_valueString = value
}
// SetString sets the string property value. Composed type representation for type []string
func (m *CustomPropertySetPayload_CustomPropertySetPayload_default_value) SetString(value []string)() {
    m.string = value
}
type CustomPropertySetPayload_CustomPropertySetPayload_default_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomPropertySetPayloadDefaultValueString()(*string)
    GetString()([]string)
    SetCustomPropertySetPayloadDefaultValueString(value *string)()
    SetString(value []string)()
}
// NewCustomPropertySetPayload instantiates a new CustomPropertySetPayload and sets the default values.
func NewCustomPropertySetPayload()(*CustomPropertySetPayload) {
    m := &CustomPropertySetPayload{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomPropertySetPayloadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomPropertySetPayloadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomPropertySetPayload(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomPropertySetPayload) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedValues gets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
// returns a []string when successful
func (m *CustomPropertySetPayload) GetAllowedValues()([]string) {
    return m.allowed_values
}
// GetDefaultValue gets the default_value property value. Default value of the property
// returns a CustomPropertySetPayload_CustomPropertySetPayload_default_valueable when successful
func (m *CustomPropertySetPayload) GetDefaultValue()(CustomPropertySetPayload_CustomPropertySetPayload_default_valueable) {
    return m.default_value
}
// GetDescription gets the description property value. Short description of the property
// returns a *string when successful
func (m *CustomPropertySetPayload) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomPropertySetPayload) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed_values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedValues(res)
        }
        return nil
    }
    res["default_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomPropertySetPayload_CustomPropertySetPayload_default_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(CustomPropertySetPayload_CustomPropertySetPayload_default_valueable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["required"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["value_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomPropertySetPayload_value_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*CustomPropertySetPayload_value_type))
        }
        return nil
    }
    return res
}
// GetRequired gets the required property value. Whether the property is required.
// returns a *bool when successful
func (m *CustomPropertySetPayload) GetRequired()(*bool) {
    return m.required
}
// GetValueType gets the value_type property value. The type of the value for the property
// returns a *CustomPropertySetPayload_value_type when successful
func (m *CustomPropertySetPayload) GetValueType()(*CustomPropertySetPayload_value_type) {
    return m.value_type
}
// Serialize serializes information the current object
func (m *CustomPropertySetPayload) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedValues() != nil {
        err := writer.WriteCollectionOfStringValues("allowed_values", m.GetAllowedValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("default_value", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("value_type", &cast)
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
func (m *CustomPropertySetPayload) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedValues sets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
func (m *CustomPropertySetPayload) SetAllowedValues(value []string)() {
    m.allowed_values = value
}
// SetDefaultValue sets the default_value property value. Default value of the property
func (m *CustomPropertySetPayload) SetDefaultValue(value CustomPropertySetPayload_CustomPropertySetPayload_default_valueable)() {
    m.default_value = value
}
// SetDescription sets the description property value. Short description of the property
func (m *CustomPropertySetPayload) SetDescription(value *string)() {
    m.description = value
}
// SetRequired sets the required property value. Whether the property is required.
func (m *CustomPropertySetPayload) SetRequired(value *bool)() {
    m.required = value
}
// SetValueType sets the value_type property value. The type of the value for the property
func (m *CustomPropertySetPayload) SetValueType(value *CustomPropertySetPayload_value_type)() {
    m.value_type = value
}
type CustomPropertySetPayloadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedValues()([]string)
    GetDefaultValue()(CustomPropertySetPayload_CustomPropertySetPayload_default_valueable)
    GetDescription()(*string)
    GetRequired()(*bool)
    GetValueType()(*CustomPropertySetPayload_value_type)
    SetAllowedValues(value []string)()
    SetDefaultValue(value CustomPropertySetPayload_CustomPropertySetPayload_default_valueable)()
    SetDescription(value *string)()
    SetRequired(value *bool)()
    SetValueType(value *CustomPropertySetPayload_value_type)()
}
