package alerts
type GetScopeQueryParameterType int

const (
    DEVELOPMENT_GETSCOPEQUERYPARAMETERTYPE GetScopeQueryParameterType = iota
    RUNTIME_GETSCOPEQUERYPARAMETERTYPE
)

func (i GetScopeQueryParameterType) String() string {
    return []string{"development", "runtime"}[i]
}
func ParseGetScopeQueryParameterType(v string) (any, error) {
    result := DEVELOPMENT_GETSCOPEQUERYPARAMETERTYPE
    switch v {
        case "development":
            result = DEVELOPMENT_GETSCOPEQUERYPARAMETERTYPE
        case "runtime":
            result = RUNTIME_GETSCOPEQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetScopeQueryParameterType(values []GetScopeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetScopeQueryParameterType) isMultiValue() bool {
    return false
}
