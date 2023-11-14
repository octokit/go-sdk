package models
import (
    "errors"
)
// 
type CheckSuite_status int

const (
    QUEUED_CHECKSUITE_STATUS CheckSuite_status = iota
    IN_PROGRESS_CHECKSUITE_STATUS
    COMPLETED_CHECKSUITE_STATUS
)

func (i CheckSuite_status) String() string {
    return []string{"queued", "in_progress", "completed"}[i]
}
func ParseCheckSuite_status(v string) (any, error) {
    result := QUEUED_CHECKSUITE_STATUS
    switch v {
        case "queued":
            result = QUEUED_CHECKSUITE_STATUS
        case "in_progress":
            result = IN_PROGRESS_CHECKSUITE_STATUS
        case "completed":
            result = COMPLETED_CHECKSUITE_STATUS
        default:
            return 0, errors.New("Unknown CheckSuite_status value: " + v)
    }
    return &result, nil
}
func SerializeCheckSuite_status(values []CheckSuite_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CheckSuite_status) isMultiValue() bool {
    return false
}
