package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemSub_issuesPriorityPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id of the sub-issue to be prioritized after (either positional argument after OR before should be specified).
    after_id *int32
    // The id of the sub-issue to be prioritized before (either positional argument after OR before should be specified).
    before_id *int32
    // The id of the sub-issue to reprioritize
    sub_issue_id *int32
}
// NewItemItemIssuesItemSub_issuesPriorityPatchRequestBody instantiates a new ItemItemIssuesItemSub_issuesPriorityPatchRequestBody and sets the default values.
func NewItemItemIssuesItemSub_issuesPriorityPatchRequestBody()(*ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) {
    m := &ItemItemIssuesItemSub_issuesPriorityPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemSub_issuesPriorityPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemSub_issuesPriorityPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemSub_issuesPriorityPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfterId gets the after_id property value. The id of the sub-issue to be prioritized after (either positional argument after OR before should be specified).
// returns a *int32 when successful
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) GetAfterId()(*int32) {
    return m.after_id
}
// GetBeforeId gets the before_id property value. The id of the sub-issue to be prioritized before (either positional argument after OR before should be specified).
// returns a *int32 when successful
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) GetBeforeId()(*int32) {
    return m.before_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["after_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfterId(val)
        }
        return nil
    }
    res["before_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBeforeId(val)
        }
        return nil
    }
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
// GetSubIssueId gets the sub_issue_id property value. The id of the sub-issue to reprioritize
// returns a *int32 when successful
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) GetSubIssueId()(*int32) {
    return m.sub_issue_id
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("after_id", m.GetAfterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("before_id", m.GetBeforeId())
        if err != nil {
            return err
        }
    }
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
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfterId sets the after_id property value. The id of the sub-issue to be prioritized after (either positional argument after OR before should be specified).
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) SetAfterId(value *int32)() {
    m.after_id = value
}
// SetBeforeId sets the before_id property value. The id of the sub-issue to be prioritized before (either positional argument after OR before should be specified).
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) SetBeforeId(value *int32)() {
    m.before_id = value
}
// SetSubIssueId sets the sub_issue_id property value. The id of the sub-issue to reprioritize
func (m *ItemItemIssuesItemSub_issuesPriorityPatchRequestBody) SetSubIssueId(value *int32)() {
    m.sub_issue_id = value
}
type ItemItemIssuesItemSub_issuesPriorityPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAfterId()(*int32)
    GetBeforeId()(*int32)
    GetSubIssueId()(*int32)
    SetAfterId(value *int32)()
    SetBeforeId(value *int32)()
    SetSubIssueId(value *int32)()
}
