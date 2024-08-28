package models
// The enablement status of secret scanning validity checks
type CodeSecurityConfiguration_secret_scanning_validity_checks int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS CodeSecurityConfiguration_secret_scanning_validity_checks = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
)

func (i CodeSecurityConfiguration_secret_scanning_validity_checks) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning_validity_checks(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_VALIDITY_CHECKS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning_validity_checks(values []CodeSecurityConfiguration_secret_scanning_validity_checks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning_validity_checks) isMultiValue() bool {
    return false
}
