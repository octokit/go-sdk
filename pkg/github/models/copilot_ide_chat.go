package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotIdeChat usage metrics for Copilot Chat in the IDE.
type CopilotIdeChat struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The editors property
    editors []CopilotIdeChat_editorsable
    // Total number of users who prompted Copilot Chat in the IDE.
    total_engaged_users *int32
}
// NewCopilotIdeChat instantiates a new CopilotIdeChat and sets the default values.
func NewCopilotIdeChat()(*CopilotIdeChat) {
    m := &CopilotIdeChat{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotIdeChatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotIdeChatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotIdeChat(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotIdeChat) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEditors gets the editors property value. The editors property
// returns a []CopilotIdeChat_editorsable when successful
func (m *CopilotIdeChat) GetEditors()([]CopilotIdeChat_editorsable) {
    return m.editors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotIdeChat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["editors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotIdeChat_editorsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotIdeChat_editorsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotIdeChat_editorsable)
                }
            }
            m.SetEditors(res)
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
// GetTotalEngagedUsers gets the total_engaged_users property value. Total number of users who prompted Copilot Chat in the IDE.
// returns a *int32 when successful
func (m *CopilotIdeChat) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotIdeChat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CopilotIdeChat) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEditors sets the editors property value. The editors property
func (m *CopilotIdeChat) SetEditors(value []CopilotIdeChat_editorsable)() {
    m.editors = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. Total number of users who prompted Copilot Chat in the IDE.
func (m *CopilotIdeChat) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotIdeChatable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEditors()([]CopilotIdeChat_editorsable)
    GetTotalEngagedUsers()(*int32)
    SetEditors(value []CopilotIdeChat_editorsable)()
    SetTotalEngagedUsers(value *int32)()
}
