package item
// The enablement status of Dependabot security updates
type WithConfiguration_PatchRequestBody_dependabot_security_updates int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES WithConfiguration_PatchRequestBody_dependabot_security_updates = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
)

func (i WithConfiguration_PatchRequestBody_dependabot_security_updates) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_dependabot_security_updates(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_dependabot_security_updates(values []WithConfiguration_PatchRequestBody_dependabot_security_updates) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_dependabot_security_updates) isMultiValue() bool {
    return false
}
