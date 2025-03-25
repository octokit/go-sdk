package models
// Color for the issue type.
type OrganizationCreateIssueType_color int

const (
    GRAY_ORGANIZATIONCREATEISSUETYPE_COLOR OrganizationCreateIssueType_color = iota
    BLUE_ORGANIZATIONCREATEISSUETYPE_COLOR
    GREEN_ORGANIZATIONCREATEISSUETYPE_COLOR
    YELLOW_ORGANIZATIONCREATEISSUETYPE_COLOR
    ORANGE_ORGANIZATIONCREATEISSUETYPE_COLOR
    RED_ORGANIZATIONCREATEISSUETYPE_COLOR
    PINK_ORGANIZATIONCREATEISSUETYPE_COLOR
    PURPLE_ORGANIZATIONCREATEISSUETYPE_COLOR
)

func (i OrganizationCreateIssueType_color) String() string {
    return []string{"gray", "blue", "green", "yellow", "orange", "red", "pink", "purple"}[i]
}
func ParseOrganizationCreateIssueType_color(v string) (any, error) {
    result := GRAY_ORGANIZATIONCREATEISSUETYPE_COLOR
    switch v {
        case "gray":
            result = GRAY_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "blue":
            result = BLUE_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "green":
            result = GREEN_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "yellow":
            result = YELLOW_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "orange":
            result = ORANGE_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "red":
            result = RED_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "pink":
            result = PINK_ORGANIZATIONCREATEISSUETYPE_COLOR
        case "purple":
            result = PURPLE_ORGANIZATIONCREATEISSUETYPE_COLOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationCreateIssueType_color(values []OrganizationCreateIssueType_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationCreateIssueType_color) isMultiValue() bool {
    return false
}
