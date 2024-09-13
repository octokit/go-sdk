package item
// The enablement status of GitHub Advanced Security
type WithConfiguration_PatchRequestBody_advanced_security int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_ADVANCED_SECURITY WithConfiguration_PatchRequestBody_advanced_security = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_ADVANCED_SECURITY
)

func (i WithConfiguration_PatchRequestBody_advanced_security) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseWithConfiguration_PatchRequestBody_advanced_security(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_ADVANCED_SECURITY
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_ADVANCED_SECURITY
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_ADVANCED_SECURITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_advanced_security(values []WithConfiguration_PatchRequestBody_advanced_security) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_advanced_security) isMultiValue() bool {
    return false
}
