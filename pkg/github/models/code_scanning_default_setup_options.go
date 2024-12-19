package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningDefaultSetupOptions feature options for code scanning default setup
type CodeScanningDefaultSetupOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
    runner_label *string
    // Whether to use labeled runners or standard GitHub runners.
    runner_type *CodeScanningDefaultSetupOptions_runner_type
}
// NewCodeScanningDefaultSetupOptions instantiates a new CodeScanningDefaultSetupOptions and sets the default values.
func NewCodeScanningDefaultSetupOptions()(*CodeScanningDefaultSetupOptions) {
    m := &CodeScanningDefaultSetupOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningDefaultSetupOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningDefaultSetupOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningDefaultSetupOptions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningDefaultSetupOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningDefaultSetupOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetupOptions_runner_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunnerType(val.(*CodeScanningDefaultSetupOptions_runner_type))
        }
        return nil
    }
    return res
}
// GetRunnerLabel gets the runner_label property value. The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
// returns a *string when successful
func (m *CodeScanningDefaultSetupOptions) GetRunnerLabel()(*string) {
    return m.runner_label
}
// GetRunnerType gets the runner_type property value. Whether to use labeled runners or standard GitHub runners.
// returns a *CodeScanningDefaultSetupOptions_runner_type when successful
func (m *CodeScanningDefaultSetupOptions) GetRunnerType()(*CodeScanningDefaultSetupOptions_runner_type) {
    return m.runner_type
}
// Serialize serializes information the current object
func (m *CodeScanningDefaultSetupOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CodeScanningDefaultSetupOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRunnerLabel sets the runner_label property value. The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
func (m *CodeScanningDefaultSetupOptions) SetRunnerLabel(value *string)() {
    m.runner_label = value
}
// SetRunnerType sets the runner_type property value. Whether to use labeled runners or standard GitHub runners.
func (m *CodeScanningDefaultSetupOptions) SetRunnerType(value *CodeScanningDefaultSetupOptions_runner_type)() {
    m.runner_type = value
}
type CodeScanningDefaultSetupOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRunnerLabel()(*string)
    GetRunnerType()(*CodeScanningDefaultSetupOptions_runner_type)
    SetRunnerLabel(value *string)()
    SetRunnerType(value *CodeScanningDefaultSetupOptions_runner_type)()
}
