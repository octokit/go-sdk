package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotDotcomChat usage metrics for Copilot Chat in github.com
type CopilotDotcomChat struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of model metrics for a custom models and the default model.
    models []CopilotDotcomChat_modelsable
    // Total number of users who prompted Copilot Chat on github.com at least once.
    total_engaged_users *int32
}
// NewCopilotDotcomChat instantiates a new CopilotDotcomChat and sets the default values.
func NewCopilotDotcomChat()(*CopilotDotcomChat) {
    m := &CopilotDotcomChat{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotDotcomChatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotDotcomChatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotDotcomChat(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotDotcomChat) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotDotcomChat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["models"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotDotcomChat_modelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotDotcomChat_modelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotDotcomChat_modelsable)
                }
            }
            m.SetModels(res)
        }
        return nil
    }
    res["total_engaged_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEngagedUsers(val)
        }
        return nil
    }
    return res
}
// GetModels gets the models property value. List of model metrics for a custom models and the default model.
// returns a []CopilotDotcomChat_modelsable when successful
func (m *CopilotDotcomChat) GetModels()([]CopilotDotcomChat_modelsable) {
    return m.models
}
// GetTotalEngagedUsers gets the total_engaged_users property value. Total number of users who prompted Copilot Chat on github.com at least once.
// returns a *int32 when successful
func (m *CopilotDotcomChat) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotDotcomChat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetModels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModels()))
        for i, v := range m.GetModels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("models", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_engaged_users", m.GetTotalEngagedUsers())
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
func (m *CopilotDotcomChat) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetModels sets the models property value. List of model metrics for a custom models and the default model.
func (m *CopilotDotcomChat) SetModels(value []CopilotDotcomChat_modelsable)() {
    m.models = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Total number of users who prompted Copilot Chat on github.com at least once.
func (m *CopilotDotcomChat) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotDotcomChatable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetModels()([]CopilotDotcomChat_modelsable)
    GetTotalEngagedUsers()(*int32)
    SetModels(value []CopilotDotcomChat_modelsable)()
    SetTotalEngagedUsers(value *int32)()
}
