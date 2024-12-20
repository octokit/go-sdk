package attach
// The type of repositories to attach the configuration to. `selected` means the configuration will be attached to only the repositories specified by `selected_repository_ids`
type AttachPostRequestBody_scope int

const (
    ALL_ATTACHPOSTREQUESTBODY_SCOPE AttachPostRequestBody_scope = iota
    ALL_WITHOUT_CONFIGURATIONS_ATTACHPOSTREQUESTBODY_SCOPE
)

func (i AttachPostRequestBody_scope) String() string {
    return []string{"all", "all_without_configurations"}[i]
}
func ParseAttachPostRequestBody_scope(v string) (any, error) {
    result := ALL_ATTACHPOSTREQUESTBODY_SCOPE
    switch v {
        case "all":
            result = ALL_ATTACHPOSTREQUESTBODY_SCOPE
        case "all_without_configurations":
            result = ALL_WITHOUT_CONFIGURATIONS_ATTACHPOSTREQUESTBODY_SCOPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAttachPostRequestBody_scope(values []AttachPostRequestBody_scope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AttachPostRequestBody_scope) isMultiValue() bool {
    return false
}
