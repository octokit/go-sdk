package privateregistries
// The registry type.
type PrivateRegistriesPostRequestBody_registry_type int

const (
    MAVEN_REPOSITORY_PRIVATEREGISTRIESPOSTREQUESTBODY_REGISTRY_TYPE PrivateRegistriesPostRequestBody_registry_type = iota
)

func (i PrivateRegistriesPostRequestBody_registry_type) String() string {
    return []string{"maven_repository"}[i]
}
func ParsePrivateRegistriesPostRequestBody_registry_type(v string) (any, error) {
    result := MAVEN_REPOSITORY_PRIVATEREGISTRIESPOSTREQUESTBODY_REGISTRY_TYPE
    switch v {
        case "maven_repository":
            result = MAVEN_REPOSITORY_PRIVATEREGISTRIESPOSTREQUESTBODY_REGISTRY_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePrivateRegistriesPostRequestBody_registry_type(values []PrivateRegistriesPostRequestBody_registry_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivateRegistriesPostRequestBody_registry_type) isMultiValue() bool {
    return false
}
