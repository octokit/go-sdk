package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\runs\{run_id}\deployment_protection_rule
type ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Deployment_protection_rulePostRequestBody composed type wrapper for classes reviewCustomGatesCommentRequired, reviewCustomGatesStateRequired
type Deployment_protection_rulePostRequestBody struct {
    // Composed type representation for type reviewCustomGatesCommentRequired
    deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type reviewCustomGatesStateRequired
    deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable
    // Composed type representation for type reviewCustomGatesCommentRequired
    reviewCustomGatesCommentRequired i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable
    // Composed type representation for type reviewCustomGatesStateRequired
    reviewCustomGatesStateRequired i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable
}
// NewDeployment_protection_rulePostRequestBody instantiates a new deployment_protection_rulePostRequestBody and sets the default values.
func NewDeployment_protection_rulePostRequestBody()(*Deployment_protection_rulePostRequestBody) {
    m := &Deployment_protection_rulePostRequestBody{
    }
    return m
}
// CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeployment_protection_rulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeployment_protection_rulePostRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable); ok {
                result.SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReviewCustomGatesCommentRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable); ok {
                result.SetReviewCustomGatesCommentRequired(cast)
            }
        } else if val, err := parseNode.GetObjectValue(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateReviewCustomGatesStateRequiredFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable); ok {
                result.SetReviewCustomGatesStateRequired(cast)
            }
        }
    }
    return result, nil
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type reviewCustomGatesCommentRequired
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired
}
// GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type reviewCustomGatesStateRequired
func (m *Deployment_protection_rulePostRequestBody) GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable) {
    return m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Deployment_protection_rulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *Deployment_protection_rulePostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetReviewCustomGatesCommentRequired gets the reviewCustomGatesCommentRequired property value. Composed type representation for type reviewCustomGatesCommentRequired
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesCommentRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable) {
    return m.reviewCustomGatesCommentRequired
}
// GetReviewCustomGatesStateRequired gets the reviewCustomGatesStateRequired property value. Composed type representation for type reviewCustomGatesStateRequired
func (m *Deployment_protection_rulePostRequestBody) GetReviewCustomGatesStateRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable) {
    return m.reviewCustomGatesStateRequired
}
// Serialize serializes information the current object
func (m *Deployment_protection_rulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesCommentRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesCommentRequired())
        if err != nil {
            return err
        }
    } else if m.GetReviewCustomGatesStateRequired() != nil {
        err := writer.WriteObjectValue("", m.GetReviewCustomGatesStateRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type reviewCustomGatesCommentRequired
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesCommentRequired = value
}
// SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type reviewCustomGatesStateRequired
func (m *Deployment_protection_rulePostRequestBody) SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)() {
    m.deployment_protection_rulePostRequestBodyReviewCustomGatesStateRequired = value
}
// SetReviewCustomGatesCommentRequired sets the reviewCustomGatesCommentRequired property value. Composed type representation for type reviewCustomGatesCommentRequired
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesCommentRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)() {
    m.reviewCustomGatesCommentRequired = value
}
// SetReviewCustomGatesStateRequired sets the reviewCustomGatesStateRequired property value. Composed type representation for type reviewCustomGatesStateRequired
func (m *Deployment_protection_rulePostRequestBody) SetReviewCustomGatesStateRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)() {
    m.reviewCustomGatesStateRequired = value
}
// ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Deployment_protection_rulePostRequestBodyable 
type Deployment_protection_rulePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)
    GetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)
    GetReviewCustomGatesCommentRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)
    GetReviewCustomGatesStateRequired()(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesCommentRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)()
    SetDeploymentProtectionRulePostRequestBodyReviewCustomGatesStateRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)()
    SetReviewCustomGatesCommentRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesCommentRequiredable)()
    SetReviewCustomGatesStateRequired(value i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ReviewCustomGatesStateRequiredable)()
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal instantiates a new Deployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    m := &ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/runs/{run_id}/deployment_protection_rule", pathParameters),
    }
    return m
}
// NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder instantiates a new Deployment_protection_ruleRequestBuilder and sets the default values.
func NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."**Note:** GitHub Apps can only review their own custom deployment protection rules.To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have read and write permission for **Deployments** to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/workflow-runs#review-custom-deployment-protection-rules-for-a-workflow-run
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) Post(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation approve or reject custom deployment protection rules provided by a GitHub App for a workflow run. For more information, see "[Using environments for deployment](https://docs.github.com/actions/deployment/targeting-different-environments/using-environments-for-deployment)."**Note:** GitHub Apps can only review their own custom deployment protection rules.To approve or reject pending deployments that are waiting for review from a specific person or team, see [`POST /repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments`](/rest/actions/workflow-runs#review-pending-deployments-for-a-workflow-run).If the repository is private, you must use an access token with the `repo` scope.GitHub Apps must have read and write permission for **Deployments** to use this endpoint.
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) ToPostRequestInformation(ctx context.Context, body Deployment_protection_rulePostRequestBodyable, requestConfiguration *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder) {
    return NewItemItemActionsRunsItemDeployment_protection_ruleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
