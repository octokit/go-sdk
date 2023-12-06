package appmanifests

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemConversionsPostResponse 
type ItemConversionsPostResponse struct {
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Integration
}
// NewItemConversionsPostResponse instantiates a new ItemConversionsPostResponse and sets the default values.
func NewItemConversionsPostResponse()(*ItemConversionsPostResponse) {
    m := &ItemConversionsPostResponse{
        Integration: *i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.NewIntegration(),
    }
    return m
}
// CreateItemConversionsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemConversionsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemConversionsPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemConversionsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Integration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ItemConversionsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Integration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// ItemConversionsPostResponseable 
type ItemConversionsPostResponseable interface {
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Integrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
