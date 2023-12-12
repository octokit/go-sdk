package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryAdvisory_cwes 
type RepositoryAdvisory_cwes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Common Weakness Enumeration (CWE) identifier.
    cwe_id *string
    // The name of the CWE.
    name *string
}
// NewRepositoryAdvisory_cwes instantiates a new repositoryAdvisory_cwes and sets the default values.
func NewRepositoryAdvisory_cwes()(*RepositoryAdvisory_cwes) {
    m := &RepositoryAdvisory_cwes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryAdvisory_cwesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryAdvisory_cwesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisory_cwes(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryAdvisory_cwes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCweId gets the cwe_id property value. The Common Weakness Enumeration (CWE) identifier.
func (m *RepositoryAdvisory_cwes) GetCweId()(*string) {
    return m.cwe_id
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryAdvisory_cwes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cwe_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCweId(val)
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
// GetName gets the name property value. The name of the CWE.
func (m *RepositoryAdvisory_cwes) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *RepositoryAdvisory_cwes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cwe_id", m.GetCweId())
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
func (m *RepositoryAdvisory_cwes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCweId sets the cwe_id property value. The Common Weakness Enumeration (CWE) identifier.
func (m *RepositoryAdvisory_cwes) SetCweId(value *string)() {
    m.cwe_id = value
}
// SetName sets the name property value. The name of the CWE.
func (m *RepositoryAdvisory_cwes) SetName(value *string)() {
    m.name = value
}
// RepositoryAdvisory_cwesable 
type RepositoryAdvisory_cwesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCweId()(*string)
    GetName()(*string)
    SetCweId(value *string)()
    SetName(value *string)()
}
