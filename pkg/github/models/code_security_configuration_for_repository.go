package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSecurityConfigurationForRepository code security configuration associated with a repository and attachment status
type CodeSecurityConfigurationForRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A code security configuration
    configuration CodeSecurityConfigurationable
    // The attachment status of the code security configuration on the repository.
    status *CodeSecurityConfigurationForRepository_status
}
// NewCodeSecurityConfigurationForRepository instantiates a new CodeSecurityConfigurationForRepository and sets the default values.
func NewCodeSecurityConfigurationForRepository()(*CodeSecurityConfigurationForRepository) {
    m := &CodeSecurityConfigurationForRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityConfigurationForRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityConfigurationForRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityConfigurationForRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityConfigurationForRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfiguration gets the configuration property value. A code security configuration
// returns a CodeSecurityConfigurationable when successful
func (m *CodeSecurityConfigurationForRepository) GetConfiguration()(CodeSecurityConfigurationable) {
    return m.configuration
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityConfigurationForRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeSecurityConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(CodeSecurityConfigurationable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfigurationForRepository_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CodeSecurityConfigurationForRepository_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The attachment status of the code security configuration on the repository.
// returns a *CodeSecurityConfigurationForRepository_status when successful
func (m *CodeSecurityConfigurationForRepository) GetStatus()(*CodeSecurityConfigurationForRepository_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *CodeSecurityConfigurationForRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CodeSecurityConfigurationForRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfiguration sets the configuration property value. A code security configuration
func (m *CodeSecurityConfigurationForRepository) SetConfiguration(value CodeSecurityConfigurationable)() {
    m.configuration = value
}
// SetStatus sets the status property value. The attachment status of the code security configuration on the repository.
func (m *CodeSecurityConfigurationForRepository) SetStatus(value *CodeSecurityConfigurationForRepository_status)() {
    m.status = value
}
type CodeSecurityConfigurationForRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(CodeSecurityConfigurationable)
    GetStatus()(*CodeSecurityConfigurationForRepository_status)
    SetConfiguration(value CodeSecurityConfigurationable)()
    SetStatus(value *CodeSecurityConfigurationForRepository_status)()
}
