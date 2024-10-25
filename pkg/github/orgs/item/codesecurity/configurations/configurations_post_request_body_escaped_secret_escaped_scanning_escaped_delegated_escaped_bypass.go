package configurations
// The enablement status of secret scanning delegated bypass
type ConfigurationsPostRequestBody_secret_scanning_delegated_bypass int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS ConfigurationsPostRequestBody_secret_scanning_delegated_bypass = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
)

func (i ConfigurationsPostRequestBody_secret_scanning_delegated_bypass) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_delegated_bypass(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_delegated_bypass(values []ConfigurationsPostRequestBody_secret_scanning_delegated_bypass) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_delegated_bypass) isMultiValue() bool {
    return false
}
