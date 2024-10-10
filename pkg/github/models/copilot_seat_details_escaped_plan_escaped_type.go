package models
// The Copilot plan of the organization, or the parent enterprise, when applicable.
type CopilotSeatDetails_plan_type int

const (
    BUSINESS_COPILOTSEATDETAILS_PLAN_TYPE CopilotSeatDetails_plan_type = iota
    ENTERPRISE_COPILOTSEATDETAILS_PLAN_TYPE
    UNKNOWN_COPILOTSEATDETAILS_PLAN_TYPE
)

func (i CopilotSeatDetails_plan_type) String() string {
    return []string{"business", "enterprise", "unknown"}[i]
}
func ParseCopilotSeatDetails_plan_type(v string) (any, error) {
    result := BUSINESS_COPILOTSEATDETAILS_PLAN_TYPE
    switch v {
        case "business":
            result = BUSINESS_COPILOTSEATDETAILS_PLAN_TYPE
        case "enterprise":
            result = ENTERPRISE_COPILOTSEATDETAILS_PLAN_TYPE
        case "unknown":
            result = UNKNOWN_COPILOTSEATDETAILS_PLAN_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCopilotSeatDetails_plan_type(values []CopilotSeatDetails_plan_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CopilotSeatDetails_plan_type) isMultiValue() bool {
    return false
}
