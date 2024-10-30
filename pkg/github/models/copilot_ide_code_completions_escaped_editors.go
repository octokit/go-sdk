package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotIdeCodeCompletions_editors copilot code completion metrics for active editors.
type CopilotIdeCodeCompletions_editors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of model metrics for custom models and the default model.
    models []CopilotIdeCodeCompletions_editors_modelsable
    // Name of the given editor.
    name *string
    // Number of users who accepted at least one Copilot code completion suggestion for the given editor. Includes both full and partial acceptances.
    total_engaged_users *int32
}
// NewCopilotIdeCodeCompletions_editors instantiates a new CopilotIdeCodeCompletions_editors and sets the default values.
func NewCopilotIdeCodeCompletions_editors()(*CopilotIdeCodeCompletions_editors) {
    m := &CopilotIdeCodeCompletions_editors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeCodeCompletions_editorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeCodeCompletions_editorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeCodeCompletions_editors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeCodeCompletions_editors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeCodeCompletions_editors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["models"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotIdeCodeCompletions_editors_modelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotIdeCodeCompletions_editors_modelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotIdeCodeCompletions_editors_modelsable)
                }
            }
            m.SetModels(res)
        }
        return nil
    }
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
// GetModels gets the models property value. List of model metrics for custom models and the default model.
// returns a []CopilotIdeCodeCompletions_editors_modelsable when successful
func (m *CopilotIdeCodeCompletions_editors) GetModels()([]CopilotIdeCodeCompletions_editors_modelsable) {
    return m.models
}
// GetName gets the name property value. Name of the given editor.
// returns a *string when successful
func (m *CopilotIdeCodeCompletions_editors) GetName()(*string) {
    return m.name
}
// GetTotalEngagedUsers gets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor. Includes both full and partial acceptances.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeCodeCompletions_editors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *CopilotIdeCodeCompletions_editors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetModels sets the models property value. List of model metrics for custom models and the default model.
func (m *CopilotIdeCodeCompletions_editors) SetModels(value []CopilotIdeCodeCompletions_editors_modelsable)() {
    m.models = value
}
// SetName sets the name property value. Name of the given editor.
func (m *CopilotIdeCodeCompletions_editors) SetName(value *string)() {
    m.name = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor. Includes both full and partial acceptances.
func (m *CopilotIdeCodeCompletions_editors) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeCodeCompletions_editorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetModels()([]CopilotIdeCodeCompletions_editors_modelsable)
    GetName()(*string)
    GetTotalEngagedUsers()(*int32)
    SetModels(value []CopilotIdeCodeCompletions_editors_modelsable)()
    SetName(value *string)()
    SetTotalEngagedUsers(value *int32)()
}
