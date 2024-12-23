package models
// The type of the value for the property
type CustomPropertySetPayload_value_type int

const (
    STRING_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE CustomPropertySetPayload_value_type = iota
    SINGLE_SELECT_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
    MULTI_SELECT_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
    TRUE_FALSE_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
)

func (i CustomPropertySetPayload_value_type) String() string {
    return []string{"string", "single_select", "multi_select", "true_false"}[i]
}
func ParseCustomPropertySetPayload_value_type(v string) (any, error) {
    result := STRING_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
    switch v {
        case "string":
            result = STRING_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
        case "single_select":
            result = SINGLE_SELECT_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
        case "multi_select":
            result = MULTI_SELECT_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
        case "true_false":
            result = TRUE_FALSE_CUSTOMPROPERTYSETPAYLOAD_VALUE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCustomPropertySetPayload_value_type(values []CustomPropertySetPayload_value_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomPropertySetPayload_value_type) isMultiValue() bool {
    return false
}
