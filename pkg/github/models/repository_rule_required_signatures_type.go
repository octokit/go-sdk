package models
type RepositoryRuleRequiredSignatures_type int

const (
    REQUIRED_SIGNATURES_REPOSITORYRULEREQUIREDSIGNATURES_TYPE RepositoryRuleRequiredSignatures_type = iota
)

func (i RepositoryRuleRequiredSignatures_type) String() string {
    return []string{"required_signatures"}[i]
}
func ParseRepositoryRuleRequiredSignatures_type(v string) (any, error) {
    result := REQUIRED_SIGNATURES_REPOSITORYRULEREQUIREDSIGNATURES_TYPE
    switch v {
        case "required_signatures":
            result = REQUIRED_SIGNATURES_REPOSITORYRULEREQUIREDSIGNATURES_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleRequiredSignatures_type(values []RepositoryRuleRequiredSignatures_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleRequiredSignatures_type) isMultiValue() bool {
    return false
}
