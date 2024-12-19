package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSecurityConfiguration a code security configuration
type CodeSecurityConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enablement status of GitHub Advanced Security
    advanced_security *CodeSecurityConfiguration_advanced_security
    // The enablement status of code scanning default setup
    code_scanning_default_setup *CodeSecurityConfiguration_code_scanning_default_setup
    // Feature options for code scanning default setup
    code_scanning_default_setup_options CodeSecurityConfiguration_code_scanning_default_setup_optionsable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The enablement status of Dependabot alerts
    dependabot_alerts *CodeSecurityConfiguration_dependabot_alerts
    // The enablement status of Dependabot security updates
    dependabot_security_updates *CodeSecurityConfiguration_dependabot_security_updates
    // The enablement status of Dependency Graph
    dependency_graph *CodeSecurityConfiguration_dependency_graph
    // The enablement status of Automatic dependency submission
    dependency_graph_autosubmit_action *CodeSecurityConfiguration_dependency_graph_autosubmit_action
    // Feature options for Automatic dependency submission
    dependency_graph_autosubmit_action_options CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable
    // A description of the code security configuration
    description *string
    // The enforcement status for a security configuration
    enforcement *CodeSecurityConfiguration_enforcement
    // The URL of the configuration
    html_url *string
    // The ID of the code security configuration
    id *int32
    // The name of the code security configuration. Must be unique within the organization.
    name *string
    // The enablement status of private vulnerability reporting
    private_vulnerability_reporting *CodeSecurityConfiguration_private_vulnerability_reporting
    // The enablement status of secret scanning
    secret_scanning *CodeSecurityConfiguration_secret_scanning
    // The enablement status of secret scanning delegated bypass
    secret_scanning_delegated_bypass *CodeSecurityConfiguration_secret_scanning_delegated_bypass
    // Feature options for secret scanning delegated bypass
    secret_scanning_delegated_bypass_options CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable
    // The enablement status of secret scanning non-provider patterns
    secret_scanning_non_provider_patterns *CodeSecurityConfiguration_secret_scanning_non_provider_patterns
    // The enablement status of secret scanning push protection
    secret_scanning_push_protection *CodeSecurityConfiguration_secret_scanning_push_protection
    // The enablement status of secret scanning validity checks
    secret_scanning_validity_checks *CodeSecurityConfiguration_secret_scanning_validity_checks
    // The type of the code security configuration.
    target_type *CodeSecurityConfiguration_target_type
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL of the configuration
    url *string
}
// NewCodeSecurityConfiguration instantiates a new CodeSecurityConfiguration and sets the default values.
func NewCodeSecurityConfiguration()(*CodeSecurityConfiguration) {
    m := &CodeSecurityConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurity gets the advanced_security property value. The enablement status of GitHub Advanced Security
// returns a *CodeSecurityConfiguration_advanced_security when successful
func (m *CodeSecurityConfiguration) GetAdvancedSecurity()(*CodeSecurityConfiguration_advanced_security) {
    return m.advanced_security
}
// GetCodeScanningDefaultSetup gets the code_scanning_default_setup property value. The enablement status of code scanning default setup
// returns a *CodeSecurityConfiguration_code_scanning_default_setup when successful
func (m *CodeSecurityConfiguration) GetCodeScanningDefaultSetup()(*CodeSecurityConfiguration_code_scanning_default_setup) {
    return m.code_scanning_default_setup
}
// GetCodeScanningDefaultSetupOptions gets the code_scanning_default_setup_options property value. Feature options for code scanning default setup
// returns a CodeSecurityConfiguration_code_scanning_default_setup_optionsable when successful
func (m *CodeSecurityConfiguration) GetCodeScanningDefaultSetupOptions()(CodeSecurityConfiguration_code_scanning_default_setup_optionsable) {
    return m.code_scanning_default_setup_options
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *CodeSecurityConfiguration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDependabotAlerts gets the dependabot_alerts property value. The enablement status of Dependabot alerts
// returns a *CodeSecurityConfiguration_dependabot_alerts when successful
func (m *CodeSecurityConfiguration) GetDependabotAlerts()(*CodeSecurityConfiguration_dependabot_alerts) {
    return m.dependabot_alerts
}
// GetDependabotSecurityUpdates gets the dependabot_security_updates property value. The enablement status of Dependabot security updates
// returns a *CodeSecurityConfiguration_dependabot_security_updates when successful
func (m *CodeSecurityConfiguration) GetDependabotSecurityUpdates()(*CodeSecurityConfiguration_dependabot_security_updates) {
    return m.dependabot_security_updates
}
// GetDependencyGraph gets the dependency_graph property value. The enablement status of Dependency Graph
// returns a *CodeSecurityConfiguration_dependency_graph when successful
func (m *CodeSecurityConfiguration) GetDependencyGraph()(*CodeSecurityConfiguration_dependency_graph) {
    return m.dependency_graph
}
// GetDependencyGraphAutosubmitAction gets the dependency_graph_autosubmit_action property value. The enablement status of Automatic dependency submission
// returns a *CodeSecurityConfiguration_dependency_graph_autosubmit_action when successful
func (m *CodeSecurityConfiguration) GetDependencyGraphAutosubmitAction()(*CodeSecurityConfiguration_dependency_graph_autosubmit_action) {
    return m.dependency_graph_autosubmit_action
}
// GetDependencyGraphAutosubmitActionOptions gets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
// returns a CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable when successful
func (m *CodeSecurityConfiguration) GetDependencyGraphAutosubmitActionOptions()(CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable) {
    return m.dependency_graph_autosubmit_action_options
}
// GetDescription gets the description property value. A description of the code security configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetDescription()(*string) {
    return m.description
}
// GetEnforcement gets the enforcement property value. The enforcement status for a security configuration
// returns a *CodeSecurityConfiguration_enforcement when successful
func (m *CodeSecurityConfiguration) GetEnforcement()(*CodeSecurityConfiguration_enforcement) {
    return m.enforcement
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_advanced_security)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurity(val.(*CodeSecurityConfiguration_advanced_security))
        }
        return nil
    }
    res["code_scanning_default_setup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_code_scanning_default_setup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningDefaultSetup(val.(*CodeSecurityConfiguration_code_scanning_default_setup))
        }
        return nil
    }
    res["code_scanning_default_setup_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeSecurityConfiguration_code_scanning_default_setup_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningDefaultSetupOptions(val.(CodeSecurityConfiguration_code_scanning_default_setup_optionsable))
        }
        return nil
    }
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
    res["dependabot_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependabot_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotAlerts(val.(*CodeSecurityConfiguration_dependabot_alerts))
        }
        return nil
    }
    res["dependabot_security_updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependabot_security_updates)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecurityUpdates(val.(*CodeSecurityConfiguration_dependabot_security_updates))
        }
        return nil
    }
    res["dependency_graph"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependency_graph)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraph(val.(*CodeSecurityConfiguration_dependency_graph))
        }
        return nil
    }
    res["dependency_graph_autosubmit_action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependency_graph_autosubmit_action)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraphAutosubmitAction(val.(*CodeSecurityConfiguration_dependency_graph_autosubmit_action))
        }
        return nil
    }
    res["dependency_graph_autosubmit_action_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraphAutosubmitActionOptions(val.(CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["enforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_enforcement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcement(val.(*CodeSecurityConfiguration_enforcement))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["private_vulnerability_reporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_private_vulnerability_reporting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateVulnerabilityReporting(val.(*CodeSecurityConfiguration_private_vulnerability_reporting))
        }
        return nil
    }
    res["secret_scanning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanning(val.(*CodeSecurityConfiguration_secret_scanning))
        }
        return nil
    }
    res["secret_scanning_delegated_bypass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning_delegated_bypass)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningDelegatedBypass(val.(*CodeSecurityConfiguration_secret_scanning_delegated_bypass))
        }
        return nil
    }
    res["secret_scanning_delegated_bypass_options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningDelegatedBypassOptions(val.(CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable))
        }
        return nil
    }
    res["secret_scanning_non_provider_patterns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning_non_provider_patterns)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningNonProviderPatterns(val.(*CodeSecurityConfiguration_secret_scanning_non_provider_patterns))
        }
        return nil
    }
    res["secret_scanning_push_protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning_push_protection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtection(val.(*CodeSecurityConfiguration_secret_scanning_push_protection))
        }
        return nil
    }
    res["secret_scanning_validity_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning_validity_checks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningValidityChecks(val.(*CodeSecurityConfiguration_secret_scanning_validity_checks))
        }
        return nil
    }
    res["target_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_target_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*CodeSecurityConfiguration_target_type))
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
    return res
}
// GetHtmlUrl gets the html_url property value. The URL of the configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The ID of the code security configuration
// returns a *int32 when successful
func (m *CodeSecurityConfiguration) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the code security configuration. Must be unique within the organization.
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetName()(*string) {
    return m.name
}
// GetPrivateVulnerabilityReporting gets the private_vulnerability_reporting property value. The enablement status of private vulnerability reporting
// returns a *CodeSecurityConfiguration_private_vulnerability_reporting when successful
func (m *CodeSecurityConfiguration) GetPrivateVulnerabilityReporting()(*CodeSecurityConfiguration_private_vulnerability_reporting) {
    return m.private_vulnerability_reporting
}
// GetSecretScanning gets the secret_scanning property value. The enablement status of secret scanning
// returns a *CodeSecurityConfiguration_secret_scanning when successful
func (m *CodeSecurityConfiguration) GetSecretScanning()(*CodeSecurityConfiguration_secret_scanning) {
    return m.secret_scanning
}
// GetSecretScanningDelegatedBypass gets the secret_scanning_delegated_bypass property value. The enablement status of secret scanning delegated bypass
// returns a *CodeSecurityConfiguration_secret_scanning_delegated_bypass when successful
func (m *CodeSecurityConfiguration) GetSecretScanningDelegatedBypass()(*CodeSecurityConfiguration_secret_scanning_delegated_bypass) {
    return m.secret_scanning_delegated_bypass
}
// GetSecretScanningDelegatedBypassOptions gets the secret_scanning_delegated_bypass_options property value. Feature options for secret scanning delegated bypass
// returns a CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable when successful
func (m *CodeSecurityConfiguration) GetSecretScanningDelegatedBypassOptions()(CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable) {
    return m.secret_scanning_delegated_bypass_options
}
// GetSecretScanningNonProviderPatterns gets the secret_scanning_non_provider_patterns property value. The enablement status of secret scanning non-provider patterns
// returns a *CodeSecurityConfiguration_secret_scanning_non_provider_patterns when successful
func (m *CodeSecurityConfiguration) GetSecretScanningNonProviderPatterns()(*CodeSecurityConfiguration_secret_scanning_non_provider_patterns) {
    return m.secret_scanning_non_provider_patterns
}
// GetSecretScanningPushProtection gets the secret_scanning_push_protection property value. The enablement status of secret scanning push protection
// returns a *CodeSecurityConfiguration_secret_scanning_push_protection when successful
func (m *CodeSecurityConfiguration) GetSecretScanningPushProtection()(*CodeSecurityConfiguration_secret_scanning_push_protection) {
    return m.secret_scanning_push_protection
}
// GetSecretScanningValidityChecks gets the secret_scanning_validity_checks property value. The enablement status of secret scanning validity checks
// returns a *CodeSecurityConfiguration_secret_scanning_validity_checks when successful
func (m *CodeSecurityConfiguration) GetSecretScanningValidityChecks()(*CodeSecurityConfiguration_secret_scanning_validity_checks) {
    return m.secret_scanning_validity_checks
}
// GetTargetType gets the target_type property value. The type of the code security configuration.
// returns a *CodeSecurityConfiguration_target_type when successful
func (m *CodeSecurityConfiguration) GetTargetType()(*CodeSecurityConfiguration_target_type) {
    return m.target_type
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CodeSecurityConfiguration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The URL of the configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeSecurityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdvancedSecurity() != nil {
        cast := (*m.GetAdvancedSecurity()).String()
        err := writer.WriteStringValue("advanced_security", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCodeScanningDefaultSetup() != nil {
        cast := (*m.GetCodeScanningDefaultSetup()).String()
        err := writer.WriteStringValue("code_scanning_default_setup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("code_scanning_default_setup_options", m.GetCodeScanningDefaultSetupOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetDependabotAlerts() != nil {
        cast := (*m.GetDependabotAlerts()).String()
        err := writer.WriteStringValue("dependabot_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependabotSecurityUpdates() != nil {
        cast := (*m.GetDependabotSecurityUpdates()).String()
        err := writer.WriteStringValue("dependabot_security_updates", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependencyGraph() != nil {
        cast := (*m.GetDependencyGraph()).String()
        err := writer.WriteStringValue("dependency_graph", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependencyGraphAutosubmitAction() != nil {
        cast := (*m.GetDependencyGraphAutosubmitAction()).String()
        err := writer.WriteStringValue("dependency_graph_autosubmit_action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dependency_graph_autosubmit_action_options", m.GetDependencyGraphAutosubmitActionOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcement() != nil {
        cast := (*m.GetEnforcement()).String()
        err := writer.WriteStringValue("enforcement", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
    if m.GetPrivateVulnerabilityReporting() != nil {
        cast := (*m.GetPrivateVulnerabilityReporting()).String()
        err := writer.WriteStringValue("private_vulnerability_reporting", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanning() != nil {
        cast := (*m.GetSecretScanning()).String()
        err := writer.WriteStringValue("secret_scanning", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningDelegatedBypass() != nil {
        cast := (*m.GetSecretScanningDelegatedBypass()).String()
        err := writer.WriteStringValue("secret_scanning_delegated_bypass", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_delegated_bypass_options", m.GetSecretScanningDelegatedBypassOptions())
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningNonProviderPatterns() != nil {
        cast := (*m.GetSecretScanningNonProviderPatterns()).String()
        err := writer.WriteStringValue("secret_scanning_non_provider_patterns", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningPushProtection() != nil {
        cast := (*m.GetSecretScanningPushProtection()).String()
        err := writer.WriteStringValue("secret_scanning_push_protection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningValidityChecks() != nil {
        cast := (*m.GetSecretScanningValidityChecks()).String()
        err := writer.WriteStringValue("secret_scanning_validity_checks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("target_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *CodeSecurityConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurity sets the advanced_security property value. The enablement status of GitHub Advanced Security
func (m *CodeSecurityConfiguration) SetAdvancedSecurity(value *CodeSecurityConfiguration_advanced_security)() {
    m.advanced_security = value
}
// SetCodeScanningDefaultSetup sets the code_scanning_default_setup property value. The enablement status of code scanning default setup
func (m *CodeSecurityConfiguration) SetCodeScanningDefaultSetup(value *CodeSecurityConfiguration_code_scanning_default_setup)() {
    m.code_scanning_default_setup = value
}
// SetCodeScanningDefaultSetupOptions sets the code_scanning_default_setup_options property value. Feature options for code scanning default setup
func (m *CodeSecurityConfiguration) SetCodeScanningDefaultSetupOptions(value CodeSecurityConfiguration_code_scanning_default_setup_optionsable)() {
    m.code_scanning_default_setup_options = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *CodeSecurityConfiguration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDependabotAlerts sets the dependabot_alerts property value. The enablement status of Dependabot alerts
func (m *CodeSecurityConfiguration) SetDependabotAlerts(value *CodeSecurityConfiguration_dependabot_alerts)() {
    m.dependabot_alerts = value
}
// SetDependabotSecurityUpdates sets the dependabot_security_updates property value. The enablement status of Dependabot security updates
func (m *CodeSecurityConfiguration) SetDependabotSecurityUpdates(value *CodeSecurityConfiguration_dependabot_security_updates)() {
    m.dependabot_security_updates = value
}
// SetDependencyGraph sets the dependency_graph property value. The enablement status of Dependency Graph
func (m *CodeSecurityConfiguration) SetDependencyGraph(value *CodeSecurityConfiguration_dependency_graph)() {
    m.dependency_graph = value
}
// SetDependencyGraphAutosubmitAction sets the dependency_graph_autosubmit_action property value. The enablement status of Automatic dependency submission
func (m *CodeSecurityConfiguration) SetDependencyGraphAutosubmitAction(value *CodeSecurityConfiguration_dependency_graph_autosubmit_action)() {
    m.dependency_graph_autosubmit_action = value
}
// SetDependencyGraphAutosubmitActionOptions sets the dependency_graph_autosubmit_action_options property value. Feature options for Automatic dependency submission
func (m *CodeSecurityConfiguration) SetDependencyGraphAutosubmitActionOptions(value CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable)() {
    m.dependency_graph_autosubmit_action_options = value
}
// SetDescription sets the description property value. A description of the code security configuration
func (m *CodeSecurityConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetEnforcement sets the enforcement property value. The enforcement status for a security configuration
func (m *CodeSecurityConfiguration) SetEnforcement(value *CodeSecurityConfiguration_enforcement)() {
    m.enforcement = value
}
// SetHtmlUrl sets the html_url property value. The URL of the configuration
func (m *CodeSecurityConfiguration) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The ID of the code security configuration
func (m *CodeSecurityConfiguration) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the code security configuration. Must be unique within the organization.
func (m *CodeSecurityConfiguration) SetName(value *string)() {
    m.name = value
}
// SetPrivateVulnerabilityReporting sets the private_vulnerability_reporting property value. The enablement status of private vulnerability reporting
func (m *CodeSecurityConfiguration) SetPrivateVulnerabilityReporting(value *CodeSecurityConfiguration_private_vulnerability_reporting)() {
    m.private_vulnerability_reporting = value
}
// SetSecretScanning sets the secret_scanning property value. The enablement status of secret scanning
func (m *CodeSecurityConfiguration) SetSecretScanning(value *CodeSecurityConfiguration_secret_scanning)() {
    m.secret_scanning = value
}
// SetSecretScanningDelegatedBypass sets the secret_scanning_delegated_bypass property value. The enablement status of secret scanning delegated bypass
func (m *CodeSecurityConfiguration) SetSecretScanningDelegatedBypass(value *CodeSecurityConfiguration_secret_scanning_delegated_bypass)() {
    m.secret_scanning_delegated_bypass = value
}
// SetSecretScanningDelegatedBypassOptions sets the secret_scanning_delegated_bypass_options property value. Feature options for secret scanning delegated bypass
func (m *CodeSecurityConfiguration) SetSecretScanningDelegatedBypassOptions(value CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable)() {
    m.secret_scanning_delegated_bypass_options = value
}
// SetSecretScanningNonProviderPatterns sets the secret_scanning_non_provider_patterns property value. The enablement status of secret scanning non-provider patterns
func (m *CodeSecurityConfiguration) SetSecretScanningNonProviderPatterns(value *CodeSecurityConfiguration_secret_scanning_non_provider_patterns)() {
    m.secret_scanning_non_provider_patterns = value
}
// SetSecretScanningPushProtection sets the secret_scanning_push_protection property value. The enablement status of secret scanning push protection
func (m *CodeSecurityConfiguration) SetSecretScanningPushProtection(value *CodeSecurityConfiguration_secret_scanning_push_protection)() {
    m.secret_scanning_push_protection = value
}
// SetSecretScanningValidityChecks sets the secret_scanning_validity_checks property value. The enablement status of secret scanning validity checks
func (m *CodeSecurityConfiguration) SetSecretScanningValidityChecks(value *CodeSecurityConfiguration_secret_scanning_validity_checks)() {
    m.secret_scanning_validity_checks = value
}
// SetTargetType sets the target_type property value. The type of the code security configuration.
func (m *CodeSecurityConfiguration) SetTargetType(value *CodeSecurityConfiguration_target_type)() {
    m.target_type = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CodeSecurityConfiguration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The URL of the configuration
func (m *CodeSecurityConfiguration) SetUrl(value *string)() {
    m.url = value
}
type CodeSecurityConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(*CodeSecurityConfiguration_advanced_security)
    GetCodeScanningDefaultSetup()(*CodeSecurityConfiguration_code_scanning_default_setup)
    GetCodeScanningDefaultSetupOptions()(CodeSecurityConfiguration_code_scanning_default_setup_optionsable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDependabotAlerts()(*CodeSecurityConfiguration_dependabot_alerts)
    GetDependabotSecurityUpdates()(*CodeSecurityConfiguration_dependabot_security_updates)
    GetDependencyGraph()(*CodeSecurityConfiguration_dependency_graph)
    GetDependencyGraphAutosubmitAction()(*CodeSecurityConfiguration_dependency_graph_autosubmit_action)
    GetDependencyGraphAutosubmitActionOptions()(CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable)
    GetDescription()(*string)
    GetEnforcement()(*CodeSecurityConfiguration_enforcement)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetPrivateVulnerabilityReporting()(*CodeSecurityConfiguration_private_vulnerability_reporting)
    GetSecretScanning()(*CodeSecurityConfiguration_secret_scanning)
    GetSecretScanningDelegatedBypass()(*CodeSecurityConfiguration_secret_scanning_delegated_bypass)
    GetSecretScanningDelegatedBypassOptions()(CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable)
    GetSecretScanningNonProviderPatterns()(*CodeSecurityConfiguration_secret_scanning_non_provider_patterns)
    GetSecretScanningPushProtection()(*CodeSecurityConfiguration_secret_scanning_push_protection)
    GetSecretScanningValidityChecks()(*CodeSecurityConfiguration_secret_scanning_validity_checks)
    GetTargetType()(*CodeSecurityConfiguration_target_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAdvancedSecurity(value *CodeSecurityConfiguration_advanced_security)()
    SetCodeScanningDefaultSetup(value *CodeSecurityConfiguration_code_scanning_default_setup)()
    SetCodeScanningDefaultSetupOptions(value CodeSecurityConfiguration_code_scanning_default_setup_optionsable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDependabotAlerts(value *CodeSecurityConfiguration_dependabot_alerts)()
    SetDependabotSecurityUpdates(value *CodeSecurityConfiguration_dependabot_security_updates)()
    SetDependencyGraph(value *CodeSecurityConfiguration_dependency_graph)()
    SetDependencyGraphAutosubmitAction(value *CodeSecurityConfiguration_dependency_graph_autosubmit_action)()
    SetDependencyGraphAutosubmitActionOptions(value CodeSecurityConfiguration_dependency_graph_autosubmit_action_optionsable)()
    SetDescription(value *string)()
    SetEnforcement(value *CodeSecurityConfiguration_enforcement)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetPrivateVulnerabilityReporting(value *CodeSecurityConfiguration_private_vulnerability_reporting)()
    SetSecretScanning(value *CodeSecurityConfiguration_secret_scanning)()
    SetSecretScanningDelegatedBypass(value *CodeSecurityConfiguration_secret_scanning_delegated_bypass)()
    SetSecretScanningDelegatedBypassOptions(value CodeSecurityConfiguration_secret_scanning_delegated_bypass_optionsable)()
    SetSecretScanningNonProviderPatterns(value *CodeSecurityConfiguration_secret_scanning_non_provider_patterns)()
    SetSecretScanningPushProtection(value *CodeSecurityConfiguration_secret_scanning_push_protection)()
    SetSecretScanningValidityChecks(value *CodeSecurityConfiguration_secret_scanning_validity_checks)()
    SetTargetType(value *CodeSecurityConfiguration_target_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
