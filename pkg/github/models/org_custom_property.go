package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgCustomProperty custom property defined on an organization
type OrgCustomProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
    allowed_values []string
    // Default value of the property
    default_value *string
    // Short description of the property
    description *string
    // The name of the property
    property_name *string
    // Whether the property is required.
    required *bool
    // The type of the value for the property
    value_type *OrgCustomProperty_value_type
    // Who can edit the values of the property
    values_editable_by *OrgCustomProperty_values_editable_by
}
// NewOrgCustomProperty instantiates a new orgCustomProperty and sets the default values.
func NewOrgCustomProperty()(*OrgCustomProperty) {
    m := &OrgCustomProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgCustomPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrgCustomPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgCustomProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrgCustomProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedValues gets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
func (m *OrgCustomProperty) GetAllowedValues()([]string) {
    return m.allowed_values
}
// GetDefaultValue gets the default_value property value. Default value of the property
func (m *OrgCustomProperty) GetDefaultValue()(*string) {
    return m.default_value
}
// GetDescription gets the description property value. Short description of the property
func (m *OrgCustomProperty) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrgCustomProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
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
    res["value_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgCustomProperty_value_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*OrgCustomProperty_value_type))
        }
        return nil
    }
    res["values_editable_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgCustomProperty_values_editable_by)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuesEditableBy(val.(*OrgCustomProperty_values_editable_by))
        }
        return nil
    }
    return res
}
// GetPropertyName gets the property_name property value. The name of the property
func (m *OrgCustomProperty) GetPropertyName()(*string) {
    return m.property_name
}
// GetRequired gets the required property value. Whether the property is required.
func (m *OrgCustomProperty) GetRequired()(*bool) {
    return m.required
}
// GetValuesEditableBy gets the values_editable_by property value. Who can edit the values of the property
func (m *OrgCustomProperty) GetValuesEditableBy()(*OrgCustomProperty_values_editable_by) {
    return m.values_editable_by
}
// GetValueType gets the value_type property value. The type of the value for the property
func (m *OrgCustomProperty) GetValueType()(*OrgCustomProperty_value_type) {
    return m.value_type
}
// Serialize serializes information the current object
func (m *OrgCustomProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedValues() != nil {
        err := writer.WriteCollectionOfStringValues("allowed_values", m.GetAllowedValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("default_value", m.GetDefaultValue())
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
func (m *OrgCustomProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedValues sets the allowed_values property value. An ordered list of the allowed values of the property.The property can have up to 200 allowed values.
func (m *OrgCustomProperty) SetAllowedValues(value []string)() {
    m.allowed_values = value
}
// SetDefaultValue sets the default_value property value. Default value of the property
func (m *OrgCustomProperty) SetDefaultValue(value *string)() {
    m.default_value = value
}
// SetDescription sets the description property value. Short description of the property
func (m *OrgCustomProperty) SetDescription(value *string)() {
    m.description = value
}
// SetPropertyName sets the property_name property value. The name of the property
func (m *OrgCustomProperty) SetPropertyName(value *string)() {
    m.property_name = value
}
// SetRequired sets the required property value. Whether the property is required.
func (m *OrgCustomProperty) SetRequired(value *bool)() {
    m.required = value
}
// SetValuesEditableBy sets the values_editable_by property value. Who can edit the values of the property
func (m *OrgCustomProperty) SetValuesEditableBy(value *OrgCustomProperty_values_editable_by)() {
    m.values_editable_by = value
}
// SetValueType sets the value_type property value. The type of the value for the property
func (m *OrgCustomProperty) SetValueType(value *OrgCustomProperty_value_type)() {
    m.value_type = value
}
// OrgCustomPropertyable 
type OrgCustomPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedValues()([]string)
    GetDefaultValue()(*string)
    GetDescription()(*string)
    GetPropertyName()(*string)
    GetRequired()(*bool)
    GetValuesEditableBy()(*OrgCustomProperty_values_editable_by)
    GetValueType()(*OrgCustomProperty_value_type)
    SetAllowedValues(value []string)()
    SetDefaultValue(value *string)()
    SetDescription(value *string)()
    SetPropertyName(value *string)()
    SetRequired(value *bool)()
    SetValuesEditableBy(value *OrgCustomProperty_values_editable_by)()
    SetValueType(value *OrgCustomProperty_value_type)()
}
