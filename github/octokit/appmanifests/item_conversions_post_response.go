package appmanifests

import (
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemConversionsPostResponse 
type ItemConversionsPostResponse struct {
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integration
}
// NewItemConversionsPostResponse instantiates a new ItemConversionsPostResponse and sets the default values.
func NewItemConversionsPostResponse()(*ItemConversionsPostResponse) {
    m := &ItemConversionsPostResponse{
        Integration: *i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.NewIntegration(),
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
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Integrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
