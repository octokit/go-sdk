package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphSpdxSbom_sbom_relationships struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The SPDX identifier of the package that is the target of the relationship.
    relatedSpdxElement *string
    // The type of relationship between the two SPDX elements.
    relationshipType *string
    // The SPDX identifier of the package that is the source of the relationship.
    spdxElementId *string
}
// NewDependencyGraphSpdxSbom_sbom_relationships instantiates a new DependencyGraphSpdxSbom_sbom_relationships and sets the default values.
func NewDependencyGraphSpdxSbom_sbom_relationships()(*DependencyGraphSpdxSbom_sbom_relationships) {
    m := &DependencyGraphSpdxSbom_sbom_relationships{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbom_sbom_relationshipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbom_sbom_relationshipsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom_sbom_relationships(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom_sbom_relationships) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom_sbom_relationships) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["relatedSpdxElement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelatedSpdxElement(val)
        }
        return nil
    }
    res["relationshipType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationshipType(val)
        }
        return nil
    }
    res["spdxElementId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpdxElementId(val)
        }
        return nil
    }
    return res
}
// GetRelatedSpdxElement gets the relatedSpdxElement property value. The SPDX identifier of the package that is the target of the relationship.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_relationships) GetRelatedSpdxElement()(*string) {
    return m.relatedSpdxElement
}
// GetRelationshipType gets the relationshipType property value. The type of relationship between the two SPDX elements.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_relationships) GetRelationshipType()(*string) {
    return m.relationshipType
}
// GetSpdxElementId gets the spdxElementId property value. The SPDX identifier of the package that is the source of the relationship.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_relationships) GetSpdxElementId()(*string) {
    return m.spdxElementId
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom_sbom_relationships) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("relatedSpdxElement", m.GetRelatedSpdxElement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relationshipType", m.GetRelationshipType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("spdxElementId", m.GetSpdxElementId())
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
func (m *DependencyGraphSpdxSbom_sbom_relationships) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRelatedSpdxElement sets the relatedSpdxElement property value. The SPDX identifier of the package that is the target of the relationship.
func (m *DependencyGraphSpdxSbom_sbom_relationships) SetRelatedSpdxElement(value *string)() {
    m.relatedSpdxElement = value
}
// SetRelationshipType sets the relationshipType property value. The type of relationship between the two SPDX elements.
func (m *DependencyGraphSpdxSbom_sbom_relationships) SetRelationshipType(value *string)() {
    m.relationshipType = value
}
// SetSpdxElementId sets the spdxElementId property value. The SPDX identifier of the package that is the source of the relationship.
func (m *DependencyGraphSpdxSbom_sbom_relationships) SetSpdxElementId(value *string)() {
    m.spdxElementId = value
}
type DependencyGraphSpdxSbom_sbom_relationshipsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRelatedSpdxElement()(*string)
    GetRelationshipType()(*string)
    GetSpdxElementId()(*string)
    SetRelatedSpdxElement(value *string)()
    SetRelationshipType(value *string)()
    SetSpdxElementId(value *string)()
}
