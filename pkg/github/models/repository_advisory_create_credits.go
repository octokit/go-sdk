package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryAdvisoryCreate_credits 
type RepositoryAdvisoryCreate_credits struct {
    // The username of the user credited.
    login *string
    // The type of credit the user is receiving.
    typeEscaped *SecurityAdvisoryCreditTypes
}
// NewRepositoryAdvisoryCreate_credits instantiates a new repositoryAdvisoryCreate_credits and sets the default values.
func NewRepositoryAdvisoryCreate_credits()(*RepositoryAdvisoryCreate_credits) {
    m := &RepositoryAdvisoryCreate_credits{
    }
    return m
}
// CreateRepositoryAdvisoryCreate_creditsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoryAdvisoryCreate_creditsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryCreate_credits(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RepositoryAdvisoryCreate_credits) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAdvisoryCreditTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SecurityAdvisoryCreditTypes))
        }
        return nil
    }
    return res
}
// GetLogin gets the login property value. The username of the user credited.
func (m *RepositoryAdvisoryCreate_credits) GetLogin()(*string) {
    return m.login
}
// GetTypeEscaped gets the type property value. The type of credit the user is receiving.
func (m *RepositoryAdvisoryCreate_credits) GetTypeEscaped()(*SecurityAdvisoryCreditTypes) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryCreate_credits) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLogin sets the login property value. The username of the user credited.
func (m *RepositoryAdvisoryCreate_credits) SetLogin(value *string)() {
    m.login = value
}
// SetTypeEscaped sets the type property value. The type of credit the user is receiving.
func (m *RepositoryAdvisoryCreate_credits) SetTypeEscaped(value *SecurityAdvisoryCreditTypes)() {
    m.typeEscaped = value
}
// RepositoryAdvisoryCreate_creditsable 
type RepositoryAdvisoryCreate_creditsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLogin()(*string)
    GetTypeEscaped()(*SecurityAdvisoryCreditTypes)
    SetLogin(value *string)()
    SetTypeEscaped(value *SecurityAdvisoryCreditTypes)()
}
