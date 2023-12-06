package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody 
type ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // **Required when the `state` is `resolved`.** The reason for resolving the alert.
    resolution *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution
    // An optional comment when closing an alert. Cannot be updated or deleted. Must be `null` when changing `state` to `open`.
    resolution_comment *string
    // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
    state *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState
}
// NewItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody instantiates a new ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody and sets the default values.
func NewItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody()(*ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) {
    m := &ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ParseSecretScanningAlertResolution)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolution(val.(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution))
        }
        return nil
    }
    res["resolution_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolutionComment(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ParseSecretScanningAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState))
        }
        return nil
    }
    return res
}
// GetResolution gets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) GetResolution()(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution) {
    return m.resolution
}
// GetResolutionComment gets the resolution_comment property value. An optional comment when closing an alert. Cannot be updated or deleted. Must be `null` when changing `state` to `open`.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) GetResolutionComment()(*string) {
    return m.resolution_comment
}
// GetState gets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) GetState()(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetResolution() != nil {
        cast := (*m.GetResolution()).String()
        err := writer.WriteStringValue("resolution", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resolution_comment", m.GetResolutionComment())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetResolution sets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) SetResolution(value *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution)() {
    m.resolution = value
}
// SetResolutionComment sets the resolution_comment property value. An optional comment when closing an alert. Cannot be updated or deleted. Must be `null` when changing `state` to `open`.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) SetResolutionComment(value *string)() {
    m.resolution_comment = value
}
// SetState sets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
func (m *ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBody) SetState(value *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState)() {
    m.state = value
}
// ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyable 
type ItemItemSecretScanningAlertsItemWithAlert_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResolution()(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution)
    GetResolutionComment()(*string)
    GetState()(*i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState)
    SetResolution(value *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertResolution)()
    SetResolutionComment(value *string)()
    SetState(value *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SecretScanningAlertState)()
}
