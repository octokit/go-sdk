package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationRole organization roles
type OrganizationRole struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The system role from which this role inherits permissions.
    base_role *OrganizationRole_base_role
    // The date and time the role was created.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A short description about who this role is for or what permissions it grants.
    description *string
    // The unique identifier of the role.
    id *int64
    // The name of the role.
    name *string
    // A GitHub user.
    organization NullableSimpleUserable
    // A list of permissions included in this role.
    permissions []string
    // Source answers the question, "where did this role come from?"
    source *OrganizationRole_source
    // The date and time the role was last updated.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewOrganizationRole instantiates a new OrganizationRole and sets the default values.
func NewOrganizationRole()(*OrganizationRole) {
    m := &OrganizationRole{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationRole(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationRole) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBaseRole gets the base_role property value. The system role from which this role inherits permissions.
// returns a *OrganizationRole_base_role when successful
func (m *OrganizationRole) GetBaseRole()(*OrganizationRole_base_role) {
    return m.base_role
}
// GetCreatedAt gets the created_at property value. The date and time the role was created.
// returns a *Time when successful
func (m *OrganizationRole) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDescription gets the description property value. A short description about who this role is for or what permissions it grants.
// returns a *string when successful
func (m *OrganizationRole) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base_role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationRole_base_role)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseRole(val.(*OrganizationRole_base_role))
        }
        return nil
    }
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPermissions(res)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationRole_source)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*OrganizationRole_source))
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
    return res
}
// GetId gets the id property value. The unique identifier of the role.
// returns a *int64 when successful
func (m *OrganizationRole) GetId()(*int64) {
    return m.id
}
// GetName gets the name property value. The name of the role.
// returns a *string when successful
func (m *OrganizationRole) GetName()(*string) {
    return m.name
}
// GetOrganization gets the organization property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *OrganizationRole) GetOrganization()(NullableSimpleUserable) {
    return m.organization
}
// GetPermissions gets the permissions property value. A list of permissions included in this role.
// returns a []string when successful
func (m *OrganizationRole) GetPermissions()([]string) {
    return m.permissions
}
// GetSource gets the source property value. Source answers the question, "where did this role come from?"
// returns a *OrganizationRole_source when successful
func (m *OrganizationRole) GetSource()(*OrganizationRole_source) {
    return m.source
}
// GetUpdatedAt gets the updated_at property value. The date and time the role was last updated.
// returns a *Time when successful
func (m *OrganizationRole) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *OrganizationRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBaseRole() != nil {
        cast := (*m.GetBaseRole()).String()
        err := writer.WriteStringValue("base_role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteInt64Value("id", m.GetId())
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
        err := writer.WriteObjectValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        err := writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err := writer.WriteStringValue("source", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationRole) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBaseRole sets the base_role property value. The system role from which this role inherits permissions.
func (m *OrganizationRole) SetBaseRole(value *OrganizationRole_base_role)() {
    m.base_role = value
}
// SetCreatedAt sets the created_at property value. The date and time the role was created.
func (m *OrganizationRole) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDescription sets the description property value. A short description about who this role is for or what permissions it grants.
func (m *OrganizationRole) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The unique identifier of the role.
func (m *OrganizationRole) SetId(value *int64)() {
    m.id = value
}
// SetName sets the name property value. The name of the role.
func (m *OrganizationRole) SetName(value *string)() {
    m.name = value
}
// SetOrganization sets the organization property value. A GitHub user.
func (m *OrganizationRole) SetOrganization(value NullableSimpleUserable)() {
    m.organization = value
}
// SetPermissions sets the permissions property value. A list of permissions included in this role.
func (m *OrganizationRole) SetPermissions(value []string)() {
    m.permissions = value
}
// SetSource sets the source property value. Source answers the question, "where did this role come from?"
func (m *OrganizationRole) SetSource(value *OrganizationRole_source)() {
    m.source = value
}
// SetUpdatedAt sets the updated_at property value. The date and time the role was last updated.
func (m *OrganizationRole) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type OrganizationRoleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseRole()(*OrganizationRole_base_role)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetId()(*int64)
    GetName()(*string)
    GetOrganization()(NullableSimpleUserable)
    GetPermissions()([]string)
    GetSource()(*OrganizationRole_source)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBaseRole(value *OrganizationRole_base_role)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetId(value *int64)()
    SetName(value *string)()
    SetOrganization(value NullableSimpleUserable)()
    SetPermissions(value []string)()
    SetSource(value *OrganizationRole_source)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
