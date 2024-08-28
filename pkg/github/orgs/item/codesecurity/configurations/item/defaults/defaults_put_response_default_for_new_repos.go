package defaults
// Specifies which types of repository this security configuration is applied to by default.
type DefaultsPutResponse_default_for_new_repos int

const (
    ALL_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS DefaultsPutResponse_default_for_new_repos = iota
    NONE_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
    PRIVATE_AND_INTERNAL_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
    PUBLIC_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
)

func (i DefaultsPutResponse_default_for_new_repos) String() string {
    return []string{"all", "none", "private_and_internal", "public"}[i]
}
func ParseDefaultsPutResponse_default_for_new_repos(v string) (any, error) {
    result := ALL_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
    switch v {
        case "all":
            result = ALL_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
        case "none":
            result = NONE_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
        case "private_and_internal":
            result = PRIVATE_AND_INTERNAL_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
        case "public":
            result = PUBLIC_DEFAULTSPUTRESPONSE_DEFAULT_FOR_NEW_REPOS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefaultsPutResponse_default_for_new_repos(values []DefaultsPutResponse_default_for_new_repos) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefaultsPutResponse_default_for_new_repos) isMultiValue() bool {
    return false
}
