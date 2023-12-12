package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NullableCommunityHealthFile 
type NullableCommunityHealthFile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The html_url property
    html_url *string
    // The url property
    url *string
}
// NewNullableCommunityHealthFile instantiates a new nullableCommunityHealthFile and sets the default values.
func NewNullableCommunityHealthFile()(*NullableCommunityHealthFile) {
    m := &NullableCommunityHealthFile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNullableCommunityHealthFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNullableCommunityHealthFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNullableCommunityHealthFile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NullableCommunityHealthFile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NullableCommunityHealthFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
func (m *NullableCommunityHealthFile) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetUrl gets the url property value. The url property
func (m *NullableCommunityHealthFile) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *NullableCommunityHealthFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *NullableCommunityHealthFile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *NullableCommunityHealthFile) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetUrl sets the url property value. The url property
func (m *NullableCommunityHealthFile) SetUrl(value *string)() {
    m.url = value
}
// NullableCommunityHealthFileable 
type NullableCommunityHealthFileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHtmlUrl()(*string)
    GetUrl()(*string)
    SetHtmlUrl(value *string)()
    SetUrl(value *string)()
}
