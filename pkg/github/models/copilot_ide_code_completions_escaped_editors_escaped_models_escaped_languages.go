package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotIdeCodeCompletions_editors_models_languages usage metrics for a given language for the given editor for Copilot code completions.
type CopilotIdeCodeCompletions_editors_models_languages struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Name of the language used for Copilot code completion suggestions, for the given editor.
    name *string
    // The number of Copilot code suggestions accepted for the given editor, for the given language. Includes both full and partial acceptances.
    total_code_acceptances *int32
    // The number of lines of code accepted from Copilot code suggestions for the given editor, for the given language.
    total_code_lines_accepted *int32
    // The number of lines of code suggested by Copilot code completions for the given editor, for the given language.
    total_code_lines_suggested *int32
    // The number of Copilot code suggestions generated for the given editor, for the given language.
    total_code_suggestions *int32
    // Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language. Includes both full and partial acceptances.
    total_engaged_users *int32
}
// NewCopilotIdeCodeCompletions_editors_models_languages instantiates a new CopilotIdeCodeCompletions_editors_models_languages and sets the default values.
func NewCopilotIdeCodeCompletions_editors_models_languages()(*CopilotIdeCodeCompletions_editors_models_languages) {
    m := &CopilotIdeCodeCompletions_editors_models_languages{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeCodeCompletions_editors_models_languagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeCodeCompletions_editors_models_languagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeCodeCompletions_editors_models_languages(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["total_code_acceptances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCodeAcceptances(val)
        }
        return nil
    }
    res["total_code_lines_accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCodeLinesAccepted(val)
        }
        return nil
    }
    res["total_code_lines_suggested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCodeLinesSuggested(val)
        }
        return nil
    }
    res["total_code_suggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCodeSuggestions(val)
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
// GetName gets the name property value. Name of the language used for Copilot code completion suggestions, for the given editor.
// returns a *string when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetName()(*string) {
    return m.name
}
// GetTotalCodeAcceptances gets the total_code_acceptances property value. The number of Copilot code suggestions accepted for the given editor, for the given language. Includes both full and partial acceptances.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetTotalCodeAcceptances()(*int32) {
    return m.total_code_acceptances
}
// GetTotalCodeLinesAccepted gets the total_code_lines_accepted property value. The number of lines of code accepted from Copilot code suggestions for the given editor, for the given language.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetTotalCodeLinesAccepted()(*int32) {
    return m.total_code_lines_accepted
}
// GetTotalCodeLinesSuggested gets the total_code_lines_suggested property value. The number of lines of code suggested by Copilot code completions for the given editor, for the given language.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetTotalCodeLinesSuggested()(*int32) {
    return m.total_code_lines_suggested
}
// GetTotalCodeSuggestions gets the total_code_suggestions property value. The number of Copilot code suggestions generated for the given editor, for the given language.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetTotalCodeSuggestions()(*int32) {
    return m.total_code_suggestions
}
// GetTotalEngagedUsers gets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language. Includes both full and partial acceptances.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions_editors_models_languages) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeCodeCompletions_editors_models_languages) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_code_acceptances", m.GetTotalCodeAcceptances())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_code_lines_accepted", m.GetTotalCodeLinesAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_code_lines_suggested", m.GetTotalCodeLinesSuggested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_code_suggestions", m.GetTotalCodeSuggestions())
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
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the language used for Copilot code completion suggestions, for the given editor.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetName(value *string)() {
    m.name = value
}
// SetTotalCodeAcceptances sets the total_code_acceptances property value. The number of Copilot code suggestions accepted for the given editor, for the given language. Includes both full and partial acceptances.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetTotalCodeAcceptances(value *int32)() {
    m.total_code_acceptances = value
}
// SetTotalCodeLinesAccepted sets the total_code_lines_accepted property value. The number of lines of code accepted from Copilot code suggestions for the given editor, for the given language.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetTotalCodeLinesAccepted(value *int32)() {
    m.total_code_lines_accepted = value
}
// SetTotalCodeLinesSuggested sets the total_code_lines_suggested property value. The number of lines of code suggested by Copilot code completions for the given editor, for the given language.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetTotalCodeLinesSuggested(value *int32)() {
    m.total_code_lines_suggested = value
}
// SetTotalCodeSuggestions sets the total_code_suggestions property value. The number of Copilot code suggestions generated for the given editor, for the given language.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetTotalCodeSuggestions(value *int32)() {
    m.total_code_suggestions = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Number of users who accepted at least one Copilot code completion suggestion for the given editor, for the given language. Includes both full and partial acceptances.
func (m *CopilotIdeCodeCompletions_editors_models_languages) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeCodeCompletions_editors_models_languagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetTotalCodeAcceptances()(*int32)
    GetTotalCodeLinesAccepted()(*int32)
    GetTotalCodeLinesSuggested()(*int32)
    GetTotalCodeSuggestions()(*int32)
    GetTotalEngagedUsers()(*int32)
    SetName(value *string)()
    SetTotalCodeAcceptances(value *int32)()
    SetTotalCodeLinesAccepted(value *int32)()
    SetTotalCodeLinesSuggested(value *int32)()
    SetTotalCodeSuggestions(value *int32)()
    SetTotalEngagedUsers(value *int32)()
}
