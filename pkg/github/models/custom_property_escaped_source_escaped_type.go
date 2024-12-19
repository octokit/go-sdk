package models
// The source type of the property
type CustomProperty_source_type int

const (
    ORGANIZATION_CUSTOMPROPERTY_SOURCE_TYPE CustomProperty_source_type = iota
    ENTERPRISE_CUSTOMPROPERTY_SOURCE_TYPE
)

func (i CustomProperty_source_type) String() string {
    return []string{"organization", "enterprise"}[i]
}
func ParseCustomProperty_source_type(v string) (any, error) {
    result := ORGANIZATION_CUSTOMPROPERTY_SOURCE_TYPE
    switch v {
        case "organization":
            result = ORGANIZATION_CUSTOMPROPERTY_SOURCE_TYPE
        case "enterprise":
            result = ENTERPRISE_CUSTOMPROPERTY_SOURCE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCustomProperty_source_type(values []CustomProperty_source_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomProperty_source_type) isMultiValue() bool {
    return false
}
