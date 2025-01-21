package item
// The hosted compute service to use for the network configuration.
type WithNetwork_configuration_PatchRequestBody_compute_service int

const (
    NONE_WITHNETWORK_CONFIGURATION_PATCHREQUESTBODY_COMPUTE_SERVICE WithNetwork_configuration_PatchRequestBody_compute_service = iota
    ACTIONS_WITHNETWORK_CONFIGURATION_PATCHREQUESTBODY_COMPUTE_SERVICE
)

func (i WithNetwork_configuration_PatchRequestBody_compute_service) String() string {
    return []string{"none", "actions"}[i]
}
func ParseWithNetwork_configuration_PatchRequestBody_compute_service(v string) (any, error) {
    result := NONE_WITHNETWORK_CONFIGURATION_PATCHREQUESTBODY_COMPUTE_SERVICE
    switch v {
        case "none":
            result = NONE_WITHNETWORK_CONFIGURATION_PATCHREQUESTBODY_COMPUTE_SERVICE
        case "actions":
            result = ACTIONS_WITHNETWORK_CONFIGURATION_PATCHREQUESTBODY_COMPUTE_SERVICE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithNetwork_configuration_PatchRequestBody_compute_service(values []WithNetwork_configuration_PatchRequestBody_compute_service) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithNetwork_configuration_PatchRequestBody_compute_service) isMultiValue() bool {
    return false
}
