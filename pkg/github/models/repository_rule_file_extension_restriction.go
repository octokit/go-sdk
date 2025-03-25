package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleFileExtensionRestriction prevent commits that include files with specified file extensions from being pushed to the commit graph.
type RepositoryRuleFileExtensionRestriction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters RepositoryRuleFileExtensionRestriction_parametersable
    // The type property
    typeEscaped *RepositoryRuleFileExtensionRestriction_type
}
// NewRepositoryRuleFileExtensionRestriction instantiates a new RepositoryRuleFileExtensionRestriction and sets the default values.
func NewRepositoryRuleFileExtensionRestriction()(*RepositoryRuleFileExtensionRestriction) {
    m := &RepositoryRuleFileExtensionRestriction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleFileExtensionRestrictionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleFileExtensionRestrictionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleFileExtensionRestriction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleFileExtensionRestriction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleFileExtensionRestriction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryRuleFileExtensionRestriction_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(RepositoryRuleFileExtensionRestriction_parametersable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleFileExtensionRestriction_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*RepositoryRuleFileExtensionRestriction_type))
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
// returns a RepositoryRuleFileExtensionRestriction_parametersable when successful
func (m *RepositoryRuleFileExtensionRestriction) GetParameters()(RepositoryRuleFileExtensionRestriction_parametersable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
// returns a *RepositoryRuleFileExtensionRestriction_type when successful
func (m *RepositoryRuleFileExtensionRestriction) GetTypeEscaped()(*RepositoryRuleFileExtensionRestriction_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryRuleFileExtensionRestriction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RepositoryRuleFileExtensionRestriction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *RepositoryRuleFileExtensionRestriction) SetParameters(value RepositoryRuleFileExtensionRestriction_parametersable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *RepositoryRuleFileExtensionRestriction) SetTypeEscaped(value *RepositoryRuleFileExtensionRestriction_type)() {
    m.typeEscaped = value
}
type RepositoryRuleFileExtensionRestrictionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()(RepositoryRuleFileExtensionRestriction_parametersable)
    GetTypeEscaped()(*RepositoryRuleFileExtensionRestriction_type)
    SetParameters(value RepositoryRuleFileExtensionRestriction_parametersable)()
    SetTypeEscaped(value *RepositoryRuleFileExtensionRestriction_type)()
}
