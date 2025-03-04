package item
// The enablement status of Copilot secret scanning
type WithConfiguration_PatchRequestBody_secret_scanning_generic_secrets int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS WithConfiguration_PatchRequestBody_secret_scanning_generic_secrets = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_generic_secrets) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_generic_secrets(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_generic_secrets(values []WithConfiguration_PatchRequestBody_secret_scanning_generic_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_generic_secrets) isMultiValue() bool {
    return false
}
