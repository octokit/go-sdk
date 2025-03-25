package models
type RepositoryRuleMaxFilePathLength_type int

const (
    MAX_FILE_PATH_LENGTH_REPOSITORYRULEMAXFILEPATHLENGTH_TYPE RepositoryRuleMaxFilePathLength_type = iota
)

func (i RepositoryRuleMaxFilePathLength_type) String() string {
    return []string{"max_file_path_length"}[i]
}
func ParseRepositoryRuleMaxFilePathLength_type(v string) (any, error) {
    result := MAX_FILE_PATH_LENGTH_REPOSITORYRULEMAXFILEPATHLENGTH_TYPE
    switch v {
        case "max_file_path_length":
            result = MAX_FILE_PATH_LENGTH_REPOSITORYRULEMAXFILEPATHLENGTH_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMaxFilePathLength_type(values []RepositoryRuleMaxFilePathLength_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMaxFilePathLength_type) isMultiValue() bool {
    return false
}
