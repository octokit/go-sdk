package configurations
// The enablement status of Copilot secret scanning
type ConfigurationsPostRequestBody_secret_scanning_generic_secrets int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS ConfigurationsPostRequestBody_secret_scanning_generic_secrets = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
)

func (i ConfigurationsPostRequestBody_secret_scanning_generic_secrets) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_generic_secrets(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_GENERIC_SECRETS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_generic_secrets(values []ConfigurationsPostRequestBody_secret_scanning_generic_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_generic_secrets) isMultiValue() bool {
    return false
}
