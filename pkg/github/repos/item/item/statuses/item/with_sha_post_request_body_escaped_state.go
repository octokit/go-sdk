package item
// The state of the status.
type WithShaPostRequestBody_state int

const (
    ERROR_WITHSHAPOSTREQUESTBODY_STATE WithShaPostRequestBody_state = iota
    FAILURE_WITHSHAPOSTREQUESTBODY_STATE
    PENDING_WITHSHAPOSTREQUESTBODY_STATE
    SUCCESS_WITHSHAPOSTREQUESTBODY_STATE
)

func (i WithShaPostRequestBody_state) String() string {
    return []string{"error", "failure", "pending", "success"}[i]
}
func ParseWithShaPostRequestBody_state(v string) (any, error) {
    result := ERROR_WITHSHAPOSTREQUESTBODY_STATE
    switch v {
        case "error":
            result = ERROR_WITHSHAPOSTREQUESTBODY_STATE
        case "failure":
            result = FAILURE_WITHSHAPOSTREQUESTBODY_STATE
        case "pending":
            result = PENDING_WITHSHAPOSTREQUESTBODY_STATE
        case "success":
            result = SUCCESS_WITHSHAPOSTREQUESTBODY_STATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithShaPostRequestBody_state(values []WithShaPostRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithShaPostRequestBody_state) isMultiValue() bool {
    return false
}
