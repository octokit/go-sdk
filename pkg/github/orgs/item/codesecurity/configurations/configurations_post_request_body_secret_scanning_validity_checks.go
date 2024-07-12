package configurations
import (
    "errors"
)
// The enablement status of secret scanning validity checks
type ConfigurationsPostRequestBody_secret_scanning_validity_checks int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS ConfigurationsPostRequestBody_secret_scanning_validity_checks = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
)

func (i ConfigurationsPostRequestBody_secret_scanning_validity_checks) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_validity_checks(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_VALIDITY_CHECKS
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_secret_scanning_validity_checks value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_validity_checks(values []ConfigurationsPostRequestBody_secret_scanning_validity_checks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_validity_checks) isMultiValue() bool {
    return false
}
