package models
// The enablement status of secret scanning
type CodeSecurityConfiguration_secret_scanning int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING CodeSecurityConfiguration_secret_scanning = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING
)

func (i CodeSecurityConfiguration_secret_scanning) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning(values []CodeSecurityConfiguration_secret_scanning) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning) isMultiValue() bool {
    return false
}
