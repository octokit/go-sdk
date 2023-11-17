package orgs

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOrganizationRolesGetResponse 
type ItemOrganizationRolesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of organization roles available to the organization.
    roles []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable
    // The total number of organization roles available to the organization.
    total_count *int32
}
// NewItemOrganizationRolesGetResponse instantiates a new ItemOrganizationRolesGetResponse and sets the default values.
func NewItemOrganizationRolesGetResponse()(*ItemOrganizationRolesGetResponse) {
    m := &ItemOrganizationRolesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOrganizationRolesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOrganizationRolesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOrganizationRolesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOrganizationRolesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOrganizationRolesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateOrganizationRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetRoles gets the roles property value. The list of organization roles available to the organization.
func (m *ItemOrganizationRolesGetResponse) GetRoles()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable) {
    return m.roles
}
// GetTotalCount gets the total_count property value. The total number of organization roles available to the organization.
func (m *ItemOrganizationRolesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemOrganizationRolesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemOrganizationRolesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRoles sets the roles property value. The list of organization roles available to the organization.
func (m *ItemOrganizationRolesGetResponse) SetRoles(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable)() {
    m.roles = value
}
// SetTotalCount sets the total_count property value. The total number of organization roles available to the organization.
func (m *ItemOrganizationRolesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemOrganizationRolesGetResponseable 
type ItemOrganizationRolesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoles()([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable)
    GetTotalCount()(*int32)
    SetRoles(value []i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.OrganizationRoleable)()
    SetTotalCount(value *int32)()
}
