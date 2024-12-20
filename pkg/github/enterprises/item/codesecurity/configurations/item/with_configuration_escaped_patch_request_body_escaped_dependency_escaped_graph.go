package item
// The enablement status of Dependency Graph
type WithConfiguration_PatchRequestBody_dependency_graph int

const (
    ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH WithConfiguration_PatchRequestBody_dependency_graph = iota
    DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
    NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
)

func (i WithConfiguration_PatchRequestBody_dependency_graph) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseWithConfiguration_PatchRequestBody_dependency_graph(v string) (any, error) {
    result := ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
    switch v {
        case "enabled":
            result = ENABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
        case "disabled":
            result = DISABLED_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
        case "not_set":
            result = NOT_SET_WITHCONFIGURATION_PATCHREQUESTBODY_DEPENDENCY_GRAPH
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithConfiguration_PatchRequestBody_dependency_graph(values []WithConfiguration_PatchRequestBody_dependency_graph) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithConfiguration_PatchRequestBody_dependency_graph) isMultiValue() bool {
    return false
}
