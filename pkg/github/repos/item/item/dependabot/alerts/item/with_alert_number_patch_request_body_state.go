package item
import (
    "errors"
)
// The state of the Dependabot alert.A `dismissed_reason` must be provided when setting the state to `dismissed`.
type WithAlert_numberPatchRequestBody_state int

const (
    DISMISSED_WITHALERT_NUMBERPATCHREQUESTBODY_STATE WithAlert_numberPatchRequestBody_state = iota
    OPEN_WITHALERT_NUMBERPATCHREQUESTBODY_STATE
)

func (i WithAlert_numberPatchRequestBody_state) String() string {
    return []string{"dismissed", "open"}[i]
}
func ParseWithAlert_numberPatchRequestBody_state(v string) (any, error) {
    result := DISMISSED_WITHALERT_NUMBERPATCHREQUESTBODY_STATE
    switch v {
        case "dismissed":
            result = DISMISSED_WITHALERT_NUMBERPATCHREQUESTBODY_STATE
        case "open":
            result = OPEN_WITHALERT_NUMBERPATCHREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown WithAlert_numberPatchRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializeWithAlert_numberPatchRequestBody_state(values []WithAlert_numberPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithAlert_numberPatchRequestBody_state) isMultiValue() bool {
    return false
}
