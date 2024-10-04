package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomProperty custom property defined on an organization
type CustomProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
    allowed_values []string
    // Default value of the property
    default_value CustomProperty_CustomProperty_default_valueable
    // Short description of the property
    description *string
    // The name of the property
    property_name *string
    // Whether the property is required.
    required *bool
    // The URL that can be used to fetch, update, or delete info about this property via the API.
    url *string
    // The type of the value for the property
    value_type *CustomProperty_value_type
    // Who can edit the values of the property
    values_editable_by *CustomProperty_values_editable_by
}
// CustomProperty_CustomProperty_default_value composed type wrapper for classes string, []string
type CustomProperty_CustomProperty_default_value struct {
    // Composed type representation for type string
    customProperty_default_valueString *string
    // Composed type representation for type []string
    string []string
}
// NewCustomProperty_CustomProperty_default_value instantiates a new CustomProperty_CustomProperty_default_value and sets the default values.
func NewCustomProperty_CustomProperty_default_value()(*CustomProperty_CustomProperty_default_value) {
    m := &CustomProperty_CustomProperty_default_value{
    }
    return m
}
// CreateCustomProperty_CustomProperty_default_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomProperty_CustomProperty_default_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCustomProperty_CustomProperty_default_value()
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
        result.SetCustomPropertyDefaultValueString(val)
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
// GetCustomPropertyDefaultValueString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *CustomProperty_CustomProperty_default_value) GetCustomPropertyDefaultValueString()(*string) {
    return m.customProperty_default_valueString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomProperty_CustomProperty_default_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CustomProperty_CustomProperty_default_value) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type []string
// returns a []string when successful
func (m *CustomProperty_CustomProperty_default_value) GetString()([]string) {
    return m.string
}
// Serialize serializes information the current object
func (m *CustomProperty_CustomProperty_default_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustomPropertyDefaultValueString() != nil {
        err := writer.WriteStringValue("", m.GetCustomPropertyDefaultValueString())
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
// SetCustomPropertyDefaultValueString sets the string property value. Composed type representation for type string
func (m *CustomProperty_CustomProperty_default_value) SetCustomPropertyDefaultValueString(value *string)() {
    m.customProperty_default_valueString = value
}
// SetString sets the string property value. Composed type representation for type []string
func (m *CustomProperty_CustomProperty_default_value) SetString(value []string)() {
    m.string = value
}
type CustomProperty_CustomProperty_default_valueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomPropertyDefaultValueString()(*string)
    GetString()([]string)
    SetCustomPropertyDefaultValueString(value *string)()
    SetString(value []string)()
}
// NewCustomProperty instantiates a new CustomProperty and sets the default values.
func NewCustomProperty()(*CustomProperty) {
    m := &CustomProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedValues gets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
// returns a []string when successful
func (m *CustomProperty) GetAllowedValues()([]string) {
    return m.allowed_values
}
// GetDefaultValue gets the default_value property value. Default value of the property
// returns a CustomProperty_CustomProperty_default_valueable when successful
func (m *CustomProperty) GetDefaultValue()(CustomProperty_CustomProperty_default_valueable) {
    return m.default_value
}
// GetDescription gets the description property value. Short description of the property
// returns a *string when successful
func (m *CustomProperty) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateCustomProperty_CustomProperty_default_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(CustomProperty_CustomProperty_default_valueable))
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
    res["value_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomProperty_value_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*CustomProperty_value_type))
        }
        return nil
    }
    res["values_editable_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCustomProperty_values_editable_by)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuesEditableBy(val.(*CustomProperty_values_editable_by))
        }
        return nil
    }
    return res
}
// GetPropertyName gets the property_name property value. The name of the property
// returns a *string when successful
func (m *CustomProperty) GetPropertyName()(*string) {
    return m.property_name
}
// GetRequired gets the required property value. Whether the property is required.
// returns a *bool when successful
func (m *CustomProperty) GetRequired()(*bool) {
    return m.required
}
// GetUrl gets the url property value. The URL that can be used to fetch, update, or delete info about this property via the API.
// returns a *string when successful
func (m *CustomProperty) GetUrl()(*string) {
    return m.url
}
// GetValuesEditableBy gets the values_editable_by property value. Who can edit the values of the property
// returns a *CustomProperty_values_editable_by when successful
func (m *CustomProperty) GetValuesEditableBy()(*CustomProperty_values_editable_by) {
    return m.values_editable_by
}
// GetValueType gets the value_type property value. The type of the value for the property
// returns a *CustomProperty_value_type when successful
func (m *CustomProperty) GetValueType()(*CustomProperty_value_type) {
    return m.value_type
}
// Serialize serializes information the current object
func (m *CustomProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("property_name", m.GetPropertyName())
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
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    if m.GetValuesEditableBy() != nil {
        cast := (*m.GetValuesEditableBy()).String()
        err := writer.WriteStringValue("values_editable_by", &cast)
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
func (m *CustomProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedValues sets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
func (m *CustomProperty) SetAllowedValues(value []string)() {
    m.allowed_values = value
}
// SetDefaultValue sets the default_value property value. Default value of the property
func (m *CustomProperty) SetDefaultValue(value CustomProperty_CustomProperty_default_valueable)() {
    m.default_value = value
}
// SetDescription sets the description property value. Short description of the property
func (m *CustomProperty) SetDescription(value *string)() {
    m.description = value
}
// SetPropertyName sets the property_name property value. The name of the property
func (m *CustomProperty) SetPropertyName(value *string)() {
    m.property_name = value
}
// SetRequired sets the required property value. Whether the property is required.
func (m *CustomProperty) SetRequired(value *bool)() {
    m.required = value
}
// SetUrl sets the url property value. The URL that can be used to fetch, update, or delete info about this property via the API.
func (m *CustomProperty) SetUrl(value *string)() {
    m.url = value
}
// SetValuesEditableBy sets the values_editable_by property value. Who can edit the values of the property
func (m *CustomProperty) SetValuesEditableBy(value *CustomProperty_values_editable_by)() {
    m.values_editable_by = value
}
// SetValueType sets the value_type property value. The type of the value for the property
func (m *CustomProperty) SetValueType(value *CustomProperty_value_type)() {
    m.value_type = value
}
type CustomPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedValues()([]string)
    GetDefaultValue()(CustomProperty_CustomProperty_default_valueable)
    GetDescription()(*string)
    GetPropertyName()(*string)
    GetRequired()(*bool)
    GetUrl()(*string)
    GetValuesEditableBy()(*CustomProperty_values_editable_by)
    GetValueType()(*CustomProperty_value_type)
    SetAllowedValues(value []string)()
    SetDefaultValue(value CustomProperty_CustomProperty_default_valueable)()
    SetDescription(value *string)()
    SetPropertyName(value *string)()
    SetRequired(value *bool)()
    SetUrl(value *string)()
    SetValuesEditableBy(value *CustomProperty_values_editable_by)()
    SetValueType(value *CustomProperty_value_type)()
}
