package teams
// The notification setting the team has chosen. The options are:   * `notifications_enabled` - team members receive notifications when the team is @mentioned.   * `notifications_disabled` - no one receives notifications.  Default: `notifications_enabled`
type TeamsPostRequestBody_notification_setting int

const (
    NOTIFICATIONS_ENABLED_TEAMSPOSTREQUESTBODY_NOTIFICATION_SETTING TeamsPostRequestBody_notification_setting = iota
    NOTIFICATIONS_DISABLED_TEAMSPOSTREQUESTBODY_NOTIFICATION_SETTING
)

func (i TeamsPostRequestBody_notification_setting) String() string {
    return []string{"notifications_enabled", "notifications_disabled"}[i]
}
func ParseTeamsPostRequestBody_notification_setting(v string) (any, error) {
    result := NOTIFICATIONS_ENABLED_TEAMSPOSTREQUESTBODY_NOTIFICATION_SETTING
    switch v {
        case "notifications_enabled":
            result = NOTIFICATIONS_ENABLED_TEAMSPOSTREQUESTBODY_NOTIFICATION_SETTING
        case "notifications_disabled":
            result = NOTIFICATIONS_DISABLED_TEAMSPOSTREQUESTBODY_NOTIFICATION_SETTING
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamsPostRequestBody_notification_setting(values []TeamsPostRequestBody_notification_setting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsPostRequestBody_notification_setting) isMultiValue() bool {
    return false
}
