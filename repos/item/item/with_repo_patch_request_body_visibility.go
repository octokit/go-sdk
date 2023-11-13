package item
import (
    "errors"
)
// The visibility of the repository.
type WithRepoPatchRequestBody_visibility int

const (
    PUBLIC_WITHREPOPATCHREQUESTBODY_VISIBILITY WithRepoPatchRequestBody_visibility = iota
    PRIVATE_WITHREPOPATCHREQUESTBODY_VISIBILITY
)

func (i WithRepoPatchRequestBody_visibility) String() string {
    return []string{"public", "private"}[i]
}
func ParseWithRepoPatchRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_WITHREPOPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_WITHREPOPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_WITHREPOPATCHREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown WithRepoPatchRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeWithRepoPatchRequestBody_visibility(values []WithRepoPatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithRepoPatchRequestBody_visibility) isMultiValue() bool {
    return false
}
