package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleMaxFilePathLength prevent commits that include file paths that exceed a specified character limit from being pushed to the commit graph.
type RepositoryRuleMaxFilePathLength struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleMaxFilePathLength_parametersable
    // The type property
    typeEscaped *RepositoryRuleMaxFilePathLength_type
}
// NewRepositoryRuleMaxFilePathLength instantiates a new RepositoryRuleMaxFilePathLength and sets the default values.
func NewRepositoryRuleMaxFilePathLength()(*RepositoryRuleMaxFilePathLength) {
    m := &RepositoryRuleMaxFilePathLength{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleMaxFilePathLengthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleMaxFilePathLengthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleMaxFilePathLength(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleMaxFilePathLength) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleMaxFilePathLength) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleMaxFilePathLength_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleMaxFilePathLength_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleMaxFilePathLength_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleMaxFilePathLength_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleMaxFilePathLength_parametersable when successful
func (m *RepositoryRuleMaxFilePathLength) GetParameters()(RepositoryRuleMaxFilePathLength_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleMaxFilePathLength_type when successful
func (m *RepositoryRuleMaxFilePathLength) GetTypeEscaped()(*RepositoryRuleMaxFilePathLength_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleMaxFilePathLength) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleMaxFilePathLength) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleMaxFilePathLength) SetParameters(value RepositoryRuleMaxFilePathLength_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleMaxFilePathLength) SetTypeEscaped(value *RepositoryRuleMaxFilePathLength_type)() {
    m.typeEscaped = value
}
type RepositoryRuleMaxFilePathLengthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleMaxFilePathLength_parametersable)
    GetTypeEscaped()(*RepositoryRuleMaxFilePathLength_type)
    SetParameters(value RepositoryRuleMaxFilePathLength_parametersable)()
    SetTypeEscaped(value *RepositoryRuleMaxFilePathLength_type)()
}
