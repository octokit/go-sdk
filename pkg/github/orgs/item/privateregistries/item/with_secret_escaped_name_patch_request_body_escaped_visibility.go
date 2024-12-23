package item
// Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
type WithSecret_namePatchRequestBody_visibility int

const (
    ALL_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY WithSecret_namePatchRequestBody_visibility = iota
    PRIVATE_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
    SELECTED_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
)

func (i WithSecret_namePatchRequestBody_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseWithSecret_namePatchRequestBody_visibility(v string) (any, error) {
    result := ALL_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "all":
            result = ALL_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
        case "selected":
            result = SELECTED_WITHSECRET_NAMEPATCHREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSecret_namePatchRequestBody_visibility(values []WithSecret_namePatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSecret_namePatchRequestBody_visibility) isMultiValue() bool {
    return false
}
