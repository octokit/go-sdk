package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options feature options for secret scanning delegated bypass
type ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The bypass reviewers for secret scanning delegated bypass
    reviewers []ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable
}
// NewItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options instantiates a new ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options and sets the default values.
func NewItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options()(*ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) {
    m := &ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable)
                }
            }
            m.SetReviewers(res)
        }
        return nil
    }
    return res
}
// GetReviewers gets the reviewers property value. The bypass reviewers for secret scanning delegated bypass
// returns a []ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable when successful
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) GetReviewers()([]ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable) {
    return m.reviewers
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
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
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewers sets the reviewers property value. The bypass reviewers for secret scanning delegated bypass
func (m *ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options) SetReviewers(value []ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable)() {
    m.reviewers = value
}
type ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_optionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewers()([]ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable)
    SetReviewers(value []ItemCodeSecurityConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewersable)()
}
