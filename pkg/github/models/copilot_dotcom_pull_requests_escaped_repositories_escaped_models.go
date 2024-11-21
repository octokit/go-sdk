package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotDotcomPullRequests_repositories_models struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The training date for the custom model.
    custom_model_training_date *string
    // Indicates whether a model is custom or default.
    is_custom_model *bool
    // Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
    name *string
    // The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository and model.
    total_engaged_users *int32
    // The number of pull request summaries generated using Copilot for Pull Requests in the given repository.
    total_pr_summaries_created *int32
}
// NewCopilotDotcomPullRequests_repositories_models instantiates a new CopilotDotcomPullRequests_repositories_models and sets the default values.
func NewCopilotDotcomPullRequests_repositories_models()(*CopilotDotcomPullRequests_repositories_models) {
    m := &CopilotDotcomPullRequests_repositories_models{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotDotcomPullRequests_repositories_modelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotDotcomPullRequests_repositories_modelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotDotcomPullRequests_repositories_models(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomModelTrainingDate gets the custom_model_training_date property value. The training date for the custom model.
// returns a *string when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetCustomModelTrainingDate()(*string) {
    return m.custom_model_training_date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_model_training_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomModelTrainingDate(val)
        }
        return nil
    }
    res["is_custom_model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomModel(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["total_engaged_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEngagedUsers(val)
        }
        return nil
    }
    res["total_pr_summaries_created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPrSummariesCreated(val)
        }
        return nil
    }
    return res
}
// GetIsCustomModel gets the is_custom_model property value. Indicates whether a model is custom or default.
// returns a *bool when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetIsCustomModel()(*bool) {
    return m.is_custom_model
}
// GetName gets the name property value. Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
// returns a *string when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetName()(*string) {
    return m.name
}
// GetTotalEngagedUsers gets the total_engaged_users property value. The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository and model.
// returns a *int32 when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// GetTotalPrSummariesCreated gets the total_pr_summaries_created property value. The number of pull request summaries generated using Copilot for Pull Requests in the given repository.
// returns a *int32 when successful
func (m *CopilotDotcomPullRequests_repositories_models) GetTotalPrSummariesCreated()(*int32) {
    return m.total_pr_summaries_created
}
// Serialize serializes information the current object
func (m *CopilotDotcomPullRequests_repositories_models) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("custom_model_training_date", m.GetCustomModelTrainingDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_custom_model", m.GetIsCustomModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_engaged_users", m.GetTotalEngagedUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_pr_summaries_created", m.GetTotalPrSummariesCreated())
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
func (m *CopilotDotcomPullRequests_repositories_models) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomModelTrainingDate sets the custom_model_training_date property value. The training date for the custom model.
func (m *CopilotDotcomPullRequests_repositories_models) SetCustomModelTrainingDate(value *string)() {
    m.custom_model_training_date = value
}
// SetIsCustomModel sets the is_custom_model property value. Indicates whether a model is custom or default.
func (m *CopilotDotcomPullRequests_repositories_models) SetIsCustomModel(value *bool)() {
    m.is_custom_model = value
}
// SetName sets the name property value. Name of the model used for Copilot code completion suggestions. If the default model is used will appear as 'default'.
func (m *CopilotDotcomPullRequests_repositories_models) SetName(value *string)() {
    m.name = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository and model.
func (m *CopilotDotcomPullRequests_repositories_models) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
// SetTotalPrSummariesCreated sets the total_pr_summaries_created property value. The number of pull request summaries generated using Copilot for Pull Requests in the given repository.
func (m *CopilotDotcomPullRequests_repositories_models) SetTotalPrSummariesCreated(value *int32)() {
    m.total_pr_summaries_created = value
}
type CopilotDotcomPullRequests_repositories_modelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomModelTrainingDate()(*string)
    GetIsCustomModel()(*bool)
    GetName()(*string)
    GetTotalEngagedUsers()(*int32)
    GetTotalPrSummariesCreated()(*int32)
    SetCustomModelTrainingDate(value *string)()
    SetIsCustomModel(value *bool)()
    SetName(value *string)()
    SetTotalEngagedUsers(value *int32)()
    SetTotalPrSummariesCreated(value *int32)()
}
