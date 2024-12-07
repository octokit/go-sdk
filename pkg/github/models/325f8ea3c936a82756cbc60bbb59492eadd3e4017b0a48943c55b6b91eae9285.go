package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSecurityConfiguration_code_scanning_default_setup_options feature options for code scanning default setup
type CodeSecurityConfiguration_code_scanning_default_setup_options struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The label of the runner to use for code scanning when runner_type is 'labeled'.
    runner_label *string
    // Whether to use labeled runners or standard GitHub runners.
    runner_type *CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type
}
// NewCodeSecurityConfiguration_code_scanning_default_setup_options instantiates a new CodeSecurityConfiguration_code_scanning_default_setup_options and sets the default values.
func NewCodeSecurityConfiguration_code_scanning_default_setup_options()(*CodeSecurityConfiguration_code_scanning_default_setup_options) {
    m := &CodeSecurityConfiguration_code_scanning_default_setup_options{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityConfiguration_code_scanning_default_setup_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityConfiguration_code_scanning_default_setup_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityConfiguration_code_scanning_default_setup_options(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["runner_label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerLabel(val)
        }
        return nil
    }
    res["runner_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_code_scanning_default_setup_options_runner_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerType(val.(*CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type))
        }
        return nil
    }
    return res
}
// GetRunnerLabel gets the runner_label property value. The label of the runner to use for code scanning when runner_type is 'labeled'.
// returns a *string when successful
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) GetRunnerLabel()(*string) {
    return m.runner_label
}
// GetRunnerType gets the runner_type property value. Whether to use labeled runners or standard GitHub runners.
// returns a *CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type when successful
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) GetRunnerType()(*CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type) {
    return m.runner_type
}
// Serialize serializes information the current object
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("runner_label", m.GetRunnerLabel())
        if err != nil {
            return err
        }
    }
    if m.GetRunnerType() != nil {
        cast := (*m.GetRunnerType()).String()
        err := writer.WriteStringValue("runner_type", &cast)
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
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRunnerLabel sets the runner_label property value. The label of the runner to use for code scanning when runner_type is 'labeled'.
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) SetRunnerLabel(value *string)() {
    m.runner_label = value
}
// SetRunnerType sets the runner_type property value. Whether to use labeled runners or standard GitHub runners.
func (m *CodeSecurityConfiguration_code_scanning_default_setup_options) SetRunnerType(value *CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type)() {
    m.runner_type = value
}
type CodeSecurityConfiguration_code_scanning_default_setup_optionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRunnerLabel()(*string)
    GetRunnerType()(*CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type)
    SetRunnerLabel(value *string)()
    SetRunnerType(value *CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type)()
}
