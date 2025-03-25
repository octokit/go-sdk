package models
// The enablement status of Copilot secret scanning
type CodeSecurityConfiguration_secret_scanning_generic_secrets int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS CodeSecurityConfiguration_secret_scanning_generic_secrets = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
)

func (i CodeSecurityConfiguration_secret_scanning_generic_secrets) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning_generic_secrets(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_GENERIC_SECRETS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning_generic_secrets(values []CodeSecurityConfiguration_secret_scanning_generic_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning_generic_secrets) isMultiValue() bool {
    return false
}
