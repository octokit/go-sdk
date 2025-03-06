package configurations
// The enablement status of secret scanning delegated alert dismissal
type ConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL ConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
)

func (i ConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal(values []ConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_delegated_alert_dismissal) isMultiValue() bool {
    return false
}
