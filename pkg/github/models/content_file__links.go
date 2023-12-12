package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentFile__links 
type ContentFile__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The git property
    git *string
    // The html property
    html *string
    // The self property
    self *string
}
// NewContentFile__links instantiates a new contentFile__links and sets the default values.
func NewContentFile__links()(*ContentFile__links) {
    m := &ContentFile__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentFile__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentFile__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentFile__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentFile__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentFile__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["git"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGit(val)
        }
        return nil
    }
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val)
        }
        return nil
    }
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
        }
        return nil
    }
    return res
}
// GetGit gets the git property value. The git property
func (m *ContentFile__links) GetGit()(*string) {
    return m.git
}
// GetHtml gets the html property value. The html property
func (m *ContentFile__links) GetHtml()(*string) {
    return m.html
}
// GetSelf gets the self property value. The self property
func (m *ContentFile__links) GetSelf()(*string) {
    return m.self
}
// Serialize serializes information the current object
func (m *ContentFile__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("git", m.GetGit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("self", m.GetSelf())
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
func (m *ContentFile__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGit sets the git property value. The git property
func (m *ContentFile__links) SetGit(value *string)() {
    m.git = value
}
// SetHtml sets the html property value. The html property
func (m *ContentFile__links) SetHtml(value *string)() {
    m.html = value
}
// SetSelf sets the self property value. The self property
func (m *ContentFile__links) SetSelf(value *string)() {
    m.self = value
}
// ContentFile__linksable 
type ContentFile__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGit()(*string)
    GetHtml()(*string)
    GetSelf()(*string)
    SetGit(value *string)()
    SetHtml(value *string)()
    SetSelf(value *string)()
}
