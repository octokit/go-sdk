package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotIdeCodeCompletions usage metrics for Copilot editor code completions in the IDE.
type CopilotIdeCodeCompletions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The editors property
    editors []CopilotIdeCodeCompletions_editorsable
    // Code completion metrics for active languages.
    languages []CopilotIdeCodeCompletions_languagesable
    // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
    total_engaged_users *int32
}
// NewCopilotIdeCodeCompletions instantiates a new CopilotIdeCodeCompletions and sets the default values.
func NewCopilotIdeCodeCompletions()(*CopilotIdeCodeCompletions) {
    m := &CopilotIdeCodeCompletions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeCodeCompletionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeCodeCompletionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeCodeCompletions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeCodeCompletions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEditors gets the editors property value. The editors property
// returns a []CopilotIdeCodeCompletions_editorsable when successful
func (m *CopilotIdeCodeCompletions) GetEditors()([]CopilotIdeCodeCompletions_editorsable) {
    return m.editors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeCodeCompletions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["editors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotIdeCodeCompletions_editorsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotIdeCodeCompletions_editorsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotIdeCodeCompletions_editorsable)
                }
            }
            m.SetEditors(res)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotIdeCodeCompletions_languagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotIdeCodeCompletions_languagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotIdeCodeCompletions_languagesable)
                }
            }
            m.SetLanguages(res)
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
// GetLanguages gets the languages property value. Code completion metrics for active languages.
// returns a []CopilotIdeCodeCompletions_languagesable when successful
func (m *CopilotIdeCodeCompletions) GetLanguages()([]CopilotIdeCodeCompletions_languagesable) {
    return m.languages
}
// GetTotalEngagedUsers gets the total_engaged_users property value. Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
// returns a *int32 when successful
func (m *CopilotIdeCodeCompletions) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeCodeCompletions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEditors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEditors()))
        for i, v := range m.GetEditors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("editors", cast)
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
func (m *CopilotIdeCodeCompletions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEditors sets the editors property value. The editors property
func (m *CopilotIdeCodeCompletions) SetEditors(value []CopilotIdeCodeCompletions_editorsable)() {
    m.editors = value
}
// SetLanguages sets the languages property value. Code completion metrics for active languages.
func (m *CopilotIdeCodeCompletions) SetLanguages(value []CopilotIdeCodeCompletions_languagesable)() {
    m.languages = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
func (m *CopilotIdeCodeCompletions) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeCodeCompletionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEditors()([]CopilotIdeCodeCompletions_editorsable)
    GetLanguages()([]CopilotIdeCodeCompletions_languagesable)
    GetTotalEngagedUsers()(*int32)
    SetEditors(value []CopilotIdeCodeCompletions_editorsable)()
    SetLanguages(value []CopilotIdeCodeCompletions_languagesable)()
    SetTotalEngagedUsers(value *int32)()
}
