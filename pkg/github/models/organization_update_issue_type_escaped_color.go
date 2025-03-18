package models
// Color for the issue type.
type OrganizationUpdateIssueType_color int

const (
    GRAY_ORGANIZATIONUPDATEISSUETYPE_COLOR OrganizationUpdateIssueType_color = iota
    BLUE_ORGANIZATIONUPDATEISSUETYPE_COLOR
    GREEN_ORGANIZATIONUPDATEISSUETYPE_COLOR
    YELLOW_ORGANIZATIONUPDATEISSUETYPE_COLOR
    ORANGE_ORGANIZATIONUPDATEISSUETYPE_COLOR
    RED_ORGANIZATIONUPDATEISSUETYPE_COLOR
    PINK_ORGANIZATIONUPDATEISSUETYPE_COLOR
    PURPLE_ORGANIZATIONUPDATEISSUETYPE_COLOR
)

func (i OrganizationUpdateIssueType_color) String() string {
    return []string{"gray", "blue", "green", "yellow", "orange", "red", "pink", "purple"}[i]
}
func ParseOrganizationUpdateIssueType_color(v string) (any, error) {
    result := GRAY_ORGANIZATIONUPDATEISSUETYPE_COLOR
    switch v {
        case "gray":
            result = GRAY_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "blue":
            result = BLUE_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "green":
            result = GREEN_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "yellow":
            result = YELLOW_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "orange":
            result = ORANGE_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "red":
            result = RED_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "pink":
            result = PINK_ORGANIZATIONUPDATEISSUETYPE_COLOR
        case "purple":
            result = PURPLE_ORGANIZATIONUPDATEISSUETYPE_COLOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrganizationUpdateIssueType_color(values []OrganizationUpdateIssueType_color) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationUpdateIssueType_color) isMultiValue() bool {
    return false
}
