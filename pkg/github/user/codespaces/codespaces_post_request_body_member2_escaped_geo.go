package codespaces
// The geographic area for this codespace. If not specified, the value is assigned by IP. This property replaces `location`, which is closing down.
type CodespacesPostRequestBodyMember2_geo int

const (
    EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO CodespacesPostRequestBodyMember2_geo = iota
    SOUTHEASTASIA_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
    USEAST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
    USWEST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
)

func (i CodespacesPostRequestBodyMember2_geo) String() string {
    return []string{"EuropeWest", "SoutheastAsia", "UsEast", "UsWest"}[i]
}
func ParseCodespacesPostRequestBodyMember2_geo(v string) (any, error) {
    result := EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
    switch v {
        case "EuropeWest":
            result = EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
        case "SoutheastAsia":
            result = SOUTHEASTASIA_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
        case "UsEast":
            result = USEAST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
        case "UsWest":
            result = USWEST_CODESPACESPOSTREQUESTBODYMEMBER2_GEO
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodespacesPostRequestBodyMember2_geo(values []CodespacesPostRequestBodyMember2_geo) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespacesPostRequestBodyMember2_geo) isMultiValue() bool {
    return false
}
