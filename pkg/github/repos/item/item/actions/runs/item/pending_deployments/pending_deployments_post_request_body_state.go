package pending_deployments
import (
    "errors"
)
// Whether to approve or reject deployment to the specified environments.
type Pending_deploymentsPostRequestBody_state int

const (
    APPROVED_PENDING_DEPLOYMENTSPOSTREQUESTBODY_STATE Pending_deploymentsPostRequestBody_state = iota
    REJECTED_PENDING_DEPLOYMENTSPOSTREQUESTBODY_STATE
)

func (i Pending_deploymentsPostRequestBody_state) String() string {
    return []string{"approved", "rejected"}[i]
}
func ParsePending_deploymentsPostRequestBody_state(v string) (any, error) {
    result := APPROVED_PENDING_DEPLOYMENTSPOSTREQUESTBODY_STATE
    switch v {
        case "approved":
            result = APPROVED_PENDING_DEPLOYMENTSPOSTREQUESTBODY_STATE
        case "rejected":
            result = REJECTED_PENDING_DEPLOYMENTSPOSTREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown Pending_deploymentsPostRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializePending_deploymentsPostRequestBody_state(values []Pending_deploymentsPostRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Pending_deploymentsPostRequestBody_state) isMultiValue() bool {
    return false
}
