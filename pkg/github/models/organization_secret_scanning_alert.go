package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationSecretScanningAlert struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The GitHub URL of the alert resource.
    html_url *string
    // The REST API URL of the code locations for this alert.
    locations_url *string
    // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
    multi_repo *bool
    // The security alert number.
    number *int32
    // Whether the secret was publicly leaked.
    publicly_leaked *bool
    // An optional comment when requesting a push protection bypass.
    push_protection_bypass_request_comment *string
    // The URL to a push protection bypass request.
    push_protection_bypass_request_html_url *string
    // A GitHub user.
    push_protection_bypass_request_reviewer NullableSimpleUserable
    // Whether push protection was bypassed for the detected secret.
    push_protection_bypassed *bool
    // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    push_protection_bypassed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    push_protection_bypassed_by NullableSimpleUserable
    // A GitHub repository.
    repository SimpleRepositoryable
    // **Required when the `state` is `resolved`.** The reason for resolving the alert.
    resolution *SecretScanningAlertResolution
    // The comment that was optionally added when this alert was closed
    resolution_comment *string
    // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    resolved_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    resolved_by NullableSimpleUserable
    // The secret that was detected.
    secret *string
    // The type of secret that secret scanning detected.
    secret_type *string
    // User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
    secret_type_display_name *string
    // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
    state *SecretScanningAlertState
    // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The REST API URL of the alert resource.
    url *string
    // The token status as of the latest validity check.
    validity *OrganizationSecretScanningAlert_validity
}
// NewOrganizationSecretScanningAlert instantiates a new OrganizationSecretScanningAlert and sets the default values.
func NewOrganizationSecretScanningAlert()(*OrganizationSecretScanningAlert) {
    m := &OrganizationSecretScanningAlert{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationSecretScanningAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationSecretScanningAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationSecretScanningAlert(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationSecretScanningAlert) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *OrganizationSecretScanningAlert) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationSecretScanningAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["locations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationsUrl(val)
        }
        return nil
    }
    res["multi_repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultiRepo(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["publicly_leaked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPubliclyLeaked(val)
        }
        return nil
    }
    res["push_protection_bypass_request_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassRequestComment(val)
        }
        return nil
    }
    res["push_protection_bypass_request_html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassRequestHtmlUrl(val)
        }
        return nil
    }
    res["push_protection_bypass_request_reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassRequestReviewer(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["push_protection_bypassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassed(val)
        }
        return nil
    }
    res["push_protection_bypassed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassedAt(val)
        }
        return nil
    }
    res["push_protection_bypassed_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(SimpleRepositoryable))
        }
        return nil
    }
    res["resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningAlertResolution)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolution(val.(*SecretScanningAlertResolution))
        }
        return nil
    }
    res["resolution_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolutionComment(val)
        }
        return nil
    }
    res["resolved_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedAt(val)
        }
        return nil
    }
    res["resolved_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    res["secret_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretType(val)
        }
        return nil
    }
    res["secret_type_display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretTypeDisplayName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecretScanningAlertState))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["validity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationSecretScanningAlert_validity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidity(val.(*OrganizationSecretScanningAlert_validity))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The GitHub URL of the alert resource.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLocationsUrl gets the locations_url property value. The REST API URL of the code locations for this alert.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetLocationsUrl()(*string) {
    return m.locations_url
}
// GetMultiRepo gets the multi_repo property value. Whether the detected secret was found in multiple repositories in the same organization or enterprise.
// returns a *bool when successful
func (m *OrganizationSecretScanningAlert) GetMultiRepo()(*bool) {
    return m.multi_repo
}
// GetNumber gets the number property value. The security alert number.
// returns a *int32 when successful
func (m *OrganizationSecretScanningAlert) GetNumber()(*int32) {
    return m.number
}
// GetPubliclyLeaked gets the publicly_leaked property value. Whether the secret was publicly leaked.
// returns a *bool when successful
func (m *OrganizationSecretScanningAlert) GetPubliclyLeaked()(*bool) {
    return m.publicly_leaked
}
// GetPushProtectionBypassed gets the push_protection_bypassed property value. Whether push protection was bypassed for the detected secret.
// returns a *bool when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassed()(*bool) {
    return m.push_protection_bypassed
}
// GetPushProtectionBypassedAt gets the push_protection_bypassed_at property value. The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.push_protection_bypassed_at
}
// GetPushProtectionBypassedBy gets the push_protection_bypassed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassedBy()(NullableSimpleUserable) {
    return m.push_protection_bypassed_by
}
// GetPushProtectionBypassRequestComment gets the push_protection_bypass_request_comment property value. An optional comment when requesting a push protection bypass.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassRequestComment()(*string) {
    return m.push_protection_bypass_request_comment
}
// GetPushProtectionBypassRequestHtmlUrl gets the push_protection_bypass_request_html_url property value. The URL to a push protection bypass request.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassRequestHtmlUrl()(*string) {
    return m.push_protection_bypass_request_html_url
}
// GetPushProtectionBypassRequestReviewer gets the push_protection_bypass_request_reviewer property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *OrganizationSecretScanningAlert) GetPushProtectionBypassRequestReviewer()(NullableSimpleUserable) {
    return m.push_protection_bypass_request_reviewer
}
// GetRepository gets the repository property value. A GitHub repository.
// returns a SimpleRepositoryable when successful
func (m *OrganizationSecretScanningAlert) GetRepository()(SimpleRepositoryable) {
    return m.repository
}
// GetResolution gets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
// returns a *SecretScanningAlertResolution when successful
func (m *OrganizationSecretScanningAlert) GetResolution()(*SecretScanningAlertResolution) {
    return m.resolution
}
// GetResolutionComment gets the resolution_comment property value. The comment that was optionally added when this alert was closed
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetResolutionComment()(*string) {
    return m.resolution_comment
}
// GetResolvedAt gets the resolved_at property value. The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *OrganizationSecretScanningAlert) GetResolvedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.resolved_at
}
// GetResolvedBy gets the resolved_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *OrganizationSecretScanningAlert) GetResolvedBy()(NullableSimpleUserable) {
    return m.resolved_by
}
// GetSecret gets the secret property value. The secret that was detected.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetSecret()(*string) {
    return m.secret
}
// GetSecretType gets the secret_type property value. The type of secret that secret scanning detected.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetSecretType()(*string) {
    return m.secret_type
}
// GetSecretTypeDisplayName gets the secret_type_display_name property value. User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetSecretTypeDisplayName()(*string) {
    return m.secret_type_display_name
}
// GetState gets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
// returns a *SecretScanningAlertState when successful
func (m *OrganizationSecretScanningAlert) GetState()(*SecretScanningAlertState) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *OrganizationSecretScanningAlert) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The REST API URL of the alert resource.
// returns a *string when successful
func (m *OrganizationSecretScanningAlert) GetUrl()(*string) {
    return m.url
}
// GetValidity gets the validity property value. The token status as of the latest validity check.
// returns a *OrganizationSecretScanningAlert_validity when successful
func (m *OrganizationSecretScanningAlert) GetValidity()(*OrganizationSecretScanningAlert_validity) {
    return m.validity
}
// Serialize serializes information the current object
func (m *OrganizationSecretScanningAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locations_url", m.GetLocationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("multi_repo", m.GetMultiRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("publicly_leaked", m.GetPubliclyLeaked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("push_protection_bypassed", m.GetPushProtectionBypassed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("push_protection_bypassed_at", m.GetPushProtectionBypassedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("push_protection_bypassed_by", m.GetPushProtectionBypassedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("push_protection_bypass_request_comment", m.GetPushProtectionBypassRequestComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("push_protection_bypass_request_html_url", m.GetPushProtectionBypassRequestHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("push_protection_bypass_request_reviewer", m.GetPushProtectionBypassRequestReviewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    if m.GetResolution() != nil {
        cast := (*m.GetResolution()).String()
        err := writer.WriteStringValue("resolution", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resolution_comment", m.GetResolutionComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("resolved_at", m.GetResolvedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resolved_by", m.GetResolvedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_type", m.GetSecretType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_type_display_name", m.GetSecretTypeDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetValidity() != nil {
        cast := (*m.GetValidity()).String()
        err := writer.WriteStringValue("validity", &cast)
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
func (m *OrganizationSecretScanningAlert) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *OrganizationSecretScanningAlert) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The GitHub URL of the alert resource.
func (m *OrganizationSecretScanningAlert) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLocationsUrl sets the locations_url property value. The REST API URL of the code locations for this alert.
func (m *OrganizationSecretScanningAlert) SetLocationsUrl(value *string)() {
    m.locations_url = value
}
// SetMultiRepo sets the multi_repo property value. Whether the detected secret was found in multiple repositories in the same organization or enterprise.
func (m *OrganizationSecretScanningAlert) SetMultiRepo(value *bool)() {
    m.multi_repo = value
}
// SetNumber sets the number property value. The security alert number.
func (m *OrganizationSecretScanningAlert) SetNumber(value *int32)() {
    m.number = value
}
// SetPubliclyLeaked sets the publicly_leaked property value. Whether the secret was publicly leaked.
func (m *OrganizationSecretScanningAlert) SetPubliclyLeaked(value *bool)() {
    m.publicly_leaked = value
}
// SetPushProtectionBypassed sets the push_protection_bypassed property value. Whether push protection was bypassed for the detected secret.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassed(value *bool)() {
    m.push_protection_bypassed = value
}
// SetPushProtectionBypassedAt sets the push_protection_bypassed_at property value. The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.push_protection_bypassed_at = value
}
// SetPushProtectionBypassedBy sets the push_protection_bypassed_by property value. A GitHub user.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassedBy(value NullableSimpleUserable)() {
    m.push_protection_bypassed_by = value
}
// SetPushProtectionBypassRequestComment sets the push_protection_bypass_request_comment property value. An optional comment when requesting a push protection bypass.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassRequestComment(value *string)() {
    m.push_protection_bypass_request_comment = value
}
// SetPushProtectionBypassRequestHtmlUrl sets the push_protection_bypass_request_html_url property value. The URL to a push protection bypass request.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassRequestHtmlUrl(value *string)() {
    m.push_protection_bypass_request_html_url = value
}
// SetPushProtectionBypassRequestReviewer sets the push_protection_bypass_request_reviewer property value. A GitHub user.
func (m *OrganizationSecretScanningAlert) SetPushProtectionBypassRequestReviewer(value NullableSimpleUserable)() {
    m.push_protection_bypass_request_reviewer = value
}
// SetRepository sets the repository property value. A GitHub repository.
func (m *OrganizationSecretScanningAlert) SetRepository(value SimpleRepositoryable)() {
    m.repository = value
}
// SetResolution sets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
func (m *OrganizationSecretScanningAlert) SetResolution(value *SecretScanningAlertResolution)() {
    m.resolution = value
}
// SetResolutionComment sets the resolution_comment property value. The comment that was optionally added when this alert was closed
func (m *OrganizationSecretScanningAlert) SetResolutionComment(value *string)() {
    m.resolution_comment = value
}
// SetResolvedAt sets the resolved_at property value. The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *OrganizationSecretScanningAlert) SetResolvedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.resolved_at = value
}
// SetResolvedBy sets the resolved_by property value. A GitHub user.
func (m *OrganizationSecretScanningAlert) SetResolvedBy(value NullableSimpleUserable)() {
    m.resolved_by = value
}
// SetSecret sets the secret property value. The secret that was detected.
func (m *OrganizationSecretScanningAlert) SetSecret(value *string)() {
    m.secret = value
}
// SetSecretType sets the secret_type property value. The type of secret that secret scanning detected.
func (m *OrganizationSecretScanningAlert) SetSecretType(value *string)() {
    m.secret_type = value
}
// SetSecretTypeDisplayName sets the secret_type_display_name property value. User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
func (m *OrganizationSecretScanningAlert) SetSecretTypeDisplayName(value *string)() {
    m.secret_type_display_name = value
}
// SetState sets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
func (m *OrganizationSecretScanningAlert) SetState(value *SecretScanningAlertState)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *OrganizationSecretScanningAlert) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The REST API URL of the alert resource.
func (m *OrganizationSecretScanningAlert) SetUrl(value *string)() {
    m.url = value
}
// SetValidity sets the validity property value. The token status as of the latest validity check.
func (m *OrganizationSecretScanningAlert) SetValidity(value *OrganizationSecretScanningAlert_validity)() {
    m.validity = value
}
type OrganizationSecretScanningAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetLocationsUrl()(*string)
    GetMultiRepo()(*bool)
    GetNumber()(*int32)
    GetPubliclyLeaked()(*bool)
    GetPushProtectionBypassed()(*bool)
    GetPushProtectionBypassedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPushProtectionBypassedBy()(NullableSimpleUserable)
    GetPushProtectionBypassRequestComment()(*string)
    GetPushProtectionBypassRequestHtmlUrl()(*string)
    GetPushProtectionBypassRequestReviewer()(NullableSimpleUserable)
    GetRepository()(SimpleRepositoryable)
    GetResolution()(*SecretScanningAlertResolution)
    GetResolutionComment()(*string)
    GetResolvedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResolvedBy()(NullableSimpleUserable)
    GetSecret()(*string)
    GetSecretType()(*string)
    GetSecretTypeDisplayName()(*string)
    GetState()(*SecretScanningAlertState)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetValidity()(*OrganizationSecretScanningAlert_validity)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetLocationsUrl(value *string)()
    SetMultiRepo(value *bool)()
    SetNumber(value *int32)()
    SetPubliclyLeaked(value *bool)()
    SetPushProtectionBypassed(value *bool)()
    SetPushProtectionBypassedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPushProtectionBypassedBy(value NullableSimpleUserable)()
    SetPushProtectionBypassRequestComment(value *string)()
    SetPushProtectionBypassRequestHtmlUrl(value *string)()
    SetPushProtectionBypassRequestReviewer(value NullableSimpleUserable)()
    SetRepository(value SimpleRepositoryable)()
    SetResolution(value *SecretScanningAlertResolution)()
    SetResolutionComment(value *string)()
    SetResolvedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResolvedBy(value NullableSimpleUserable)()
    SetSecret(value *string)()
    SetSecretType(value *string)()
    SetSecretTypeDisplayName(value *string)()
    SetState(value *SecretScanningAlertState)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetValidity(value *OrganizationSecretScanningAlert_validity)()
}
