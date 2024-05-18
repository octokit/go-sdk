package access
// Which users can access codespaces in the organization. `disabled` means that no users can access codespaces in the organization.
type AccessPutRequestBody_visibility int

const (
    DISABLED_ACCESSPUTREQUESTBODY_VISIBILITY AccessPutRequestBody_visibility = iota
    SELECTED_MEMBERS_ACCESSPUTREQUESTBODY_VISIBILITY
    ALL_MEMBERS_ACCESSPUTREQUESTBODY_VISIBILITY
    ALL_MEMBERS_AND_OUTSIDE_COLLABORATORS_ACCESSPUTREQUESTBODY_VISIBILITY
)

func (i AccessPutRequestBody_visibility) String() string {
    return []string{"disabled", "selected_members", "all_members", "all_members_and_outside_collaborators"}[i]
}
func ParseAccessPutRequestBody_visibility(v string) (any, error) {
    result := DISABLED_ACCESSPUTREQUESTBODY_VISIBILITY
    switch v {
        case "disabled":
            result = DISABLED_ACCESSPUTREQUESTBODY_VISIBILITY
        case "selected_members":
            result = SELECTED_MEMBERS_ACCESSPUTREQUESTBODY_VISIBILITY
        case "all_members":
            result = ALL_MEMBERS_ACCESSPUTREQUESTBODY_VISIBILITY
        case "all_members_and_outside_collaborators":
            result = ALL_MEMBERS_AND_OUTSIDE_COLLABORATORS_ACCESSPUTREQUESTBODY_VISIBILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAccessPutRequestBody_visibility(values []AccessPutRequestBody_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccessPutRequestBody_visibility) isMultiValue() bool {
    return false
}
