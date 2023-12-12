package item
import (
    "errors"
)
// 
type WithCheck_run_PatchRequestBodyMember2_status int

const (
    QUEUED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER2_STATUS WithCheck_run_PatchRequestBodyMember2_status = iota
    IN_PROGRESS_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER2_STATUS
)

func (i WithCheck_run_PatchRequestBodyMember2_status) String() string {
    return []string{"queued", "in_progress"}[i]
}
func ParseWithCheck_run_PatchRequestBodyMember2_status(v string) (any, error) {
    result := QUEUED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER2_STATUS
    switch v {
        case "queued":
            result = QUEUED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER2_STATUS
        case "in_progress":
            result = IN_PROGRESS_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER2_STATUS
        default:
            return 0, errors.New("Unknown WithCheck_run_PatchRequestBodyMember2_status value: " + v)
    }
    return &result, nil
}
func SerializeWithCheck_run_PatchRequestBodyMember2_status(values []WithCheck_run_PatchRequestBodyMember2_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithCheck_run_PatchRequestBodyMember2_status) isMultiValue() bool {
    return false
}
