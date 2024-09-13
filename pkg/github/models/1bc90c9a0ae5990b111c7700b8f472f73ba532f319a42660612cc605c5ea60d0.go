package models
// The enablement status of secret scanning non-provider patterns
type CodeSecurityConfiguration_secret_scanning_non_provider_patterns int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS CodeSecurityConfiguration_secret_scanning_non_provider_patterns = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
)

func (i CodeSecurityConfiguration_secret_scanning_non_provider_patterns) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning_non_provider_patterns(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning_non_provider_patterns(values []CodeSecurityConfiguration_secret_scanning_non_provider_patterns) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning_non_provider_patterns) isMultiValue() bool {
    return false
}
