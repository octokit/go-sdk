package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgPrivateRegistryConfiguration private registry configuration for an organization
type OrgPrivateRegistryConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the private registry configuration.
    name *string
    // The registry type.
    registry_type *OrgPrivateRegistryConfiguration_registry_type
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The username to use when authenticating with the private registry.
    username *string
    // Which type of organization repositories have access to the private registry.
    visibility *OrgPrivateRegistryConfiguration_visibility
}
// NewOrgPrivateRegistryConfiguration instantiates a new OrgPrivateRegistryConfiguration and sets the default values.
func NewOrgPrivateRegistryConfiguration()(*OrgPrivateRegistryConfiguration) {
    m := &OrgPrivateRegistryConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgPrivateRegistryConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrgPrivateRegistryConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgPrivateRegistryConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrgPrivateRegistryConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *OrgPrivateRegistryConfiguration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrgPrivateRegistryConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["registry_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgPrivateRegistryConfiguration_registry_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistryType(val.(*OrgPrivateRegistryConfiguration_registry_type))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrgPrivateRegistryConfiguration_visibility)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val.(*OrgPrivateRegistryConfiguration_visibility))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the private registry configuration.
// returns a *string when successful
func (m *OrgPrivateRegistryConfiguration) GetName()(*string) {
    return m.name
}
// GetRegistryType gets the registry_type property value. The registry type.
// returns a *OrgPrivateRegistryConfiguration_registry_type when successful
func (m *OrgPrivateRegistryConfiguration) GetRegistryType()(*OrgPrivateRegistryConfiguration_registry_type) {
    return m.registry_type
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *OrgPrivateRegistryConfiguration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUsername gets the username property value. The username to use when authenticating with the private registry.
// returns a *string when successful
func (m *OrgPrivateRegistryConfiguration) GetUsername()(*string) {
    return m.username
}
// GetVisibility gets the visibility property value. Which type of organization repositories have access to the private registry.
// returns a *OrgPrivateRegistryConfiguration_visibility when successful
func (m *OrgPrivateRegistryConfiguration) GetVisibility()(*OrgPrivateRegistryConfiguration_visibility) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *OrgPrivateRegistryConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
    if m.GetRegistryType() != nil {
        cast := (*m.GetRegistryType()).String()
        err := writer.WriteStringValue("registry_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := (*m.GetVisibility()).String()
        err := writer.WriteStringValue("visibility", &cast)
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
func (m *OrgPrivateRegistryConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *OrgPrivateRegistryConfiguration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetName sets the name property value. The name of the private registry configuration.
func (m *OrgPrivateRegistryConfiguration) SetName(value *string)() {
    m.name = value
}
// SetRegistryType sets the registry_type property value. The registry type.
func (m *OrgPrivateRegistryConfiguration) SetRegistryType(value *OrgPrivateRegistryConfiguration_registry_type)() {
    m.registry_type = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *OrgPrivateRegistryConfiguration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUsername sets the username property value. The username to use when authenticating with the private registry.
func (m *OrgPrivateRegistryConfiguration) SetUsername(value *string)() {
    m.username = value
}
// SetVisibility sets the visibility property value. Which type of organization repositories have access to the private registry.
func (m *OrgPrivateRegistryConfiguration) SetVisibility(value *OrgPrivateRegistryConfiguration_visibility)() {
    m.visibility = value
}
type OrgPrivateRegistryConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetRegistryType()(*OrgPrivateRegistryConfiguration_registry_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUsername()(*string)
    GetVisibility()(*OrgPrivateRegistryConfiguration_visibility)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetRegistryType(value *OrgPrivateRegistryConfiguration_registry_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUsername(value *string)()
    SetVisibility(value *OrgPrivateRegistryConfiguration_visibility)()
}
