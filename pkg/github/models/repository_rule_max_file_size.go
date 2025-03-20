package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleMaxFileSize prevent commits that exceed a specified file size limit from being pushed to the commit graph.
type RepositoryRuleMaxFileSize struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleMaxFileSize_parametersable
    // The type property
    typeEscaped *RepositoryRuleMaxFileSize_type
}
// NewRepositoryRuleMaxFileSize instantiates a new RepositoryRuleMaxFileSize and sets the default values.
func NewRepositoryRuleMaxFileSize()(*RepositoryRuleMaxFileSize) {
    m := &RepositoryRuleMaxFileSize{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMaxFileSizeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMaxFileSizeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMaxFileSize(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMaxFileSize) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMaxFileSize) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleMaxFileSize_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleMaxFileSize_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMaxFileSize_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleMaxFileSize_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleMaxFileSize_parametersable when successful
func (m *RepositoryRuleMaxFileSize) GetParameters()(RepositoryRuleMaxFileSize_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleMaxFileSize_type when successful
func (m *RepositoryRuleMaxFileSize) GetTypeEscaped()(*RepositoryRuleMaxFileSize_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleMaxFileSize) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parameters", m.GetParameters())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *RepositoryRuleMaxFileSize) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleMaxFileSize) SetParameters(value RepositoryRuleMaxFileSize_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleMaxFileSize) SetTypeEscaped(value *RepositoryRuleMaxFileSize_type)() {
    m.typeEscaped = value
}
type RepositoryRuleMaxFileSizeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleMaxFileSize_parametersable)
    GetTypeEscaped()(*RepositoryRuleMaxFileSize_type)
    SetParameters(value RepositoryRuleMaxFileSize_parametersable)()
    SetTypeEscaped(value *RepositoryRuleMaxFileSize_type)()
}
