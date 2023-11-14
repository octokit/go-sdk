package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageDeployment the GitHub Pages deployment status.
type PageDeployment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URI to the deployed GitHub Pages.
    page_url *string
    // The URI to the deployed GitHub Pages preview.
    preview_url *string
    // The URI to monitor GitHub Pages deployment status.
    status_url *string
}
// NewPageDeployment instantiates a new pageDeployment and sets the default values.
func NewPageDeployment()(*PageDeployment) {
    m := &PageDeployment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePageDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePageDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageDeployment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PageDeployment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PageDeployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["page_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageUrl(val)
        }
        return nil
    }
    res["preview_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewUrl(val)
        }
        return nil
    }
    res["status_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusUrl(val)
        }
        return nil
    }
    return res
}
// GetPageUrl gets the page_url property value. The URI to the deployed GitHub Pages.
func (m *PageDeployment) GetPageUrl()(*string) {
    return m.page_url
}
// GetPreviewUrl gets the preview_url property value. The URI to the deployed GitHub Pages preview.
func (m *PageDeployment) GetPreviewUrl()(*string) {
    return m.preview_url
}
// GetStatusUrl gets the status_url property value. The URI to monitor GitHub Pages deployment status.
func (m *PageDeployment) GetStatusUrl()(*string) {
    return m.status_url
}
// Serialize serializes information the current object
func (m *PageDeployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("page_url", m.GetPageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("preview_url", m.GetPreviewUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status_url", m.GetStatusUrl())
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
func (m *PageDeployment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPageUrl sets the page_url property value. The URI to the deployed GitHub Pages.
func (m *PageDeployment) SetPageUrl(value *string)() {
    m.page_url = value
}
// SetPreviewUrl sets the preview_url property value. The URI to the deployed GitHub Pages preview.
func (m *PageDeployment) SetPreviewUrl(value *string)() {
    m.preview_url = value
}
// SetStatusUrl sets the status_url property value. The URI to monitor GitHub Pages deployment status.
func (m *PageDeployment) SetStatusUrl(value *string)() {
    m.status_url = value
}
// PageDeploymentable 
type PageDeploymentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPageUrl()(*string)
    GetPreviewUrl()(*string)
    GetStatusUrl()(*string)
    SetPageUrl(value *string)()
    SetPreviewUrl(value *string)()
    SetStatusUrl(value *string)()
}
