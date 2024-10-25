package codespaces
// The geographic area for this codespace. If not specified, the value is assigned by IP. This property replaces `location`, which is closing down.
type CodespacesPostRequestBodyMember1_geo int

const (
    EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO CodespacesPostRequestBodyMember1_geo = iota
    SOUTHEASTASIA_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
    USEAST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
    USWEST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
)

func (i CodespacesPostRequestBodyMember1_geo) String() string {
    return []string{"EuropeWest", "SoutheastAsia", "UsEast", "UsWest"}[i]
}
func ParseCodespacesPostRequestBodyMember1_geo(v string) (any, error) {
    result := EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
    switch v {
        case "EuropeWest":
            result = EUROPEWEST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
        case "SoutheastAsia":
            result = SOUTHEASTASIA_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
        case "UsEast":
            result = USEAST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
        case "UsWest":
            result = USWEST_CODESPACESPOSTREQUESTBODYMEMBER1_GEO
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCodespacesPostRequestBodyMember1_geo(values []CodespacesPostRequestBodyMember1_geo) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodespacesPostRequestBodyMember1_geo) isMultiValue() bool {
    return false
}
