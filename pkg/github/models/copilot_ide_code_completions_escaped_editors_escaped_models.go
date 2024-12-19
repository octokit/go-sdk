package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotIdeCodeCompletions_editors_models struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The training date for the custom model.
    custom_model_training_date *string
    // Indicates whether a model is custom or default.
    is_custom_model *bool
    // Code completion metrics for active languages, for the given editor.
    languages []CopilotIdeCodeCompletions_editors_models_languagesable
    // Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
    name *string
    // Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language and model. Includes both full and partial acceptances.
    total_engaged_users *int32
}
// NewCopilotIdeCodeCompletions_editors_models instantiates a new CopilotIdeCodeCompletions_editors_models and sets the default values.
func NewCopilotIdeCodeCompletions_editors_models()(*CopilotIdeCodeCompletions_editors_models) {
    m := &CopilotIdeCodeCompletions_editors_models{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeCodeCompletions_editors_modelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeCodeCompletions_editors_modelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeCodeCompletions_editors_models(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomModelTrainingDate gets the custom_model_training_date property value. The training date for the custom model.
// returns a *string when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetCustomModelTrainingDate()(*string) {
    return m.custom_model_training_date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_model_training_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomModelTrainingDate(val)
        }
        return nil
    }
    res["is_custom_model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomModel(val)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotIdeCodeCompletions_editors_models_languagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotIdeCodeCompletions_editors_models_languagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotIdeCodeCompletions_editors_models_languagesable)
                }
            }
            m.SetLanguages(res)
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
// GetIsCustomModel gets the is_custom_model property value. Indicates whether a model is custom or default.
// returns a *bool when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetIsCustomModel()(*bool) {
    return m.is_custom_model
}
// GetLanguages gets the languages property value. Code completion metrics for active languages, for the given editor.
// returns a []CopilotIdeCodeCompletions_editors_models_languagesable when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetLanguages()([]CopilotIdeCodeCompletions_editors_models_languagesable) {
    return m.languages
}
// GetName gets the name property value. Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
// returns a *string when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetName()(*string) {
    return m.name
}
// GetTotalEngagedUsers gets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language and model. Includes both full and partial acceptances.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeCodeCompletions_editors_models) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("custom_model_training_date", m.GetCustomModelTrainingDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_custom_model", m.GetIsCustomModel())
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("languages", cast)
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
func (m *CopilotIdeCodeCompletions_editors_models) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomModelTrainingDate sets the custom_model_training_date property value. The training date for the custom model.
func (m *CopilotIdeCodeCompletions_editors_models) SetCustomModelTrainingDate(value *string)() {
    m.custom_model_training_date = value
}
// SetIsCustomModel sets the is_custom_model property value. Indicates whether a model is custom or default.
func (m *CopilotIdeCodeCompletions_editors_models) SetIsCustomModel(value *bool)() {
    m.is_custom_model = value
}
// SetLanguages sets the languages property value. Code completion metrics for active languages, for the given editor.
func (m *CopilotIdeCodeCompletions_editors_models) SetLanguages(value []CopilotIdeCodeCompletions_editors_models_languagesable)() {
    m.languages = value
}
// SetName sets the name property value. Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
func (m *CopilotIdeCodeCompletions_editors_models) SetName(value *string)() {
    m.name = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language and model. Includes both full and partial acceptances.
func (m *CopilotIdeCodeCompletions_editors_models) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeCodeCompletions_editors_modelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomModelTrainingDate()(*string)
    GetIsCustomModel()(*bool)
    GetLanguages()([]CopilotIdeCodeCompletions_editors_models_languagesable)
    GetName()(*string)
    GetTotalEngagedUsers()(*int32)
    SetCustomModelTrainingDate(value *string)()
    SetIsCustomModel(value *bool)()
    SetLanguages(value []CopilotIdeCodeCompletions_editors_models_languagesable)()
    SetName(value *string)()
    SetTotalEngagedUsers(value *int32)()
}
