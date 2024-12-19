package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemSub_issueDeleteRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sub-issue to remove
    sub_issue_id *int32
}
// NewItemItemIssuesItemSub_issueDeleteRequestBody instantiates a new ItemItemIssuesItemSub_issueDeleteRequestBody and sets the default values.
func NewItemItemIssuesItemSub_issueDeleteRequestBody()(*ItemItemIssuesItemSub_issueDeleteRequestBody) {
    m := &ItemItemIssuesItemSub_issueDeleteRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemSub_issueDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemSub_issueDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemSub_issueDeleteRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sub_issue_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubIssueId(val)
        }
        return nil
    }
    return res
}
// GetSubIssueId gets the sub_issue_id property value. The sub-issue to remove
// returns a *int32 when successful
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) GetSubIssueId()(*int32) {
    return m.sub_issue_id
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("sub_issue_id", m.GetSubIssueId())
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
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSubIssueId sets the sub_issue_id property value. The sub-issue to remove
func (m *ItemItemIssuesItemSub_issueDeleteRequestBody) SetSubIssueId(value *int32)() {
    m.sub_issue_id = value
}
type ItemItemIssuesItemSub_issueDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSubIssueId()(*int32)
    SetSubIssueId(value *int32)()
}
