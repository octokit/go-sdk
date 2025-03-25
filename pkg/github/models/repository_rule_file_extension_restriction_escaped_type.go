package models
type RepositoryRuleFileExtensionRestriction_type int

const (
    FILE_EXTENSION_RESTRICTION_REPOSITORYRULEFILEEXTENSIONRESTRICTION_TYPE RepositoryRuleFileExtensionRestriction_type = iota
)

func (i RepositoryRuleFileExtensionRestriction_type) String() string {
    return []string{"file_extension_restriction"}[i]
}
func ParseRepositoryRuleFileExtensionRestriction_type(v string) (any, error) {
    result := FILE_EXTENSION_RESTRICTION_REPOSITORYRULEFILEEXTENSIONRESTRICTION_TYPE
    switch v {
        case "file_extension_restriction":
            result = FILE_EXTENSION_RESTRICTION_REPOSITORYRULEFILEEXTENSIONRESTRICTION_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleFileExtensionRestriction_type(values []RepositoryRuleFileExtensionRestriction_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleFileExtensionRestriction_type) isMultiValue() bool {
    return false
}
