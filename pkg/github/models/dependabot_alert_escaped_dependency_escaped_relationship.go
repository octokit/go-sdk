package models
// The vulnerable dependency's relationship to your project.> [!NOTE]> We are rolling out support for dependency relationship across ecosystems. This value will be "unknown" for all dependencies in unsupported ecosystems.
type DependabotAlert_dependency_relationship int

const (
    UNKNOWN_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP DependabotAlert_dependency_relationship = iota
    DIRECT_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
    TRANSITIVE_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
)

func (i DependabotAlert_dependency_relationship) String() string {
    return []string{"unknown", "direct", "transitive"}[i]
}
func ParseDependabotAlert_dependency_relationship(v string) (any, error) {
    result := UNKNOWN_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
    switch v {
        case "unknown":
            result = UNKNOWN_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
        case "direct":
            result = DIRECT_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
        case "transitive":
            result = TRANSITIVE_DEPENDABOTALERT_DEPENDENCY_RELATIONSHIP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDependabotAlert_dependency_relationship(values []DependabotAlert_dependency_relationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlert_dependency_relationship) isMultiValue() bool {
    return false
}
