package visibility
import (
    "errors"
)
// Denotes whether an email is publicly visible.
type VisibilityPatchRequestBody_visibility int

const (
    PUBLIC_VISIBILITYPATCHREQUESTBODY_VISIBILITY VisibilityPatchRequestBody_visibility = iota
    PRIVATE_VISIBILITYPATCHREQUESTBODY_VISIBILITY
)

func (i VisibilityPatchRequestBody_visibility) String() string {
    return []string{"public", "private"}[i]
}
func ParseVisibilityPatchRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_VISIBILITYPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_VISIBILITYPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_VISIBILITYPATCHREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown VisibilityPatchRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeVisibilityPatchRequestBody_visibility(values []VisibilityPatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VisibilityPatchRequestBody_visibility) isMultiValue() bool {
    return false
}
