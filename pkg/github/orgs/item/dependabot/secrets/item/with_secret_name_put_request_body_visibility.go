package item
import (
    "errors"
)
// Which type of organization repositories have access to the organization secret. `selected` means only the repositories specified by `selected_repository_ids` can access the secret.
type WithSecret_namePutRequestBody_visibility int

const (
    ALL_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY WithSecret_namePutRequestBody_visibility = iota
    PRIVATE_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
    SELECTED_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
)

func (i WithSecret_namePutRequestBody_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParseWithSecret_namePutRequestBody_visibility(v string) (any, error) {
    result := ALL_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
    switch v {
        case "all":
            result = ALL_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
        case "selected":
            result = SELECTED_WITHSECRET_NAMEPUTREQUESTBODY_VISIBILITY
        default:
            return 0, errors.New("Unknown WithSecret_namePutRequestBody_visibility value: " + v)
    }
    return &result, nil
}
func SerializeWithSecret_namePutRequestBody_visibility(values []WithSecret_namePutRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSecret_namePutRequestBody_visibility) isMultiValue() bool {
    return false
}
