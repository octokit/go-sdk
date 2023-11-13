package repos

import (
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse 
type ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The branch_policies property
    branch_policies []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable
    // The number of deployment branch policies for the environment.
    total_count *int32
}
// NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse instantiates a new ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse and sets the default values.
func NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse()(*ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) {
    m := &ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranchPolicies gets the branch_policies property value. The branch_policies property
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetBranchPolicies()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable) {
    return m.branch_policies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch_policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateDeploymentBranchPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable)
                }
            }
            m.SetBranchPolicies(res)
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
// GetTotalCount gets the total_count property value. The number of deployment branch policies for the environment.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBranchPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBranchPolicies()))
        for i, v := range m.GetBranchPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("branch_policies", cast)
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
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranchPolicies sets the branch_policies property value. The branch_policies property
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetBranchPolicies(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable)() {
    m.branch_policies = value
}
// SetTotalCount sets the total_count property value. The number of deployment branch policies for the environment.
func (m *ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable 
type ItemItemEnvironmentsItemDeploymentBranchPoliciesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranchPolicies()([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable)
    GetTotalCount()(*int32)
    SetBranchPolicies(value []i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.DeploymentBranchPolicyable)()
    SetTotalCount(value *int32)()
}
