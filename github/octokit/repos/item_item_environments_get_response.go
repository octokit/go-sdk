package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemEnvironmentsGetResponse 
type ItemItemEnvironmentsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The environments property
    environments []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable
    // The number of environments in this repository
    total_count *int32
}
// NewItemItemEnvironmentsGetResponse instantiates a new ItemItemEnvironmentsGetResponse and sets the default values.
func NewItemItemEnvironmentsGetResponse()(*ItemItemEnvironmentsGetResponse) {
    m := &ItemItemEnvironmentsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemEnvironmentsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnvironments gets the environments property value. The environments property
func (m *ItemItemEnvironmentsGetResponse) GetEnvironments()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable) {
    return m.environments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemEnvironmentsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["environments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateEnvironmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable)
                }
            }
            m.SetEnvironments(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The number of environments in this repository
func (m *ItemItemEnvironmentsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnvironments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnvironments()))
        for i, v := range m.GetEnvironments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("environments", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemEnvironmentsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnvironments sets the environments property value. The environments property
func (m *ItemItemEnvironmentsGetResponse) SetEnvironments(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable)() {
    m.environments = value
}
// SetTotalCount sets the total_count property value. The number of environments in this repository
func (m *ItemItemEnvironmentsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemEnvironmentsGetResponseable 
type ItemItemEnvironmentsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnvironments()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable)
    GetTotalCount()(*int32)
    SetEnvironments(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Environmentable)()
    SetTotalCount(value *int32)()
}
