package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the team or role selected as a bypass reviewer
    reviewer_id *int32
}
// NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers instantiates a new ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers and sets the default values.
func NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers()(*ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) {
    m := &ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewer_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewerId(val)
        }
        return nil
    }
    return res
}
// GetReviewerId gets the reviewer_id property value. The ID of the team or role selected as a bypass reviewer
// returns a *int32 when successful
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) GetReviewerId()(*int32) {
    return m.reviewer_id
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("reviewer_id", m.GetReviewerId())
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
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewerId sets the reviewer_id property value. The ID of the team or role selected as a bypass reviewer
func (m *ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers) SetReviewerId(value *int32)() {
    m.reviewer_id = value
}
type ItemCodeSecurityConfigurationsItemWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewerId()(*int32)
    SetReviewerId(value *int32)()
}
