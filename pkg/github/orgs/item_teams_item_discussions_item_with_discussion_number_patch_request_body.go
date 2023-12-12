package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody 
type ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The discussion post's body text.
    body *string
    // The discussion post's title.
    title *string
}
// NewItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody instantiates a new ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody and sets the default values.
func NewItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody()(*ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) {
    m := &ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBody gets the body property value. The discussion post's body text.
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetTitle gets the title property value. The discussion post's title.
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBody sets the body property value. The discussion post's body text.
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) SetBody(value *string)() {
    m.body = value
}
// SetTitle sets the title property value. The discussion post's title.
func (m *ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBody) SetTitle(value *string)() {
    m.title = value
}
// ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBodyable 
type ItemTeamsItemDiscussionsItemWithDiscussion_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    GetTitle()(*string)
    SetBody(value *string)()
    SetTitle(value *string)()
}
