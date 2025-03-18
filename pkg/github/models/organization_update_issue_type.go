package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationUpdateIssueType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Color for the issue type.
    color *OrganizationUpdateIssueType_color
    // Description of the issue type.
    description *string
    // Whether or not the issue type is enabled at the organization level.
    is_enabled *bool
    // Whether or not the issue type is restricted to issues in private repositories.
    is_private *bool
    // Name of the issue type.
    name *string
}
// NewOrganizationUpdateIssueType instantiates a new OrganizationUpdateIssueType and sets the default values.
func NewOrganizationUpdateIssueType()(*OrganizationUpdateIssueType) {
    m := &OrganizationUpdateIssueType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationUpdateIssueTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationUpdateIssueTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationUpdateIssueType(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationUpdateIssueType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. Color for the issue type.
// returns a *OrganizationUpdateIssueType_color when successful
func (m *OrganizationUpdateIssueType) GetColor()(*OrganizationUpdateIssueType_color) {
    return m.color
}
// GetDescription gets the description property value. Description of the issue type.
// returns a *string when successful
func (m *OrganizationUpdateIssueType) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationUpdateIssueType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationUpdateIssueType_color)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*OrganizationUpdateIssueType_color))
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
func (m *OrganizationUpdateIssueType) GetIsEnabled()(*bool) {
    return m.is_enabled
}
// GetIsPrivate gets the is_private property value. Whether or not the issue type is restricted to issues in private repositories.
// returns a *bool when successful
func (m *OrganizationUpdateIssueType) GetIsPrivate()(*bool) {
    return m.is_private
}
// GetName gets the name property value. Name of the issue type.
// returns a *string when successful
func (m *OrganizationUpdateIssueType) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *OrganizationUpdateIssueType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OrganizationUpdateIssueType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. Color for the issue type.
func (m *OrganizationUpdateIssueType) SetColor(value *OrganizationUpdateIssueType_color)() {
    m.color = value
}
// SetDescription sets the description property value. Description of the issue type.
func (m *OrganizationUpdateIssueType) SetDescription(value *string)() {
    m.description = value
}
// SetIsEnabled sets the is_enabled property value. Whether or not the issue type is enabled at the organization level.
func (m *OrganizationUpdateIssueType) SetIsEnabled(value *bool)() {
    m.is_enabled = value
}
// SetIsPrivate sets the is_private property value. Whether or not the issue type is restricted to issues in private repositories.
func (m *OrganizationUpdateIssueType) SetIsPrivate(value *bool)() {
    m.is_private = value
}
// SetName sets the name property value. Name of the issue type.
func (m *OrganizationUpdateIssueType) SetName(value *string)() {
    m.name = value
}
type OrganizationUpdateIssueTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*OrganizationUpdateIssueType_color)
    GetDescription()(*string)
    GetIsEnabled()(*bool)
    GetIsPrivate()(*bool)
    GetName()(*string)
    SetColor(value *OrganizationUpdateIssueType_color)()
    SetDescription(value *string)()
    SetIsEnabled(value *bool)()
    SetIsPrivate(value *bool)()
    SetName(value *string)()
}
