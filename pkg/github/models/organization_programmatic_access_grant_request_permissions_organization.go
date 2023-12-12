package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationProgrammaticAccessGrantRequest_permissions_organization 
type OrganizationProgrammaticAccessGrantRequest_permissions_organization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewOrganizationProgrammaticAccessGrantRequest_permissions_organization instantiates a new organizationProgrammaticAccessGrantRequest_permissions_organization and sets the default values.
func NewOrganizationProgrammaticAccessGrantRequest_permissions_organization()(*OrganizationProgrammaticAccessGrantRequest_permissions_organization) {
    m := &OrganizationProgrammaticAccessGrantRequest_permissions_organization{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationProgrammaticAccessGrantRequest_permissions_organizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationProgrammaticAccessGrantRequest_permissions_organizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationProgrammaticAccessGrantRequest_permissions_organization(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationProgrammaticAccessGrantRequest_permissions_organization) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationProgrammaticAccessGrantRequest_permissions_organization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *OrganizationProgrammaticAccessGrantRequest_permissions_organization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationProgrammaticAccessGrantRequest_permissions_organization) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// OrganizationProgrammaticAccessGrantRequest_permissions_organizationable 
type OrganizationProgrammaticAccessGrantRequest_permissions_organizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
