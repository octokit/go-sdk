package models
// The status of the runner.
type ActionsHostedRunner_status int

const (
    READY_ACTIONSHOSTEDRUNNER_STATUS ActionsHostedRunner_status = iota
    PROVISIONING_ACTIONSHOSTEDRUNNER_STATUS
    SHUTDOWN_ACTIONSHOSTEDRUNNER_STATUS
    DELETING_ACTIONSHOSTEDRUNNER_STATUS
    STUCK_ACTIONSHOSTEDRUNNER_STATUS
)

func (i ActionsHostedRunner_status) String() string {
    return []string{"Ready", "Provisioning", "Shutdown", "Deleting", "Stuck"}[i]
}
func ParseActionsHostedRunner_status(v string) (any, error) {
    result := READY_ACTIONSHOSTEDRUNNER_STATUS
    switch v {
        case "Ready":
            result = READY_ACTIONSHOSTEDRUNNER_STATUS
        case "Provisioning":
            result = PROVISIONING_ACTIONSHOSTEDRUNNER_STATUS
        case "Shutdown":
            result = SHUTDOWN_ACTIONSHOSTEDRUNNER_STATUS
        case "Deleting":
            result = DELETING_ACTIONSHOSTEDRUNNER_STATUS
        case "Stuck":
            result = STUCK_ACTIONSHOSTEDRUNNER_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeActionsHostedRunner_status(values []ActionsHostedRunner_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionsHostedRunner_status) isMultiValue() bool {
    return false
}
