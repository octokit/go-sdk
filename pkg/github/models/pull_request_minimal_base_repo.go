package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestMinimal_base_repo 
type PullRequestMinimal_base_repo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The name property
    name *string
    // The url property
    url *string
}
// NewPullRequestMinimal_base_repo instantiates a new pullRequestMinimal_base_repo and sets the default values.
func NewPullRequestMinimal_base_repo()(*PullRequestMinimal_base_repo) {
    m := &PullRequestMinimal_base_repo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestMinimal_base_repoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePullRequestMinimal_base_repoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestMinimal_base_repo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequestMinimal_base_repo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PullRequestMinimal_base_repo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The id property
func (m *PullRequestMinimal_base_repo) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
func (m *PullRequestMinimal_base_repo) GetName()(*string) {
    return m.name
}
// GetUrl gets the url property value. The url property
func (m *PullRequestMinimal_base_repo) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PullRequestMinimal_base_repo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
func (m *PullRequestMinimal_base_repo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *PullRequestMinimal_base_repo) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *PullRequestMinimal_base_repo) SetName(value *string)() {
    m.name = value
}
// SetUrl sets the url property value. The url property
func (m *PullRequestMinimal_base_repo) SetUrl(value *string)() {
    m.url = value
}
// PullRequestMinimal_base_repoable 
type PullRequestMinimal_base_repoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetName()(*string)
    GetUrl()(*string)
    SetId(value *int32)()
    SetName(value *string)()
    SetUrl(value *string)()
}
