package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecretScanningLocation 
type SecretScanningLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details property
    details SecretScanningLocation_SecretScanningLocation_detailsable
    // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
    typeEscaped *SecretScanningLocation_type
}
// SecretScanningLocation_SecretScanningLocation_details composed type wrapper for classes secretScanningLocationCommit, secretScanningLocationIssueBody, secretScanningLocationIssueComment, secretScanningLocationIssueTitle
type SecretScanningLocation_SecretScanningLocation_details struct {
    // Composed type representation for type secretScanningLocationCommit
    secretScanningLocationCommit SecretScanningLocationCommitable
    // Composed type representation for type secretScanningLocationIssueBody
    secretScanningLocationIssueBody SecretScanningLocationIssueBodyable
    // Composed type representation for type secretScanningLocationIssueComment
    secretScanningLocationIssueComment SecretScanningLocationIssueCommentable
    // Composed type representation for type secretScanningLocationIssueTitle
    secretScanningLocationIssueTitle SecretScanningLocationIssueTitleable
}
// NewSecretScanningLocation_SecretScanningLocation_details instantiates a new secretScanningLocation_details and sets the default values.
func NewSecretScanningLocation_SecretScanningLocation_details()(*SecretScanningLocation_SecretScanningLocation_details) {
    m := &SecretScanningLocation_SecretScanningLocation_details{
    }
    return m
}
// CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecretScanningLocation_SecretScanningLocation_details()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecretScanningLocation_SecretScanningLocation_details) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *SecretScanningLocation_SecretScanningLocation_details) GetIsComposedType()(bool) {
    return true
}
// GetSecretScanningLocationCommit gets the secretScanningLocationCommit property value. Composed type representation for type secretScanningLocationCommit
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationCommit()(SecretScanningLocationCommitable) {
    return m.secretScanningLocationCommit
}
// GetSecretScanningLocationIssueBody gets the secretScanningLocationIssueBody property value. Composed type representation for type secretScanningLocationIssueBody
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable) {
    return m.secretScanningLocationIssueBody
}
// GetSecretScanningLocationIssueComment gets the secretScanningLocationIssueComment property value. Composed type representation for type secretScanningLocationIssueComment
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable) {
    return m.secretScanningLocationIssueComment
}
// GetSecretScanningLocationIssueTitle gets the secretScanningLocationIssueTitle property value. Composed type representation for type secretScanningLocationIssueTitle
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable) {
    return m.secretScanningLocationIssueTitle
}
// Serialize serializes information the current object
func (m *SecretScanningLocation_SecretScanningLocation_details) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSecretScanningLocationCommit() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationCommit())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueBody() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueBody())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueTitle() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecretScanningLocationCommit sets the secretScanningLocationCommit property value. Composed type representation for type secretScanningLocationCommit
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)() {
    m.secretScanningLocationCommit = value
}
// SetSecretScanningLocationIssueBody sets the secretScanningLocationIssueBody property value. Composed type representation for type secretScanningLocationIssueBody
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)() {
    m.secretScanningLocationIssueBody = value
}
// SetSecretScanningLocationIssueComment sets the secretScanningLocationIssueComment property value. Composed type representation for type secretScanningLocationIssueComment
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)() {
    m.secretScanningLocationIssueComment = value
}
// SetSecretScanningLocationIssueTitle sets the secretScanningLocationIssueTitle property value. Composed type representation for type secretScanningLocationIssueTitle
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)() {
    m.secretScanningLocationIssueTitle = value
}
// SecretScanningLocation_SecretScanningLocation_detailsable 
type SecretScanningLocation_SecretScanningLocation_detailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecretScanningLocationCommit()(SecretScanningLocationCommitable)
    GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable)
    GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable)
    GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable)
    SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)()
    SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)()
    SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)()
    SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)()
}
// NewSecretScanningLocation instantiates a new secretScanningLocation and sets the default values.
func NewSecretScanningLocation()(*SecretScanningLocation) {
    m := &SecretScanningLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecretScanningLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecretScanningLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. The details property
func (m *SecretScanningLocation) GetDetails()(SecretScanningLocation_SecretScanningLocation_detailsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecretScanningLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(SecretScanningLocation_SecretScanningLocation_detailsable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningLocation_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SecretScanningLocation_type))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
func (m *SecretScanningLocation) GetTypeEscaped()(*SecretScanningLocation_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SecretScanningLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecretScanningLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. The details property
func (m *SecretScanningLocation) SetDetails(value SecretScanningLocation_SecretScanningLocation_detailsable)() {
    m.details = value
}
// SetTypeEscaped sets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found.
func (m *SecretScanningLocation) SetTypeEscaped(value *SecretScanningLocation_type)() {
    m.typeEscaped = value
}
// SecretScanningLocationable 
type SecretScanningLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()(SecretScanningLocation_SecretScanningLocation_detailsable)
    GetTypeEscaped()(*SecretScanningLocation_type)
    SetDetails(value SecretScanningLocation_SecretScanningLocation_detailsable)()
    SetTypeEscaped(value *SecretScanningLocation_type)()
}
