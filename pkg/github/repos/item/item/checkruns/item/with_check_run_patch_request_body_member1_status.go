package item
import (
    "errors"
)
type WithCheck_run_PatchRequestBodyMember1_status int

const (
    COMPLETED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER1_STATUS WithCheck_run_PatchRequestBodyMember1_status = iota
)

func (i WithCheck_run_PatchRequestBodyMember1_status) String() string {
    return []string{"completed"}[i]
}
func ParseWithCheck_run_PatchRequestBodyMember1_status(v string) (any, error) {
    result := COMPLETED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER1_STATUS
    switch v {
        case "completed":
            result = COMPLETED_WITHCHECK_RUN_PATCHREQUESTBODYMEMBER1_STATUS
        default:
            return 0, errors.New("Unknown WithCheck_run_PatchRequestBodyMember1_status value: " + v)
    }
    return &result, nil
}
func SerializeWithCheck_run_PatchRequestBodyMember1_status(values []WithCheck_run_PatchRequestBodyMember1_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithCheck_run_PatchRequestBodyMember1_status) isMultiValue() bool {
    return false
}
