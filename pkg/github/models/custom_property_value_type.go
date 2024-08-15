package models
import (
    "errors"
)
// The type of the value for the property
type CustomProperty_value_type int

const (
    STRING_CUSTOMPROPERTY_VALUE_TYPE CustomProperty_value_type = iota
    SINGLE_SELECT_CUSTOMPROPERTY_VALUE_TYPE
    MULTI_SELECT_CUSTOMPROPERTY_VALUE_TYPE
    TRUE_FALSE_CUSTOMPROPERTY_VALUE_TYPE
)

func (i CustomProperty_value_type) String() string {
    return []string{"string", "single_select", "multi_select", "true_false"}[i]
}
func ParseCustomProperty_value_type(v string) (any, error) {
    result := STRING_CUSTOMPROPERTY_VALUE_TYPE
    switch v {
        case "string":
            result = STRING_CUSTOMPROPERTY_VALUE_TYPE
        case "single_select":
            result = SINGLE_SELECT_CUSTOMPROPERTY_VALUE_TYPE
        case "multi_select":
            result = MULTI_SELECT_CUSTOMPROPERTY_VALUE_TYPE
        case "true_false":
            result = TRUE_FALSE_CUSTOMPROPERTY_VALUE_TYPE
        default:
            return 0, errors.New("Unknown CustomProperty_value_type value: " + v)
    }
    return &result, nil
}
func SerializeCustomProperty_value_type(values []CustomProperty_value_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomProperty_value_type) isMultiValue() bool {
    return false
}
