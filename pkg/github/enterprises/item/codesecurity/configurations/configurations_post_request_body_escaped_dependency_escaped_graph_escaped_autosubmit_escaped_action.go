package configurations
// The enablement status of Automatic dependency submission
type ConfigurationsPostRequestBody_dependency_graph_autosubmit_action int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION ConfigurationsPostRequestBody_dependency_graph_autosubmit_action = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
)

func (i ConfigurationsPostRequestBody_dependency_graph_autosubmit_action) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_dependency_graph_autosubmit_action(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_dependency_graph_autosubmit_action(values []ConfigurationsPostRequestBody_dependency_graph_autosubmit_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_dependency_graph_autosubmit_action) isMultiValue() bool {
    return false
}
