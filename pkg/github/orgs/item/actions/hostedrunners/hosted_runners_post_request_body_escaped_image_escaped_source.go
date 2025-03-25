package hostedrunners
// The source of the runner image.
type HostedRunnersPostRequestBody_image_source int

const (
    GITHUB_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE HostedRunnersPostRequestBody_image_source = iota
    PARTNER_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
    CUSTOM_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
)

func (i HostedRunnersPostRequestBody_image_source) String() string {
    return []string{"github", "partner", "custom"}[i]
}
func ParseHostedRunnersPostRequestBody_image_source(v string) (any, error) {
    result := GITHUB_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
    switch v {
        case "github":
            result = GITHUB_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
        case "partner":
            result = PARTNER_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
        case "custom":
            result = CUSTOM_HOSTEDRUNNERSPOSTREQUESTBODY_IMAGE_SOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHostedRunnersPostRequestBody_image_source(values []HostedRunnersPostRequestBody_image_source) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HostedRunnersPostRequestBody_image_source) isMultiValue() bool {
    return false
}
