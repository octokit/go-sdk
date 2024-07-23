package item
import (
    "errors"
)
// The target of the ruleset> [!NOTE]> The `push` target is in beta and is subject to change.
type WithRuleset_PutRequestBody_target int

const (
    BRANCH_WITHRULESET_PUTREQUESTBODY_TARGET WithRuleset_PutRequestBody_target = iota
    TAG_WITHRULESET_PUTREQUESTBODY_TARGET
    PUSH_WITHRULESET_PUTREQUESTBODY_TARGET
)

func (i WithRuleset_PutRequestBody_target) String() string {
    return []string{"branch", "tag", "push"}[i]
}
func ParseWithRuleset_PutRequestBody_target(v string) (any, error) {
    result := BRANCH_WITHRULESET_PUTREQUESTBODY_TARGET
    switch v {
        case "branch":
            result = BRANCH_WITHRULESET_PUTREQUESTBODY_TARGET
        case "tag":
            result = TAG_WITHRULESET_PUTREQUESTBODY_TARGET
        case "push":
            result = PUSH_WITHRULESET_PUTREQUESTBODY_TARGET
        default:
            return 0, errors.New("Unknown WithRuleset_PutRequestBody_target value: " + v)
    }
    return &result, nil
}
func SerializeWithRuleset_PutRequestBody_target(values []WithRuleset_PutRequestBody_target) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithRuleset_PutRequestBody_target) isMultiValue() bool {
    return false
}
