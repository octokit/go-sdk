package models
type RepositoryRuleUpdate_type int

const (
    UPDATE_REPOSITORYRULEUPDATE_TYPE RepositoryRuleUpdate_type = iota
)

func (i RepositoryRuleUpdate_type) String() string {
    return []string{"update"}[i]
}
func ParseRepositoryRuleUpdate_type(v string) (any, error) {
    result := UPDATE_REPOSITORYRULEUPDATE_TYPE
    switch v {
        case "update":
            result = UPDATE_REPOSITORYRULEUPDATE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleUpdate_type(values []RepositoryRuleUpdate_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleUpdate_type) isMultiValue() bool {
    return false
}
