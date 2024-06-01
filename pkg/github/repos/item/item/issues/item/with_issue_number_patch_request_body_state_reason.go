package item
import (
    "errors"
)
// The reason for the state change. Ignored unless `state` is changed.
type WithIssue_numberPatchRequestBody_state_reason int

const (
    COMPLETED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON WithIssue_numberPatchRequestBody_state_reason = iota
    NOT_PLANNED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
    REOPENED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
)

func (i WithIssue_numberPatchRequestBody_state_reason) String() string {
    return []string{"completed", "not_planned", "reopened"}[i]
}
func ParseWithIssue_numberPatchRequestBody_state_reason(v string) (any, error) {
    result := COMPLETED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
    switch v {
        case "completed":
            result = COMPLETED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
        case "not_planned":
            result = NOT_PLANNED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
        case "reopened":
            result = REOPENED_WITHISSUE_NUMBERPATCHREQUESTBODY_STATE_REASON
        default:
            return 0, errors.New("Unknown WithIssue_numberPatchRequestBody_state_reason value: " + v)
    }
    return &result, nil
}
func SerializeWithIssue_numberPatchRequestBody_state_reason(values []WithIssue_numberPatchRequestBody_state_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithIssue_numberPatchRequestBody_state_reason) isMultiValue() bool {
    return false
}
