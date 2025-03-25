package models
// The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
type CopilotOrganizationDetails_public_code_suggestions int

const (
    ALLOW_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS CopilotOrganizationDetails_public_code_suggestions = iota
    BLOCK_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
    UNCONFIGURED_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
)

func (i CopilotOrganizationDetails_public_code_suggestions) String() string {
    return []string{"allow", "block", "unconfigured"}[i]
}
func ParseCopilotOrganizationDetails_public_code_suggestions(v string) (any, error) {
    result := ALLOW_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
    switch v {
        case "allow":
            result = ALLOW_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
        case "block":
            result = BLOCK_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
        case "unconfigured":
            result = UNCONFIGURED_COPILOTORGANIZATIONDETAILS_PUBLIC_CODE_SUGGESTIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_public_code_suggestions(values []CopilotOrganizationDetails_public_code_suggestions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_public_code_suggestions) isMultiValue() bool {
    return false
}
