package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemContentsItemWithPathPutRequestBody_committer the person that committed the file. Default: the authenticated user.
type ItemItemContentsItemWithPathPutRequestBody_committer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date property
    date *string
    // The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
    email *string
    // The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
    name *string
}
// NewItemItemContentsItemWithPathPutRequestBody_committer instantiates a new ItemItemContentsItemWithPathPutRequestBody_committer and sets the default values.
func NewItemItemContentsItemWithPathPutRequestBody_committer()(*ItemItemContentsItemWithPathPutRequestBody_committer) {
    m := &ItemItemContentsItemWithPathPutRequestBody_committer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemContentsItemWithPathPutRequestBody_committerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemContentsItemWithPathPutRequestBody_committerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemContentsItemWithPathPutRequestBody_committer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. The date property
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) GetDate()(*string) {
    return m.date
}
// GetEmail gets the email property value. The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    return res
}
// GetName gets the name property value. The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. The date property
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) SetDate(value *string)() {
    m.date = value
}
// SetEmail sets the email property value. The email of the author or committer of the commit. You'll receive a `422` status code if `email` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name of the author or committer of the commit. You'll receive a `422` status code if `name` is omitted.
func (m *ItemItemContentsItemWithPathPutRequestBody_committer) SetName(value *string)() {
    m.name = value
}
// ItemItemContentsItemWithPathPutRequestBody_committerable 
type ItemItemContentsItemWithPathPutRequestBody_committerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*string)
    GetEmail()(*string)
    GetName()(*string)
    SetDate(value *string)()
    SetEmail(value *string)()
    SetName(value *string)()
}
