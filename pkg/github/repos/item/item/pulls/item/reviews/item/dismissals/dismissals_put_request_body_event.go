package dismissals
type DismissalsPutRequestBody_event int

const (
    DISMISS_DISMISSALSPUTREQUESTBODY_EVENT DismissalsPutRequestBody_event = iota
)

func (i DismissalsPutRequestBody_event) String() string {
    return []string{"DISMISS"}[i]
}
func ParseDismissalsPutRequestBody_event(v string) (any, error) {
    result := DISMISS_DISMISSALSPUTREQUESTBODY_EVENT
    switch v {
        case "DISMISS":
            result = DISMISS_DISMISSALSPUTREQUESTBODY_EVENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDismissalsPutRequestBody_event(values []DismissalsPutRequestBody_event) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DismissalsPutRequestBody_event) isMultiValue() bool {
    return false
}
