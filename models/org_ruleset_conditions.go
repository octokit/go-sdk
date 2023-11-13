package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgRulesetConditions conditions for an organization ruleset. The conditions object should contain both `repository_name` and `ref_name` properties or both `repository_id` and `ref_name` properties.
type OrgRulesetConditions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
    orgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs RepositoryRulesetConditionsForRepositoryIDsable
    // Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
    orgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames RepositoryRulesetConditionsForRepositoryNamesable
    // Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
    repositoryRulesetConditionsForRepositoryIDs RepositoryRulesetConditionsForRepositoryIDsable
    // Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
    repositoryRulesetConditionsForRepositoryNames RepositoryRulesetConditionsForRepositoryNamesable
}
// NewOrgRulesetConditions instantiates a new orgRulesetConditions and sets the default values.
func NewOrgRulesetConditions()(*OrgRulesetConditions) {
    m := &OrgRulesetConditions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgRulesetConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrgRulesetConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOrgRulesetConditions()
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
    return result, nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrgRulesetConditions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrgRulesetConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *OrgRulesetConditions) GetIsComposedType()(bool) {
    return true
}
// GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs gets the RepositoryRulesetConditionsForRepositoryIDs property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
func (m *OrgRulesetConditions) GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs()(RepositoryRulesetConditionsForRepositoryIDsable) {
    return m.orgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs
}
// GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames gets the RepositoryRulesetConditionsForRepositoryNames property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
func (m *OrgRulesetConditions) GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames()(RepositoryRulesetConditionsForRepositoryNamesable) {
    return m.orgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames
}
// GetRepositoryRulesetConditionsForRepositoryIDs gets the RepositoryRulesetConditionsForRepositoryIDs property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
func (m *OrgRulesetConditions) GetRepositoryRulesetConditionsForRepositoryIDs()(RepositoryRulesetConditionsForRepositoryIDsable) {
    return m.repositoryRulesetConditionsForRepositoryIDs
}
// GetRepositoryRulesetConditionsForRepositoryNames gets the RepositoryRulesetConditionsForRepositoryNames property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
func (m *OrgRulesetConditions) GetRepositoryRulesetConditionsForRepositoryNames()(RepositoryRulesetConditionsForRepositoryNamesable) {
    return m.repositoryRulesetConditionsForRepositoryNames
}
// Serialize serializes information the current object
func (m *OrgRulesetConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs() != nil {
        err := writer.WriteObjectValue("", m.GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs())
        if err != nil {
            return err
        }
    } else if m.GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames() != nil {
        err := writer.WriteObjectValue("", m.GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRulesetConditionsForRepositoryIDs() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRulesetConditionsForRepositoryIDs())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRulesetConditionsForRepositoryNames() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRulesetConditionsForRepositoryNames())
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
func (m *OrgRulesetConditions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs sets the RepositoryRulesetConditionsForRepositoryIDs property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
func (m *OrgRulesetConditions) SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs(value RepositoryRulesetConditionsForRepositoryIDsable)() {
    m.orgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs = value
}
// SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames sets the RepositoryRulesetConditionsForRepositoryNames property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
func (m *OrgRulesetConditions) SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames(value RepositoryRulesetConditionsForRepositoryNamesable)() {
    m.orgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames = value
}
// SetRepositoryRulesetConditionsForRepositoryIDs sets the RepositoryRulesetConditionsForRepositoryIDs property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryIDs
func (m *OrgRulesetConditions) SetRepositoryRulesetConditionsForRepositoryIDs(value RepositoryRulesetConditionsForRepositoryIDsable)() {
    m.repositoryRulesetConditionsForRepositoryIDs = value
}
// SetRepositoryRulesetConditionsForRepositoryNames sets the RepositoryRulesetConditionsForRepositoryNames property value. Composed type representation for type RepositoryRulesetConditionsForRepositoryNames
func (m *OrgRulesetConditions) SetRepositoryRulesetConditionsForRepositoryNames(value RepositoryRulesetConditionsForRepositoryNamesable)() {
    m.repositoryRulesetConditionsForRepositoryNames = value
}
// OrgRulesetConditionsable 
type OrgRulesetConditionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs()(RepositoryRulesetConditionsForRepositoryIDsable)
    GetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames()(RepositoryRulesetConditionsForRepositoryNamesable)
    GetRepositoryRulesetConditionsForRepositoryIDs()(RepositoryRulesetConditionsForRepositoryIDsable)
    GetRepositoryRulesetConditionsForRepositoryNames()(RepositoryRulesetConditionsForRepositoryNamesable)
    SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryIDs(value RepositoryRulesetConditionsForRepositoryIDsable)()
    SetOrgRulesetConditionsRepositoryRulesetConditionsForRepositoryNames(value RepositoryRulesetConditionsForRepositoryNamesable)()
    SetRepositoryRulesetConditionsForRepositoryIDs(value RepositoryRulesetConditionsForRepositoryIDsable)()
    SetRepositoryRulesetConditionsForRepositoryNames(value RepositoryRulesetConditionsForRepositoryNamesable)()
}
