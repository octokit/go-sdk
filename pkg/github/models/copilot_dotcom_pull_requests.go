package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotDotcomPullRequests usage metrics for Copilot for pull requests.
type CopilotDotcomPullRequests struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
    repositories []CopilotDotcomPullRequests_repositoriesable
    // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
    total_engaged_users *int32
}
// NewCopilotDotcomPullRequests instantiates a new CopilotDotcomPullRequests and sets the default values.
func NewCopilotDotcomPullRequests()(*CopilotDotcomPullRequests) {
    m := &CopilotDotcomPullRequests{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotDotcomPullRequestsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotDotcomPullRequestsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotDotcomPullRequests(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotDotcomPullRequests) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotDotcomPullRequests) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotDotcomPullRequests_repositoriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotDotcomPullRequests_repositoriesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotDotcomPullRequests_repositoriesable)
                }
            }
            m.SetRepositories(res)
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
// GetRepositories gets the repositories property value. Repositories in which users used Copilot for Pull Requests to generate pull request summaries
// returns a []CopilotDotcomPullRequests_repositoriesable when successful
func (m *CopilotDotcomPullRequests) GetRepositories()([]CopilotDotcomPullRequests_repositoriesable) {
    return m.repositories
}
// GetTotalEngagedUsers gets the total_engaged_users property value. The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
// returns a *int32 when successful
func (m *CopilotDotcomPullRequests) GetTotalEngagedUsers()(*int32) {
    return m.total_engaged_users
}
// Serialize serializes information the current object
func (m *CopilotDotcomPullRequests) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
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
func (m *CopilotDotcomPullRequests) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. Repositories in which users used Copilot for Pull Requests to generate pull request summaries
func (m *CopilotDotcomPullRequests) SetRepositories(value []CopilotDotcomPullRequests_repositoriesable)() {
    m.repositories = value
}
// SetTotalEngagedUsers sets the total_engaged_users property value. The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
func (m *CopilotDotcomPullRequests) SetTotalEngagedUsers(value *int32)() {
    m.total_engaged_users = value
}
type CopilotDotcomPullRequestsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositories()([]CopilotDotcomPullRequests_repositoriesable)
    GetTotalEngagedUsers()(*int32)
    SetRepositories(value []CopilotDotcomPullRequests_repositoriesable)()
    SetTotalEngagedUsers(value *int32)()
}
