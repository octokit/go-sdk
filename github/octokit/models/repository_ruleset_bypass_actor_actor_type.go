package models
import (
    "errors"
)
// The type of actor that can bypass a ruleset
type RepositoryRulesetBypassActor_actor_type int

const (
    REPOSITORYROLE_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE RepositoryRulesetBypassActor_actor_type = iota
    TEAM_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    INTEGRATION_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    ORGANIZATIONADMIN_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
)

func (i RepositoryRulesetBypassActor_actor_type) String() string {
    return []string{"RepositoryRole", "Team", "Integration", "OrganizationAdmin"}[i]
}
func ParseRepositoryRulesetBypassActor_actor_type(v string) (any, error) {
    result := REPOSITORYROLE_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
    switch v {
        case "RepositoryRole":
            result = REPOSITORYROLE_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "Team":
            result = TEAM_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "Integration":
            result = INTEGRATION_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        case "OrganizationAdmin":
            result = ORGANIZATIONADMIN_REPOSITORYRULESETBYPASSACTOR_ACTOR_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRulesetBypassActor_actor_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRulesetBypassActor_actor_type(values []RepositoryRulesetBypassActor_actor_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRulesetBypassActor_actor_type) isMultiValue() bool {
    return false
}
