package item
// The enablement status of secret scanning delegated alert dismissal
type WithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL WithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_ALERT_DISMISSAL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal(values []WithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_alert_dismissal) isMultiValue() bool {
    return false
}
