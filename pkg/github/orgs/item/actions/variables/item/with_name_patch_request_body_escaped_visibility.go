package item
// The type of repositories in the organization that can access the variable. `selected` means only the repositories specified by `selected_repository_ids` can access the variable.
type WithNamePatchRequestBody_visibility int

const (
    ALL_WITHNAMEPATCHREQUESTBODY_VISIBILITY WithNamePatchRequestBody_visibility = iota
    PRIVATE_WITHNAMEPATCHREQUESTBODY_VISIBILITY
    SELECTED_WITHNAMEPATCHREQUESTBODY_VISIBILITY
)

func (i WithNamePatchRequestBody_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseWithNamePatchRequestBody_visibility(v string) (any, error) {
    result := ALL_WITHNAMEPATCHREQUESTBODY_VISIBILITY
    switch v {
        case "all":
            result = ALL_WITHNAMEPATCHREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_WITHNAMEPATCHREQUESTBODY_VISIBILITY
        case "selected":
            result = SELECTED_WITHNAMEPATCHREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithNamePatchRequestBody_visibility(values []WithNamePatchRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithNamePatchRequestBody_visibility) isMultiValue() bool {
    return false
}
