package rulesets
// The target of the ruleset
type RulesetsPostRequestBody_target int

const (
    BRANCH_RULESETSPOSTREQUESTBODY_TARGET RulesetsPostRequestBody_target = iota
    TAG_RULESETSPOSTREQUESTBODY_TARGET
    PUSH_RULESETSPOSTREQUESTBODY_TARGET
    REPOSITORY_RULESETSPOSTREQUESTBODY_TARGET
)

func (i RulesetsPostRequestBody_target) String() string {
    return []string{"branch", "tag", "push", "repository"}[i]
}
func ParseRulesetsPostRequestBody_target(v string) (any, error) {
    result := BRANCH_RULESETSPOSTREQUESTBODY_TARGET
    switch v {
        case "branch":
            result = BRANCH_RULESETSPOSTREQUESTBODY_TARGET
        case "tag":
            result = TAG_RULESETSPOSTREQUESTBODY_TARGET
        case "push":
            result = PUSH_RULESETSPOSTREQUESTBODY_TARGET
        case "repository":
            result = REPOSITORY_RULESETSPOSTREQUESTBODY_TARGET
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRulesetsPostRequestBody_target(values []RulesetsPostRequestBody_target) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RulesetsPostRequestBody_target) isMultiValue() bool {
    return false
}
