package item
// The enablement status of secret scanning delegated bypass
type WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass(values []WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass) isMultiValue() bool {
    return false
}
