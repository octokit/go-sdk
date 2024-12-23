package configurations
// The enablement status of secret scanning non provider patterns
type ConfigurationsPostRequestBody_secret_scanning_non_provider_patterns int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS ConfigurationsPostRequestBody_secret_scanning_non_provider_patterns = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
)

func (i ConfigurationsPostRequestBody_secret_scanning_non_provider_patterns) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_non_provider_patterns(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_NON_PROVIDER_PATTERNS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_non_provider_patterns(values []ConfigurationsPostRequestBody_secret_scanning_non_provider_patterns) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_non_provider_patterns) isMultiValue() bool {
    return false
}
