package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAutofixCommits commit an autofix for a code scanning alert
type CodeScanningAutofixCommits struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Commit message to be used.
    message *string
    // The Git reference of target branch for the commit. Branch needs to already exist.  For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
    target_ref *string
}
// NewCodeScanningAutofixCommits instantiates a new CodeScanningAutofixCommits and sets the default values.
func NewCodeScanningAutofixCommits()(*CodeScanningAutofixCommits) {
    m := &CodeScanningAutofixCommits{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAutofixCommitsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAutofixCommitsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAutofixCommits(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAutofixCommits) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAutofixCommits) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["target_ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetRef(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Commit message to be used.
// returns a *string when successful
func (m *CodeScanningAutofixCommits) GetMessage()(*string) {
    return m.message
}
// GetTargetRef gets the target_ref property value. The Git reference of target branch for the commit. Branch needs to already exist.  For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
// returns a *string when successful
func (m *CodeScanningAutofixCommits) GetTargetRef()(*string) {
    return m.target_ref
}
// Serialize serializes information the current object
func (m *CodeScanningAutofixCommits) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_ref", m.GetTargetRef())
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
func (m *CodeScanningAutofixCommits) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. Commit message to be used.
func (m *CodeScanningAutofixCommits) SetMessage(value *string)() {
    m.message = value
}
// SetTargetRef sets the target_ref property value. The Git reference of target branch for the commit. Branch needs to already exist.  For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
func (m *CodeScanningAutofixCommits) SetTargetRef(value *string)() {
    m.target_ref = value
}
type CodeScanningAutofixCommitsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    GetTargetRef()(*string)
    SetMessage(value *string)()
    SetTargetRef(value *string)()
}
