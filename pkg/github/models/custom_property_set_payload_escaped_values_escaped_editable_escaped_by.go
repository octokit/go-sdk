package models
// Who can edit the values of the property
type CustomPropertySetPayload_values_editable_by int

const (
    ORG_ACTORS_CUSTOMPROPERTYSETPAYLOAD_VALUES_EDITABLE_BY CustomPropertySetPayload_values_editable_by = iota
    ORG_AND_REPO_ACTORS_CUSTOMPROPERTYSETPAYLOAD_VALUES_EDITABLE_BY
)

func (i CustomPropertySetPayload_values_editable_by) String() string {
    return []string{"org_actors", "org_and_repo_actors"}[i]
}
func ParseCustomPropertySetPayload_values_editable_by(v string) (any, error) {
    result := ORG_ACTORS_CUSTOMPROPERTYSETPAYLOAD_VALUES_EDITABLE_BY
    switch v {
        case "org_actors":
            result = ORG_ACTORS_CUSTOMPROPERTYSETPAYLOAD_VALUES_EDITABLE_BY
        case "org_and_repo_actors":
            result = ORG_AND_REPO_ACTORS_CUSTOMPROPERTYSETPAYLOAD_VALUES_EDITABLE_BY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCustomPropertySetPayload_values_editable_by(values []CustomPropertySetPayload_values_editable_by) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomPropertySetPayload_values_editable_by) isMultiValue() bool {
    return false
}
