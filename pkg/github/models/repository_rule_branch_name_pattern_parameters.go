package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleBranchNamePattern_parameters 
type RepositoryRuleBranchNamePattern_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How this rule will appear to users.
    name *string
    // If true, the rule will fail if the pattern matches.
    negate *bool
    // The operator to use for matching.
    operator *RepositoryRuleBranchNamePattern_parameters_operator
    // The pattern to match with.
    pattern *string
}
// NewRepositoryRuleBranchNamePattern_parameters instantiates a new repositoryRuleBranchNamePattern_parameters and sets the default values.
func NewRepositoryRuleBranchNamePattern_parameters()(*RepositoryRuleBranchNamePattern_parameters) {
    m := &RepositoryRuleBranchNamePattern_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleBranchNamePattern_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryRuleBranchNamePattern_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleBranchNamePattern_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RepositoryRuleBranchNamePattern_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryRuleBranchNamePattern_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["negate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNegate(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleBranchNamePattern_parameters_operator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*RepositoryRuleBranchNamePattern_parameters_operator))
        }
        return nil
    }
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. How this rule will appear to users.
func (m *RepositoryRuleBranchNamePattern_parameters) GetName()(*string) {
    return m.name
}
// GetNegate gets the negate property value. If true, the rule will fail if the pattern matches.
func (m *RepositoryRuleBranchNamePattern_parameters) GetNegate()(*bool) {
    return m.negate
}
// GetOperator gets the operator property value. The operator to use for matching.
func (m *RepositoryRuleBranchNamePattern_parameters) GetOperator()(*RepositoryRuleBranchNamePattern_parameters_operator) {
    return m.operator
}
// GetPattern gets the pattern property value. The pattern to match with.
func (m *RepositoryRuleBranchNamePattern_parameters) GetPattern()(*string) {
    return m.pattern
}
// Serialize serializes information the current object
func (m *RepositoryRuleBranchNamePattern_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("negate", m.GetNegate())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pattern", m.GetPattern())
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
func (m *RepositoryRuleBranchNamePattern_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. How this rule will appear to users.
func (m *RepositoryRuleBranchNamePattern_parameters) SetName(value *string)() {
    m.name = value
}
// SetNegate sets the negate property value. If true, the rule will fail if the pattern matches.
func (m *RepositoryRuleBranchNamePattern_parameters) SetNegate(value *bool)() {
    m.negate = value
}
// SetOperator sets the operator property value. The operator to use for matching.
func (m *RepositoryRuleBranchNamePattern_parameters) SetOperator(value *RepositoryRuleBranchNamePattern_parameters_operator)() {
    m.operator = value
}
// SetPattern sets the pattern property value. The pattern to match with.
func (m *RepositoryRuleBranchNamePattern_parameters) SetPattern(value *string)() {
    m.pattern = value
}
// RepositoryRuleBranchNamePattern_parametersable 
type RepositoryRuleBranchNamePattern_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNegate()(*bool)
    GetOperator()(*RepositoryRuleBranchNamePattern_parameters_operator)
    GetPattern()(*string)
    SetName(value *string)()
    SetNegate(value *bool)()
    SetOperator(value *RepositoryRuleBranchNamePattern_parameters_operator)()
    SetPattern(value *string)()
}
