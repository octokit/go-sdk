package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PullRequestReview__links_pull_request 
type PullRequestReview__links_pull_request struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The href property
    href *string
}
// NewPullRequestReview__links_pull_request instantiates a new pullRequestReview__links_pull_request and sets the default values.
func NewPullRequestReview__links_pull_request()(*PullRequestReview__links_pull_request) {
    m := &PullRequestReview__links_pull_request{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestReview__links_pull_requestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePullRequestReview__links_pull_requestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestReview__links_pull_request(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PullRequestReview__links_pull_request) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PullRequestReview__links_pull_request) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The href property
func (m *PullRequestReview__links_pull_request) GetHref()(*string) {
    return m.href
}
// Serialize serializes information the current object
func (m *PullRequestReview__links_pull_request) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
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
func (m *PullRequestReview__links_pull_request) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHref sets the href property value. The href property
func (m *PullRequestReview__links_pull_request) SetHref(value *string)() {
    m.href = value
}
// PullRequestReview__links_pull_requestable 
type PullRequestReview__links_pull_requestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHref()(*string)
    SetHref(value *string)()
}
