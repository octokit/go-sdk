package item
// The type of the bypass reviewer
type WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type int

const (
    TEAM_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type = iota
    ROLE_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
)

func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) String() string {
    return []string{"TEAM", "ROLE"}[i]
}
func ParseWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type(v string) (any, error) {
    result := TEAM_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
    switch v {
        case "TEAM":
            result = TEAM_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
        case "ROLE":
            result = ROLE_WITHCONFIGURATION_PATCHREQUESTBODY_SECRET_SCANNING_DELEGATED_BYPASS_OPTIONS_REVIEWERS_REVIEWER_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type(values []WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_secret_scanning_delegated_bypass_options_reviewers_reviewer_type) isMultiValue() bool {
    return false
}
