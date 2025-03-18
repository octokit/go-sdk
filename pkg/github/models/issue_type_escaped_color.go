package models
// The color of the issue type.
type IssueType_color int

const (
    GRAY_ISSUETYPE_COLOR IssueType_color = iota
    BLUE_ISSUETYPE_COLOR
    GREEN_ISSUETYPE_COLOR
    YELLOW_ISSUETYPE_COLOR
    ORANGE_ISSUETYPE_COLOR
    RED_ISSUETYPE_COLOR
    PINK_ISSUETYPE_COLOR
    PURPLE_ISSUETYPE_COLOR
)

func (i IssueType_color) String() string {
    return []string{"gray", "blue", "green", "yellow", "orange", "red", "pink", "purple"}[i]
}
func ParseIssueType_color(v string) (any, error) {
    result := GRAY_ISSUETYPE_COLOR
    switch v {
        case "gray":
            result = GRAY_ISSUETYPE_COLOR
        case "blue":
            result = BLUE_ISSUETYPE_COLOR
        case "green":
            result = GREEN_ISSUETYPE_COLOR
        case "yellow":
            result = YELLOW_ISSUETYPE_COLOR
        case "orange":
            result = ORANGE_ISSUETYPE_COLOR
        case "red":
            result = RED_ISSUETYPE_COLOR
        case "pink":
            result = PINK_ISSUETYPE_COLOR
        case "purple":
            result = PURPLE_ISSUETYPE_COLOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIssueType_color(values []IssueType_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IssueType_color) isMultiValue() bool {
    return false
}
