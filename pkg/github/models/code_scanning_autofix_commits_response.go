package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAutofixCommitsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // SHA of commit with autofix.
    sha *string
    // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
    target_ref *string
}
// NewCodeScanningAutofixCommitsResponse instantiates a new CodeScanningAutofixCommitsResponse and sets the default values.
func NewCodeScanningAutofixCommitsResponse()(*CodeScanningAutofixCommitsResponse) {
    m := &CodeScanningAutofixCommitsResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAutofixCommitsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAutofixCommitsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAutofixCommitsResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAutofixCommitsResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAutofixCommitsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
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
// GetSha gets the sha property value. SHA of commit with autofix.
// returns a *string when successful
func (m *CodeScanningAutofixCommitsResponse) GetSha()(*string) {
    return m.sha
}
// GetTargetRef gets the target_ref property value. The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
// returns a *string when successful
func (m *CodeScanningAutofixCommitsResponse) GetTargetRef()(*string) {
    return m.target_ref
}
// Serialize serializes information the current object
func (m *CodeScanningAutofixCommitsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sha", m.GetSha())
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
func (m *CodeScanningAutofixCommitsResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSha sets the sha property value. SHA of commit with autofix.
func (m *CodeScanningAutofixCommitsResponse) SetSha(value *string)() {
    m.sha = value
}
// SetTargetRef sets the target_ref property value. The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
func (m *CodeScanningAutofixCommitsResponse) SetTargetRef(value *string)() {
    m.target_ref = value
}
type CodeScanningAutofixCommitsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSha()(*string)
    GetTargetRef()(*string)
    SetSha(value *string)()
    SetTargetRef(value *string)()
}
