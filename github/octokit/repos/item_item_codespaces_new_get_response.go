package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemCodespacesNewGetResponse 
type ItemItemCodespacesNewGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    billable_owner i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable
    // The defaults property
    defaults ItemItemCodespacesNewGetResponse_defaultsable
}
// NewItemItemCodespacesNewGetResponse instantiates a new ItemItemCodespacesNewGetResponse and sets the default values.
func NewItemItemCodespacesNewGetResponse()(*ItemItemCodespacesNewGetResponse) {
    m := &ItemItemCodespacesNewGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesNewGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemCodespacesNewGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesNewGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemCodespacesNewGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillableOwner gets the billable_owner property value. A GitHub user.
func (m *ItemItemCodespacesNewGetResponse) GetBillableOwner()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable) {
    return m.billable_owner
}
// GetDefaults gets the defaults property value. The defaults property
func (m *ItemItemCodespacesNewGetResponse) GetDefaults()(ItemItemCodespacesNewGetResponse_defaultsable) {
    return m.defaults
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemCodespacesNewGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billable_owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillableOwner(val.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable))
        }
        return nil
    }
    res["defaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemCodespacesNewGetResponse_defaultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaults(val.(ItemItemCodespacesNewGetResponse_defaultsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesNewGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("billable_owner", m.GetBillableOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaults", m.GetDefaults())
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
func (m *ItemItemCodespacesNewGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillableOwner sets the billable_owner property value. A GitHub user.
func (m *ItemItemCodespacesNewGetResponse) SetBillableOwner(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable)() {
    m.billable_owner = value
}
// SetDefaults sets the defaults property value. The defaults property
func (m *ItemItemCodespacesNewGetResponse) SetDefaults(value ItemItemCodespacesNewGetResponse_defaultsable)() {
    m.defaults = value
}
// ItemItemCodespacesNewGetResponseable 
type ItemItemCodespacesNewGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillableOwner()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable)
    GetDefaults()(ItemItemCodespacesNewGetResponse_defaultsable)
    SetBillableOwner(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.SimpleUserable)()
    SetDefaults(value ItemItemCodespacesNewGetResponse_defaultsable)()
}
