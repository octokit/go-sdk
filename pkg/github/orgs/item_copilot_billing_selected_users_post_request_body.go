package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCopilotBillingSelected_usersPostRequestBody 
type ItemCopilotBillingSelected_usersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The usernames of the organization members to be granted access to GitHub Copilot.
    selected_usernames []string
}
// NewItemCopilotBillingSelected_usersPostRequestBody instantiates a new ItemCopilotBillingSelected_usersPostRequestBody and sets the default values.
func NewItemCopilotBillingSelected_usersPostRequestBody()(*ItemCopilotBillingSelected_usersPostRequestBody) {
    m := &ItemCopilotBillingSelected_usersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCopilotBillingSelected_usersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCopilotBillingSelected_usersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCopilotBillingSelected_usersPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCopilotBillingSelected_usersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCopilotBillingSelected_usersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_usernames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSelectedUsernames(res)
        }
        return nil
    }
    return res
}
// GetSelectedUsernames gets the selected_usernames property value. The usernames of the organization members to be granted access to GitHub Copilot.
func (m *ItemCopilotBillingSelected_usersPostRequestBody) GetSelectedUsernames()([]string) {
    return m.selected_usernames
}
// Serialize serializes information the current object
func (m *ItemCopilotBillingSelected_usersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedUsernames() != nil {
        err := writer.WriteCollectionOfStringValues("selected_usernames", m.GetSelectedUsernames())
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
func (m *ItemCopilotBillingSelected_usersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSelectedUsernames sets the selected_usernames property value. The usernames of the organization members to be granted access to GitHub Copilot.
func (m *ItemCopilotBillingSelected_usersPostRequestBody) SetSelectedUsernames(value []string)() {
    m.selected_usernames = value
}
// ItemCopilotBillingSelected_usersPostRequestBodyable 
type ItemCopilotBillingSelected_usersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedUsernames()([]string)
    SetSelectedUsernames(value []string)()
}
