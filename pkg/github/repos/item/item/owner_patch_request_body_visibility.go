package item
import (
    "errors"
)
// The visibility of the repository.
type OwnerPatchRequestBody_visibility int

const (
    PUBLIC_OWNERPATCHREQUESTBODY_VISIBILITY OwnerPatchRequestBody_visibility = iota
    PRIVATE_OWNERPATCHREQUESTBODY_VISIBILITY
)

func (i OwnerPatchRequestBody_visibility) String() string {
    return []string{"public", "private"}[i]
}
func ParseOwnerPatchRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_OWNERPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_OWNERPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_OWNERPATCHREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown OwnerPatchRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeOwnerPatchRequestBody_visibility(values []OwnerPatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OwnerPatchRequestBody_visibility) isMultiValue() bool {
    return false
}
