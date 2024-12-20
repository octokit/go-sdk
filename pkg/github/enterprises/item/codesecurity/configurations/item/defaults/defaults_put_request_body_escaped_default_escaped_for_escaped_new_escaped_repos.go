package defaults
// Specify which types of repository this security configuration should be applied to by default.
type DefaultsPutRequestBody_default_for_new_repos int

const (
    ALL_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS DefaultsPutRequestBody_default_for_new_repos = iota
    NONE_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
    PRIVATE_AND_INTERNAL_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
    PUBLIC_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
)

func (i DefaultsPutRequestBody_default_for_new_repos) String() string {
    return []string{"all", "none", "private_and_internal", "public"}[i]
}
func ParseDefaultsPutRequestBody_default_for_new_repos(v string) (any, error) {
    result := ALL_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
    switch v {
        case "all":
            result = ALL_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
        case "none":
            result = NONE_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
        case "private_and_internal":
            result = PRIVATE_AND_INTERNAL_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
        case "public":
            result = PUBLIC_DEFAULTSPUTREQUESTBODY_DEFAULT_FOR_NEW_REPOS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefaultsPutRequestBody_default_for_new_repos(values []DefaultsPutRequestBody_default_for_new_repos) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefaultsPutRequestBody_default_for_new_repos) isMultiValue() bool {
    return false
}
