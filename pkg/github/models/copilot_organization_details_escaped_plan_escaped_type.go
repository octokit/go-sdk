package models
// The Copilot plan of the organization, or the parent enterprise, when applicable.
type CopilotOrganizationDetails_plan_type int

const (
    BUSINESS_COPILOTORGANIZATIONDETAILS_PLAN_TYPE CopilotOrganizationDetails_plan_type = iota
    ENTERPRISE_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
    UNKNOWN_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
)

func (i CopilotOrganizationDetails_plan_type) String() string {
    return []string{"business", "enterprise", "unknown"}[i]
}
func ParseCopilotOrganizationDetails_plan_type(v string) (any, error) {
    result := BUSINESS_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
    switch v {
        case "business":
            result = BUSINESS_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
        case "enterprise":
            result = ENTERPRISE_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
        case "unknown":
            result = UNKNOWN_COPILOTORGANIZATIONDETAILS_PLAN_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCopilotOrganizationDetails_plan_type(values []CopilotOrganizationDetails_plan_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotOrganizationDetails_plan_type) isMultiValue() bool {
    return false
}
