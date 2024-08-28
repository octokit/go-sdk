package item
// The enablement status of Automatic dependency submission
type WithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION WithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
)

func (i WithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH_AUTOSUBMIT_ACTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action(values []WithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_dependency_graph_autosubmit_action) isMultiValue() bool {
    return false
}
