package models
// The hosted compute service the network configuration supports.
type NetworkConfiguration_compute_service int

const (
    NONE_NETWORKCONFIGURATION_COMPUTE_SERVICE NetworkConfiguration_compute_service = iota
    ACTIONS_NETWORKCONFIGURATION_COMPUTE_SERVICE
    CODESPACES_NETWORKCONFIGURATION_COMPUTE_SERVICE
)

func (i NetworkConfiguration_compute_service) String() string {
    return []string{"none", "actions", "codespaces"}[i]
}
func ParseNetworkConfiguration_compute_service(v string) (any, error) {
    result := NONE_NETWORKCONFIGURATION_COMPUTE_SERVICE
    switch v {
        case "none":
            result = NONE_NETWORKCONFIGURATION_COMPUTE_SERVICE
        case "actions":
            result = ACTIONS_NETWORKCONFIGURATION_COMPUTE_SERVICE
        case "codespaces":
            result = CODESPACES_NETWORKCONFIGURATION_COMPUTE_SERVICE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkConfiguration_compute_service(values []NetworkConfiguration_compute_service) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkConfiguration_compute_service) isMultiValue() bool {
    return false
}
