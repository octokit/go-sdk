package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationCreateIssueType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Color for the issue type.
    color *OrganizationCreateIssueType_color
    // Description of the issue type.
    description *string
    // Whether or not the issue type is enabled at the organization level.
    is_enabled *bool
    // Whether or not the issue type is restricted to issues in private repositories.
    is_private *bool
    // Name of the issue type.
    name *string
}
// NewOrganizationCreateIssueType instantiates a new OrganizationCreateIssueType and sets the default values.
func NewOrganizationCreateIssueType()(*OrganizationCreateIssueType) {
    m := &OrganizationCreateIssueType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationCreateIssueTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationCreateIssueTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationCreateIssueType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationCreateIssueType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. Color for the issue type.
// returns a *OrganizationCreateIssueType_color when successful
func (m *OrganizationCreateIssueType) GetColor()(*OrganizationCreateIssueType_color) {
    return m.color
}
// GetDescription gets the description property value. Description of the issue type.
// returns a *string when successful
func (m *OrganizationCreateIssueType) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationCreateIssueType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationCreateIssueType_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*OrganizationCreateIssueType_color))
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
    res["is_enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["is_private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivate(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the is_enabled property value. Whether or not the issue type is enabled at the organization level.
// returns a *bool when successful
func (m *OrganizationCreateIssueType) GetIsEnabled()(*bool) {
    return m.is_enabled
}
// GetIsPrivate gets the is_private property value. Whether or not the issue type is restricted to issues in private repositories.
// returns a *bool when successful
func (m *OrganizationCreateIssueType) GetIsPrivate()(*bool) {
    return m.is_private
}
// GetName gets the name property value. Name of the issue type.
// returns a *string when successful
func (m *OrganizationCreateIssueType) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *OrganizationCreateIssueType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err := writer.WriteStringValue("color", &cast)
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
        err := writer.WriteBoolValue("is_enabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_private", m.GetIsPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *OrganizationCreateIssueType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. Color for the issue type.
func (m *OrganizationCreateIssueType) SetColor(value *OrganizationCreateIssueType_color)() {
    m.color = value
}
// SetDescription sets the description property value. Description of the issue type.
func (m *OrganizationCreateIssueType) SetDescription(value *string)() {
    m.description = value
}
// SetIsEnabled sets the is_enabled property value. Whether or not the issue type is enabled at the organization level.
func (m *OrganizationCreateIssueType) SetIsEnabled(value *bool)() {
    m.is_enabled = value
}
// SetIsPrivate sets the is_private property value. Whether or not the issue type is restricted to issues in private repositories.
func (m *OrganizationCreateIssueType) SetIsPrivate(value *bool)() {
    m.is_private = value
}
// SetName sets the name property value. Name of the issue type.
func (m *OrganizationCreateIssueType) SetName(value *string)() {
    m.name = value
}
type OrganizationCreateIssueTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*OrganizationCreateIssueType_color)
    GetDescription()(*string)
    GetIsEnabled()(*bool)
    GetIsPrivate()(*bool)
    GetName()(*string)
    SetColor(value *OrganizationCreateIssueType_color)()
    SetDescription(value *string)()
    SetIsEnabled(value *bool)()
    SetIsPrivate(value *bool)()
    SetName(value *string)()
}
