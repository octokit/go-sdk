package models
type RepositoryRuleFilePathRestriction_type int

const (
    FILE_PATH_RESTRICTION_REPOSITORYRULEFILEPATHRESTRICTION_TYPE RepositoryRuleFilePathRestriction_type = iota
)

func (i RepositoryRuleFilePathRestriction_type) String() string {
    return []string{"file_path_restriction"}[i]
}
func ParseRepositoryRuleFilePathRestriction_type(v string) (any, error) {
    result := FILE_PATH_RESTRICTION_REPOSITORYRULEFILEPATHRESTRICTION_TYPE
    switch v {
        case "file_path_restriction":
            result = FILE_PATH_RESTRICTION_REPOSITORYRULEFILEPATHRESTRICTION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleFilePathRestriction_type(values []RepositoryRuleFilePathRestriction_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleFilePathRestriction_type) isMultiValue() bool {
    return false
}
