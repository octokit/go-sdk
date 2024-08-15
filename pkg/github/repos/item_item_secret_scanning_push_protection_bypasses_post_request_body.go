package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemSecretScanningPushProtectionBypassesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the push protection bypass placeholder. This value is returned on any push protected routes.
    placeholder_id *string
    // The reason for bypassing push protection.
    reason *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason
}
// NewItemItemSecretScanningPushProtectionBypassesPostRequestBody instantiates a new ItemItemSecretScanningPushProtectionBypassesPostRequestBody and sets the default values.
func NewItemItemSecretScanningPushProtectionBypassesPostRequestBody()(*ItemItemSecretScanningPushProtectionBypassesPostRequestBody) {
    m := &ItemItemSecretScanningPushProtectionBypassesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemSecretScanningPushProtectionBypassesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemSecretScanningPushProtectionBypassesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemSecretScanningPushProtectionBypassesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["placeholder_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceholderId(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseSecretScanningPushProtectionBypassReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason))
        }
        return nil
    }
    return res
}
// GetPlaceholderId gets the placeholder_id property value. The ID of the push protection bypass placeholder. This value is returned on any push protected routes.
// returns a *string when successful
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) GetPlaceholderId()(*string) {
    return m.placeholder_id
}
// GetReason gets the reason property value. The reason for bypassing push protection.
// returns a *SecretScanningPushProtectionBypassReason when successful
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) GetReason()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason) {
    return m.reason
}
// Serialize serializes information the current object
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("placeholder_id", m.GetPlaceholderId())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
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
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPlaceholderId sets the placeholder_id property value. The ID of the push protection bypass placeholder. This value is returned on any push protected routes.
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) SetPlaceholderId(value *string)() {
    m.placeholder_id = value
}
// SetReason sets the reason property value. The reason for bypassing push protection.
func (m *ItemItemSecretScanningPushProtectionBypassesPostRequestBody) SetReason(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason)() {
    m.reason = value
}
type ItemItemSecretScanningPushProtectionBypassesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPlaceholderId()(*string)
    GetReason()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason)
    SetPlaceholderId(value *string)()
    SetReason(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SecretScanningPushProtectionBypassReason)()
}
