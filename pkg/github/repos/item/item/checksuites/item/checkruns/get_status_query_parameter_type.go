package checkruns
type GetStatusQueryParameterType int

const (
    QUEUED_GETSTATUSQUERYPARAMETERTYPE GetStatusQueryParameterType = iota
    IN_PROGRESS_GETSTATUSQUERYPARAMETERTYPE
    COMPLETED_GETSTATUSQUERYPARAMETERTYPE
)

func (i GetStatusQueryParameterType) String() string {
    return []string{"queued", "in_progress", "completed"}[i]
}
func ParseGetStatusQueryParameterType(v string) (any, error) {
    result := QUEUED_GETSTATUSQUERYPARAMETERTYPE
    switch v {
        case "queued":
            result = QUEUED_GETSTATUSQUERYPARAMETERTYPE
        case "in_progress":
            result = IN_PROGRESS_GETSTATUSQUERYPARAMETERTYPE
        case "completed":
            result = COMPLETED_GETSTATUSQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetStatusQueryParameterType(values []GetStatusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStatusQueryParameterType) isMultiValue() bool {
    return false
}
