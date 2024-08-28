package item
// The visibility of the repository.
type RepoPatchRequestBody_visibility int

const (
    PUBLIC_REPOPATCHREQUESTBODY_VISIBILITY RepoPatchRequestBody_visibility = iota
    PRIVATE_REPOPATCHREQUESTBODY_VISIBILITY
)

func (i RepoPatchRequestBody_visibility) String() string {
    return []string{"public", "private"}[i]
}
func ParseRepoPatchRequestBody_visibility(v string) (any, error) {
    result := PUBLIC_REPOPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "public":
            result = PUBLIC_REPOPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_REPOPATCHREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepoPatchRequestBody_visibility(values []RepoPatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepoPatchRequestBody_visibility) isMultiValue() bool {
    return false
}
