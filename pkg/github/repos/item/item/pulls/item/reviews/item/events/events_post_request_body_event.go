package events
// The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. When you leave this blank, the API returns _HTTP 422 (Unrecognizable entity)_ and sets the review action state to `PENDING`, which means you will need to re-submit the pull request review using a review action.
type EventsPostRequestBody_event int

const (
    APPROVE_EVENTSPOSTREQUESTBODY_EVENT EventsPostRequestBody_event = iota
    REQUEST_CHANGES_EVENTSPOSTREQUESTBODY_EVENT
    COMMENT_EVENTSPOSTREQUESTBODY_EVENT
)

func (i EventsPostRequestBody_event) String() string {
    return []string{"APPROVE", "REQUEST_CHANGES", "COMMENT"}[i]
}
func ParseEventsPostRequestBody_event(v string) (any, error) {
    result := APPROVE_EVENTSPOSTREQUESTBODY_EVENT
    switch v {
        case "APPROVE":
            result = APPROVE_EVENTSPOSTREQUESTBODY_EVENT
        case "REQUEST_CHANGES":
            result = REQUEST_CHANGES_EVENTSPOSTREQUESTBODY_EVENT
        case "COMMENT":
            result = COMMENT_EVENTSPOSTREQUESTBODY_EVENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEventsPostRequestBody_event(values []EventsPostRequestBody_event) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EventsPostRequestBody_event) isMultiValue() bool {
    return false
}
