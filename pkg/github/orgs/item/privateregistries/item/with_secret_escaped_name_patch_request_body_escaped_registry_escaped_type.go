package item
// The registry type.
type WithSecret_namePatchRequestBody_registry_type int

const (
    MAVEN_REPOSITORY_WITHSECRET_NAMEPATCHREQUESTBODY_REGISTRY_TYPE WithSecret_namePatchRequestBody_registry_type = iota
)

func (i WithSecret_namePatchRequestBody_registry_type) String() string {
    return []string{"maven_repository"}[i]
}
func ParseWithSecret_namePatchRequestBody_registry_type(v string) (any, error) {
    result := MAVEN_REPOSITORY_WITHSECRET_NAMEPATCHREQUESTBODY_REGISTRY_TYPE
    switch v {
        case "maven_repository":
            result = MAVEN_REPOSITORY_WITHSECRET_NAMEPATCHREQUESTBODY_REGISTRY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithSecret_namePatchRequestBody_registry_type(values []WithSecret_namePatchRequestBody_registry_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithSecret_namePatchRequestBody_registry_type) isMultiValue() bool {
    return false
}
