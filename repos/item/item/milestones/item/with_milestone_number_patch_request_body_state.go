package item
import (
    "errors"
)
// The state of the milestone. Either `open` or `closed`.
type WithMilestone_numberPatchRequestBody_state int

const (
    OPEN_WITHMILESTONE_NUMBERPATCHREQUESTBODY_STATE WithMilestone_numberPatchRequestBody_state = iota
    CLOSED_WITHMILESTONE_NUMBERPATCHREQUESTBODY_STATE
)

func (i WithMilestone_numberPatchRequestBody_state) String() string {
    return []string{"open", "closed"}[i]
}
func ParseWithMilestone_numberPatchRequestBody_state(v string) (any, error) {
    result := OPEN_WITHMILESTONE_NUMBERPATCHREQUESTBODY_STATE
    switch v {
        case "open":
            result = OPEN_WITHMILESTONE_NUMBERPATCHREQUESTBODY_STATE
        case "closed":
            result = CLOSED_WITHMILESTONE_NUMBERPATCHREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown WithMilestone_numberPatchRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializeWithMilestone_numberPatchRequestBody_state(values []WithMilestone_numberPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithMilestone_numberPatchRequestBody_state) isMultiValue() bool {
    return false
}
