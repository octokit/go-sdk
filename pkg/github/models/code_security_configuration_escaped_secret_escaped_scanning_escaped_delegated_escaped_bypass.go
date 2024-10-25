package models
// The enablement status of secret scanning delegated bypass
type CodeSecurityConfiguration_secret_scanning_delegated_bypass int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS CodeSecurityConfiguration_secret_scanning_delegated_bypass = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
)

func (i CodeSecurityConfiguration_secret_scanning_delegated_bypass) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning_delegated_bypass(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_BYPASS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning_delegated_bypass(values []CodeSecurityConfiguration_secret_scanning_delegated_bypass) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning_delegated_bypass) isMultiValue() bool {
    return false
}
