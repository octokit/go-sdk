package models
type RepositoryRuleTagNamePattern_type int

const (
    TAG_NAME_PATTERN_REPOSITORYRULETAGNAMEPATTERN_TYPE RepositoryRuleTagNamePattern_type = iota
)

func (i RepositoryRuleTagNamePattern_type) String() string {
    return []string{"tag_name_pattern"}[i]
}
func ParseRepositoryRuleTagNamePattern_type(v string) (any, error) {
    result := TAG_NAME_PATTERN_REPOSITORYRULETAGNAMEPATTERN_TYPE
    switch v {
        case "tag_name_pattern":
            result = TAG_NAME_PATTERN_REPOSITORYRULETAGNAMEPATTERN_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleTagNamePattern_type(values []RepositoryRuleTagNamePattern_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleTagNamePattern_type) isMultiValue() bool {
    return false
}
