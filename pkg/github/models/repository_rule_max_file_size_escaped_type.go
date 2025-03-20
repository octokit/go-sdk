package models
type RepositoryRuleMaxFileSize_type int

const (
    MAX_FILE_SIZE_REPOSITORYRULEMAXFILESIZE_TYPE RepositoryRuleMaxFileSize_type = iota
)

func (i RepositoryRuleMaxFileSize_type) String() string {
    return []string{"max_file_size"}[i]
}
func ParseRepositoryRuleMaxFileSize_type(v string) (any, error) {
    result := MAX_FILE_SIZE_REPOSITORYRULEMAXFILESIZE_TYPE
    switch v {
        case "max_file_size":
            result = MAX_FILE_SIZE_REPOSITORYRULEMAXFILESIZE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMaxFileSize_type(values []RepositoryRuleMaxFileSize_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMaxFileSize_type) isMultiValue() bool {
    return false
}
