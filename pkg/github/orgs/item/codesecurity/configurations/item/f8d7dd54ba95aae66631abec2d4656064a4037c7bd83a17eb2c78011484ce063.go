package item
// The enablement status of code scanning default setup
type WithConfiguration_PatchRequestBody_code_scanning_default_setup int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP WithConfiguration_PatchRequestBody_code_scanning_default_setup = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
)

func (i WithConfiguration_PatchRequestBody_code_scanning_default_setup) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_code_scanning_default_setup(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_CODE_SCANNING_DEFAULT_SETUP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_code_scanning_default_setup(values []WithConfiguration_PatchRequestBody_code_scanning_default_setup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_code_scanning_default_setup) isMultiValue() bool {
    return false
}
