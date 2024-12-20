package privateregistries
// Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
type PrivateRegistriesPostRequestBody_visibility int

const (
    ALL_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY PrivateRegistriesPostRequestBody_visibility = iota
    PRIVATE_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
    SELECTED_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
)

func (i PrivateRegistriesPostRequestBody_visibility) String() string {
    return []string{"all", "private", "selected"}[i]
}
func ParsePrivateRegistriesPostRequestBody_visibility(v string) (any, error) {
    result := ALL_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
    switch v {
        case "all":
            result = ALL_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
        case "private":
            result = PRIVATE_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
        case "selected":
            result = SELECTED_PRIVATEREGISTRIESPOSTREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePrivateRegistriesPostRequestBody_visibility(values []PrivateRegistriesPostRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivateRegistriesPostRequestBody_visibility) isMultiValue() bool {
    return false
}
