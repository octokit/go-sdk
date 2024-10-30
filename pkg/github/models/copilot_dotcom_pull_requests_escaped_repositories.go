package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CopilotDotcomPullRequests_repositories struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of model metrics for custom models and the default model.
    models []CopilotDotcomPullRequests_repositories_modelsable
    // Repository name
    name *string
    // The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository.
    total_engaged_users *int32
}
// NewCopilotDotcomPullRequests_repositories instantiates a new CopilotDotcomPullRequests_repositories and sets the default values.
func NewCopilotDotcomPullRequests_repositories()(*CopilotDotcomPullRequests_repositories) {
    m := &CopilotDotcomPullRequests_repositories{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotDotcomPullRequests_repositoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotDotcomPullRequests_repositoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotDotcomPullRequests_repositories(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotDotcomPullRequests_repositories) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotDotcomPullRequests_repositories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["models"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotDotcomPullRequests_repositories_modelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotDotcomPullRequests_repositories_modelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotDotcomPullRequests_repositories_modelsable)
                }
            }
            m.SetModels(res)
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
    return res
}
// GetModels gets the models property value. List of model metrics for custom models and the default model.
// returns a []CopilotDotcomPullRequests_repositories_modelsable when successful
func (m *CopilotDotcomPullRequests_repositories) GetModels()([]CopilotDotcomPullRequests_repositories_modelsable) {
    return m.models
}
// GetName gets the name property value. Repository name
// returns a *string when successful
func (m *CopilotDotcomPullRequests_repositories) GetName()(*string) {
    return m.name
}
// GetTotalEngagedUsers gets the total_engaged_users property value. The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository.
// returns a *int32 when successful
func (m *CopilotDotcomPullRequests_repositories) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotDotcomPullRequests_repositories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetModels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModels()))
        for i, v := range m.GetModels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("models", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopilotDotcomPullRequests_repositories) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetModels sets the models property value. List of model metrics for custom models and the default model.
func (m *CopilotDotcomPullRequests_repositories) SetModels(value []CopilotDotcomPullRequests_repositories_modelsable)() {
    m.models = value
}
// SetName sets the name property value. Repository name
func (m *CopilotDotcomPullRequests_repositories) SetName(value *string)() {
    m.name = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. The number of users who generated pull request summaries using Copilot for Pull Requests in the given repository.
func (m *CopilotDotcomPullRequests_repositories) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotDotcomPullRequests_repositoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetModels()([]CopilotDotcomPullRequests_repositories_modelsable)
    GetName()(*string)
    GetTotalEngagedUsers()(*int32)
    SetModels(value []CopilotDotcomPullRequests_repositories_modelsable)()
    SetName(value *string)()
    SetTotalEngagedUsers(value *int32)()
}
