package reviews
import (
    "errors"
)
// The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. By leaving this blank, you set the review action state to `PENDING`, which means you will need to [submit the pull request review](https://docs.github.com/rest/pulls/reviews#submit-a-review-for-a-pull-request) when you are ready.
type ReviewsPostRequestBody_event int

const (
    APPROVE_REVIEWSPOSTREQUESTBODY_EVENT ReviewsPostRequestBody_event = iota
    REQUEST_CHANGES_REVIEWSPOSTREQUESTBODY_EVENT
    COMMENT_REVIEWSPOSTREQUESTBODY_EVENT
)

func (i ReviewsPostRequestBody_event) String() string {
    return []string{"APPROVE", "REQUEST_CHANGES", "COMMENT"}[i]
}
func ParseReviewsPostRequestBody_event(v string) (any, error) {
    result := APPROVE_REVIEWSPOSTREQUESTBODY_EVENT
    switch v {
        case "APPROVE":
            result = APPROVE_REVIEWSPOSTREQUESTBODY_EVENT
        case "REQUEST_CHANGES":
            result = REQUEST_CHANGES_REVIEWSPOSTREQUESTBODY_EVENT
        case "COMMENT":
            result = COMMENT_REVIEWSPOSTREQUESTBODY_EVENT
        default:
            return 0, errors.New("Unknown ReviewsPostRequestBody_event value: " + v)
    }
    return &result, nil
}
func SerializeReviewsPostRequestBody_event(values []ReviewsPostRequestBody_event) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReviewsPostRequestBody_event) isMultiValue() bool {
    return false
}
