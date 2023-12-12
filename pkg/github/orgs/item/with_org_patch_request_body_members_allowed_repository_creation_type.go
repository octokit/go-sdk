package item
import (
    "errors"
)
// Specifies which types of repositories non-admin organization members can create. `private` is only available to repositories that are part of an organization on GitHub Enterprise Cloud. **Note:** This parameter is deprecated and will be removed in the future. Its return value ignores internal repositories. Using this parameter overrides values set in `members_can_create_repositories`. See the parameter deprecation notice in the operation description for details.
type WithOrgPatchRequestBody_members_allowed_repository_creation_type int

const (
    ALL_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE WithOrgPatchRequestBody_members_allowed_repository_creation_type = iota
    PRIVATE_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
    NONE_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
)

func (i WithOrgPatchRequestBody_members_allowed_repository_creation_type) String() string {
    return []string{"all", "private", "none"}[i]
}
func ParseWithOrgPatchRequestBody_members_allowed_repository_creation_type(v string) (any, error) {
    result := ALL_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
    switch v {
        case "all":
            result = ALL_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
        case "private":
            result = PRIVATE_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
        case "none":
            result = NONE_WITHORGPATCHREQUESTBODY_MEMBERS_ALLOWED_REPOSITORY_CREATION_TYPE
        default:
            return 0, errors.New("Unknown WithOrgPatchRequestBody_members_allowed_repository_creation_type value: " + v)
    }
    return &result, nil
}
func SerializeWithOrgPatchRequestBody_members_allowed_repository_creation_type(values []WithOrgPatchRequestBody_members_allowed_repository_creation_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithOrgPatchRequestBody_members_allowed_repository_creation_type) isMultiValue() bool {
    return false
}
