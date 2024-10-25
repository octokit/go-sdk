package configurations
// The type of the bypass reviewer
type ConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type int

const (
    TEAM_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE ConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type = iota
    ROLE_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
)

func (i ConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) String() string {
    return []string{"TEAM", "ROLE"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type(v string) (any, error) {
    result := TEAM_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
    switch v {
        case "TEAM":
            result = TEAM_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
        case "ROLE":
            result = ROLE_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type(values []ConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) isMultiValue() bool {
    return false
}
