package models
// The enablement status of Automatic dependency submission
type CodeSecurityConfiguration_dependency_graph_autosubmit_action int

const (
    ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION CodeSecurityConfiguration_dependency_graph_autosubmit_action = iota
    DISABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    NOT_SET_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
)

func (i CodeSecurityConfiguration_dependency_graph_autosubmit_action) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_dependency_graph_autosubmit_action(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_dependency_graph_autosubmit_action(values []CodeSecurityConfiguration_dependency_graph_autosubmit_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_dependency_graph_autosubmit_action) isMultiValue() bool {
    return false
}
