package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemActionsRunsItemAttemptsItemJobsGetResponse 
type ItemItemActionsRunsItemAttemptsItemJobsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The jobs property
    jobs []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable
    // The total_count property
    total_count *int32
}
// NewItemItemActionsRunsItemAttemptsItemJobsGetResponse instantiates a new ItemItemActionsRunsItemAttemptsItemJobsGetResponse and sets the default values.
func NewItemItemActionsRunsItemAttemptsItemJobsGetResponse()(*ItemItemActionsRunsItemAttemptsItemJobsGetResponse) {
    m := &ItemItemActionsRunsItemAttemptsItemJobsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsRunsItemAttemptsItemJobsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsItemAttemptsItemJobsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemAttemptsItemJobsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable)
                }
            }
            m.SetJobs(res)
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
// GetJobs gets the jobs property value. The jobs property
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) GetJobs()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable) {
    return m.jobs
}
// GetTotalCount gets the total_count property value. The total_count property
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("jobs", cast)
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
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJobs sets the jobs property value. The jobs property
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) SetJobs(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable)() {
    m.jobs = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsRunsItemAttemptsItemJobsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemActionsRunsItemAttemptsItemJobsGetResponseable 
type ItemItemActionsRunsItemAttemptsItemJobsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetJobs()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable)
    GetTotalCount()(*int32)
    SetJobs(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Jobable)()
    SetTotalCount(value *int32)()
}
