package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotIdeChat_editors_models struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The training date for the custom model.
    custom_model_training_date *string
    // Indicates whether a model is custom or default.
    is_custom_model *bool
    // Name of the language used for Copilot code completion suggestions, for the given editor.
    name *string
    // The number of times users copied a code suggestion from Copilot Chat using the keyboard, or the 'Copy' UI element, for the given editor.
    total_chat_copy_events *int32
    // The number of times users accepted a code suggestion from Copilot Chat using the 'Insert Code' UI element, for the given editor.
    total_chat_insertion_events *int32
    // The total number of chats initiated by users in the given editor and model.
    total_chats *int32
    // The number of users who prompted Copilot Chat in the given editor and model.
    total_engaged_users *int32
}
// NewCopilotIdeChat_editors_models instantiates a new CopilotIdeChat_editors_models and sets the default values.
func NewCopilotIdeChat_editors_models()(*CopilotIdeChat_editors_models) {
    m := &CopilotIdeChat_editors_models{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeChat_editors_modelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeChat_editors_modelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeChat_editors_models(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeChat_editors_models) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomModelTrainingDate gets the custom_model_training_date property value. The training date for the custom model.
// returns a *string when successful
func (m *CopilotIdeChat_editors_models) GetCustomModelTrainingDate()(*string) {
    return m.custom_model_training_date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeChat_editors_models) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["total_chat_copy_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalChatCopyEvents(val)
        }
        return nil
    }
    res["total_chat_insertion_events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalChatInsertionEvents(val)
        }
        return nil
    }
    res["total_chats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalChats(val)
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
func (m *CopilotIdeChat_editors_models) GetIsCustomModel()(*bool) {
    return m.is_custom_model
}
// GetName gets the name property value. Name of the language used for Copilot code completion suggestions, for the given editor.
// returns a *string when successful
func (m *CopilotIdeChat_editors_models) GetName()(*string) {
    return m.name
}
// GetTotalChatCopyEvents gets the total_chat_copy_events property value. The number of times users copied a code suggestion from Copilot Chat using the keyboard, or the 'Copy' UI element, for the given editor.
// returns a *int32 when successful
func (m *CopilotIdeChat_editors_models) GetTotalChatCopyEvents()(*int32) {
    return m.total_chat_copy_events
}
// GetTotalChatInsertionEvents gets the total_chat_insertion_events property value. The number of times users accepted a code suggestion from Copilot Chat using the 'Insert Code' UI element, for the given editor.
// returns a *int32 when successful
func (m *CopilotIdeChat_editors_models) GetTotalChatInsertionEvents()(*int32) {
    return m.total_chat_insertion_events
}
// GetTotalChats gets the total_chats property value. The total number of chats initiated by users in the given editor and model.
// returns a *int32 when successful
func (m *CopilotIdeChat_editors_models) GetTotalChats()(*int32) {
    return m.total_chats
}
// GetTotalEngagedUsers gets the total_engaged_users property value. The number of users who prompted Copilot Chat in the given editor and model.
// returns a *int32 when successful
func (m *CopilotIdeChat_editors_models) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeChat_editors_models) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_chats", m.GetTotalChats())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_chat_copy_events", m.GetTotalChatCopyEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_chat_insertion_events", m.GetTotalChatInsertionEvents())
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
func (m *CopilotIdeChat_editors_models) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomModelTrainingDate sets the custom_model_training_date property value. The training date for the custom model.
func (m *CopilotIdeChat_editors_models) SetCustomModelTrainingDate(value *string)() {
    m.custom_model_training_date = value
}
// SetIsCustomModel sets the is_custom_model property value. Indicates whether a model is custom or default.
func (m *CopilotIdeChat_editors_models) SetIsCustomModel(value *bool)() {
    m.is_custom_model = value
}
// SetName sets the name property value. Name of the language used for Copilot code completion suggestions, for the given editor.
func (m *CopilotIdeChat_editors_models) SetName(value *string)() {
    m.name = value
}
// SetTotalChatCopyEvents sets the total_chat_copy_events property value. The number of times users copied a code suggestion from Copilot Chat using the keyboard, or the 'Copy' UI element, for the given editor.
func (m *CopilotIdeChat_editors_models) SetTotalChatCopyEvents(value *int32)() {
    m.total_chat_copy_events = value
}
// SetTotalChatInsertionEvents sets the total_chat_insertion_events property value. The number of times users accepted a code suggestion from Copilot Chat using the 'Insert Code' UI element, for the given editor.
func (m *CopilotIdeChat_editors_models) SetTotalChatInsertionEvents(value *int32)() {
    m.total_chat_insertion_events = value
}
// SetTotalChats sets the total_chats property value. The total number of chats initiated by users in the given editor and model.
func (m *CopilotIdeChat_editors_models) SetTotalChats(value *int32)() {
    m.total_chats = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. The number of users who prompted Copilot Chat in the given editor and model.
func (m *CopilotIdeChat_editors_models) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeChat_editors_modelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomModelTrainingDate()(*string)
    GetIsCustomModel()(*bool)
    GetName()(*string)
    GetTotalChatCopyEvents()(*int32)
    GetTotalChatInsertionEvents()(*int32)
    GetTotalChats()(*int32)
    GetTotalEngagedUsers()(*int32)
    SetCustomModelTrainingDate(value *string)()
    SetIsCustomModel(value *bool)()
    SetName(value *string)()
    SetTotalChatCopyEvents(value *int32)()
    SetTotalChatInsertionEvents(value *int32)()
    SetTotalChats(value *int32)()
    SetTotalEngagedUsers(value *int32)()
}
