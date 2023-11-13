package item
import (
    "errors"
)
// The state that the membership should be in. Only `"active"` will be accepted.
type WithOrgPatchRequestBody_state int

const (
    ACTIVE_WITHORGPATCHREQUESTBODY_STATE WithOrgPatchRequestBody_state = iota
)

func (i WithOrgPatchRequestBody_state) String() string {
    return []string{"active"}[i]
}
func ParseWithOrgPatchRequestBody_state(v string) (any, error) {
    result := ACTIVE_WITHORGPATCHREQUESTBODY_STATE
    switch v {
        case "active":
            result = ACTIVE_WITHORGPATCHREQUESTBODY_STATE
        default:
            return 0, errors.New("Unknown WithOrgPatchRequestBody_state value: " + v)
    }
    return &result, nil
}
func SerializeWithOrgPatchRequestBody_state(values []WithOrgPatchRequestBody_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithOrgPatchRequestBody_state) isMultiValue() bool {
    return false
}
