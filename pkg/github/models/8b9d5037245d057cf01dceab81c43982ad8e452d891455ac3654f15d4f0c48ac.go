package models
// The enablement status of secret scanning delegated alert dismissal
type CodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal int

const (
    ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL CodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal = iota
    DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
)

func (i CodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal(values []CodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_secret_scanning_delegated_alert_dismissal) isMultiValue() bool {
    return false
}
