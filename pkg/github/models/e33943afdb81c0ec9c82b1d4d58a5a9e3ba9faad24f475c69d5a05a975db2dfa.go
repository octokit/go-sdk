package models
// Whether to use labeled runners or standard GitHub runners.
type CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type int

const (
    STANDARD_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type = iota
    LABELED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
    NOT_SET_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
)

func (i CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type) String() string {
    return []string{"standard", "labeled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_code_scanning_default_setup_options_runner_type(v string) (any, error) {
    result := STANDARD_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
    switch v {
        case "standard":
            result = STANDARD_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
        case "labeled":
            result = LABELED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP_OPTIONS_RUNNER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_code_scanning_default_setup_options_runner_type(values []CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_code_scanning_default_setup_options_runner_type) isMultiValue() bool {
    return false
}
