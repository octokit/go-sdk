package attach
import (
    "errors"
)
// The type of repositories to attach the configuration to. `selected` means the configuration will be attached to only the repositories specified by `selected_repository_ids`
type AttachPostRequestBody_scope int

const (
    ALL_ATTACHPOSTREQUESTBODY_SCOPE AttachPostRequestBody_scope = iota
    PUBLIC_ATTACHPOSTREQUESTBODY_SCOPE
    PRIVATE_OR_INTERNAL_ATTACHPOSTREQUESTBODY_SCOPE
    SELECTED_ATTACHPOSTREQUESTBODY_SCOPE
)

func (i AttachPostRequestBody_scope) String() string {
    return []string{"all", "public", "private_or_internal", "selected"}[i]
}
func ParseAttachPostRequestBody_scope(v string) (any, error) {
    result := ALL_ATTACHPOSTREQUESTBODY_SCOPE
    switch v {
        case "all":
            result = ALL_ATTACHPOSTREQUESTBODY_SCOPE
        case "public":
            result = PUBLIC_ATTACHPOSTREQUESTBODY_SCOPE
        case "private_or_internal":
            result = PRIVATE_OR_INTERNAL_ATTACHPOSTREQUESTBODY_SCOPE
        case "selected":
            result = SELECTED_ATTACHPOSTREQUESTBODY_SCOPE
        default:
            return 0, errors.New("Unknown AttachPostRequestBody_scope value: " + v)
    }
    return &result, nil
}
func SerializeAttachPostRequestBody_scope(values []AttachPostRequestBody_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AttachPostRequestBody_scope) isMultiValue() bool {
    return false
}
