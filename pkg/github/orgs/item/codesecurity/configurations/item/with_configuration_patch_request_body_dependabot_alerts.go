package item
// The enablement status of Dependabot alerts
type WithConfiguration_PatchRequestBody_dependabot_alerts int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS WithConfiguration_PatchRequestBody_dependabot_alerts = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
)

func (i WithConfiguration_PatchRequestBody_dependabot_alerts) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_dependabot_alerts(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_ALERTS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_dependabot_alerts(values []WithConfiguration_PatchRequestBody_dependabot_alerts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_dependabot_alerts) isMultiValue() bool {
    return false
}
