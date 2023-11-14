package repos

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunnersGenerateJitconfigPostResponse 
type ItemItemActionsRunnersGenerateJitconfigPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The base64 encoded runner configuration.
    encoded_jit_config *string
    // A self hosted runner
    runner i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable
}
// NewItemItemActionsRunnersGenerateJitconfigPostResponse instantiates a new ItemItemActionsRunnersGenerateJitconfigPostResponse and sets the default values.
func NewItemItemActionsRunnersGenerateJitconfigPostResponse()(*ItemItemActionsRunnersGenerateJitconfigPostResponse) {
    m := &ItemItemActionsRunnersGenerateJitconfigPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunnersGenerateJitconfigPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunnersGenerateJitconfigPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEncodedJitConfig gets the encoded_jit_config property value. The base64 encoded runner configuration.
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) GetEncodedJitConfig()(*string) {
    return m.encoded_jit_config
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encoded_jit_config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncodedJitConfig(val)
        }
        return nil
    }
    res["runner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRunnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunner(val.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable))
        }
        return nil
    }
    return res
}
// GetRunner gets the runner property value. A self hosted runner
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) GetRunner()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable) {
    return m.runner
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("encoded_jit_config", m.GetEncodedJitConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("runner", m.GetRunner())
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
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEncodedJitConfig sets the encoded_jit_config property value. The base64 encoded runner configuration.
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) SetEncodedJitConfig(value *string)() {
    m.encoded_jit_config = value
}
// SetRunner sets the runner property value. A self hosted runner
func (m *ItemItemActionsRunnersGenerateJitconfigPostResponse) SetRunner(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable)() {
    m.runner = value
}
// ItemItemActionsRunnersGenerateJitconfigPostResponseable 
type ItemItemActionsRunnersGenerateJitconfigPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncodedJitConfig()(*string)
    GetRunner()(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable)
    SetEncodedJitConfig(value *string)()
    SetRunner(value i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Runnerable)()
}
