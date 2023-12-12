package item
import (
    "errors"
)
// The notification setting the team has chosen. Editing teams without specifying this parameter leaves `notification_setting` intact. The options are:  * `notifications_enabled` - team members receive notifications when the team is @mentioned.   * `notifications_disabled` - no one receives notifications.
type WithTeam_PatchRequestBody_notification_setting int

const (
    NOTIFICATIONS_ENABLED_WITHTEAM_PATCHREQUESTBODY_NOTIFICATION_SETTING WithTeam_PatchRequestBody_notification_setting = iota
    NOTIFICATIONS_DISABLED_WITHTEAM_PATCHREQUESTBODY_NOTIFICATION_SETTING
)

func (i WithTeam_PatchRequestBody_notification_setting) String() string {
    return []string{"notifications_enabled", "notifications_disabled"}[i]
}
func ParseWithTeam_PatchRequestBody_notification_setting(v string) (any, error) {
    result := NOTIFICATIONS_ENABLED_WITHTEAM_PATCHREQUESTBODY_NOTIFICATION_SETTING
    switch v {
        case "notifications_enabled":
            result = NOTIFICATIONS_ENABLED_WITHTEAM_PATCHREQUESTBODY_NOTIFICATION_SETTING
        case "notifications_disabled":
            result = NOTIFICATIONS_DISABLED_WITHTEAM_PATCHREQUESTBODY_NOTIFICATION_SETTING
        default:
            return 0, errors.New("Unknown WithTeam_PatchRequestBody_notification_setting value: " + v)
    }
    return &result, nil
}
func SerializeWithTeam_PatchRequestBody_notification_setting(values []WithTeam_PatchRequestBody_notification_setting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithTeam_PatchRequestBody_notification_setting) isMultiValue() bool {
    return false
}
