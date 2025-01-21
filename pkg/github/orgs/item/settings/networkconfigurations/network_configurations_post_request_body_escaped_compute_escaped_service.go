package networkconfigurations
// The hosted compute service to use for the network configuration.
type NetworkConfigurationsPostRequestBody_compute_service int

const (
    NONE_NETWORKCONFIGURATIONSPOSTREQUESTBODY_COMPUTE_SERVICE NetworkConfigurationsPostRequestBody_compute_service = iota
    ACTIONS_NETWORKCONFIGURATIONSPOSTREQUESTBODY_COMPUTE_SERVICE
)

func (i NetworkConfigurationsPostRequestBody_compute_service) String() string {
    return []string{"none", "actions"}[i]
}
func ParseNetworkConfigurationsPostRequestBody_compute_service(v string) (any, error) {
    result := NONE_NETWORKCONFIGURATIONSPOSTREQUESTBODY_COMPUTE_SERVICE
    switch v {
        case "none":
            result = NONE_NETWORKCONFIGURATIONSPOSTREQUESTBODY_COMPUTE_SERVICE
        case "actions":
            result = ACTIONS_NETWORKCONFIGURATIONSPOSTREQUESTBODY_COMPUTE_SERVICE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeNetworkConfigurationsPostRequestBody_compute_service(values []NetworkConfigurationsPostRequestBody_compute_service) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NetworkConfigurationsPostRequestBody_compute_service) isMultiValue() bool {
    return false
}
