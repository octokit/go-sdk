package repos

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemGitCommitsPostRequestBody_author information about the author of the commit. By default, the `author` will be the authenticated user and the current date. See the `author` and `committer` object below for details.
type ItemItemGitCommitsPostRequestBody_author struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The email of the author (or committer) of the commit
    email *string
    // The name of the author (or committer) of the commit
    name *string
}
// NewItemItemGitCommitsPostRequestBody_author instantiates a new ItemItemGitCommitsPostRequestBody_author and sets the default values.
func NewItemItemGitCommitsPostRequestBody_author()(*ItemItemGitCommitsPostRequestBody_author) {
    m := &ItemItemGitCommitsPostRequestBody_author{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGitCommitsPostRequestBody_authorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemGitCommitsPostRequestBody_authorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGitCommitsPostRequestBody_author(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemGitCommitsPostRequestBody_author) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDate gets the date property value. Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *ItemItemGitCommitsPostRequestBody_author) GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date
}
// GetEmail gets the email property value. The email of the author (or committer) of the commit
func (m *ItemItemGitCommitsPostRequestBody_author) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemGitCommitsPostRequestBody_author) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
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
// GetName gets the name property value. The name of the author (or committer) of the commit
func (m *ItemItemGitCommitsPostRequestBody_author) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemGitCommitsPostRequestBody_author) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("date", m.GetDate())
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
func (m *ItemItemGitCommitsPostRequestBody_author) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDate sets the date property value. Indicates when this commit was authored (or committed). This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *ItemItemGitCommitsPostRequestBody_author) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date = value
}
// SetEmail sets the email property value. The email of the author (or committer) of the commit
func (m *ItemItemGitCommitsPostRequestBody_author) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name of the author (or committer) of the commit
func (m *ItemItemGitCommitsPostRequestBody_author) SetName(value *string)() {
    m.name = value
}
// ItemItemGitCommitsPostRequestBody_authorable 
type ItemItemGitCommitsPostRequestBody_authorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmail()(*string)
    GetName()(*string)
    SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmail(value *string)()
    SetName(value *string)()
}
