package statuses
import (
    "errors"
)
// The state of the status. When you set a transient deployment to `inactive`, the deployment will be shown as `destroyed` in GitHub.
type StatusesPostRequestBody_state int

const (
    ERROR_STATUSESPOSTREQUESTBODY_STATE StatusesPostRequestBody_state = iota
    FAILURE_STATUSESPOSTREQUESTBODY_STATE
    INACTIVE_STATUSESPOSTREQUESTBODY_STATE
    IN_PROGRESS_STATUSESPOSTREQUESTBODY_STATE
    QUEUED_STATUSESPOSTREQUESTBODY_STATE
    PENDING_STATUSESPOSTREQUESTBODY_STATE
    SUCCESS_STATUSESPOSTREQUESTBODY_STATE
)

func (i StatusesPostRequestBody_state) String() string {
    return []string{"error", "failure", "inactive", "in_progress", "queued", "pending", "success"}[i]
}
func ParseStatusesPostRequestBody_state(v string) (any, error) {
    result := ERROR_STATUSESPOSTREQUESTBODY_STATE
    switch v {
        case "error":
            result = ERROR_STATUSESPOSTREQUESTBODY_STATE
        case "failure":
            result = FAILURE_STATUSESPOSTREQUESTBODY_STATE
        case "inactive":
            result = INACTIVE_STATUSESPOSTREQUESTBODY_STATE
        case "in_progress":
            result = IN_PROGRESS_STATUSESPOSTREQUESTBODY_STATE
        case "queued":
            result = QUEUED_STATUSESPOSTREQUESTBODY_STATE
        case "pending":
            result = PENDING_STATUSESPOSTREQUESTBODY_STATE
        case "success":
            result = SUCCESS_STATUSESPOSTREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown StatusesPostRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializeStatusesPostRequestBody_state(values []StatusesPostRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i StatusesPostRequestBody_state) isMultiValue() bool {
    return false
}
