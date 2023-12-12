package orgs

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i3106aaea777cbd2e189aa9c11fbc93d90c3c7272c21ca15937589998ee7e5a2e "github.com/octokit/go-sdk/pkg/github/orgs/item/personalaccesstokens"
)

// ItemPersonalAccessTokensRequestBuilder builds and executes requests for operations under \orgs\{org}\personal-access-tokens
type ItemPersonalAccessTokensRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPersonalAccessTokensRequestBuilderGetQueryParameters lists approved fine-grained personal access tokens owned by organization members that can access organization resources. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
type ItemPersonalAccessTokensRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *i3106aaea777cbd2e189aa9c11fbc93d90c3c7272c21ca15937589998ee7e5a2e.GetDirectionQueryParameterType `uriparametername:"direction"`
    // Only show fine-grained personal access tokens used after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Last_used_after *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"last_used_after"`
    // Only show fine-grained personal access tokens used before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Last_used_before *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"last_used_before"`
    // A list of owner usernames to use to filter the results.
    Owner []string `uriparametername:"owner"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The permission to use to filter the results.
    Permission *string `uriparametername:"permission"`
    // The name of the repository to use to filter the results.
    Repository *string `uriparametername:"repository"`
    // The property by which to sort the results.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property by which to sort the results.
    SortAsGetSortQueryParameterType *i3106aaea777cbd2e189aa9c11fbc93d90c3c7272c21ca15937589998ee7e5a2e.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemPersonalAccessTokensRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPersonalAccessTokensRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPersonalAccessTokensRequestBuilderGetQueryParameters
}
// ItemPersonalAccessTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPersonalAccessTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPat_id gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.personalAccessTokens.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemPersonalAccessTokensRequestBuilder) ByPat_id(pat_id string)(*ItemPersonalAccessTokensWithPat_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if pat_id != "" {
        urlTplParams["pat_id"] = pat_id
    }
    return NewItemPersonalAccessTokensWithPat_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPat_idInteger gets an item from the github.com/octokit/go-sdk/pkg/github/.orgs.item.personalAccessTokens.item collection
func (m *ItemPersonalAccessTokensRequestBuilder) ByPat_idInteger(pat_id int32)(*ItemPersonalAccessTokensWithPat_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["pat_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(pat_id), 10)
    return NewItemPersonalAccessTokensWithPat_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPersonalAccessTokensRequestBuilderInternal instantiates a new PersonalAccessTokensRequestBuilder and sets the default values.
func NewItemPersonalAccessTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokensRequestBuilder) {
    m := &ItemPersonalAccessTokensRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/personal-access-tokens{?per_page*,page*,sort*,direction*,owner*,repository*,permission*,last_used_before*,last_used_after*}", pathParameters),
    }
    return m
}
// NewItemPersonalAccessTokensRequestBuilder instantiates a new PersonalAccessTokensRequestBuilder and sets the default values.
func NewItemPersonalAccessTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPersonalAccessTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists approved fine-grained personal access tokens owned by organization members that can access organization resources. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#list-fine-grained-personal-access-tokens-with-access-to-organization-resources
func (m *ItemPersonalAccessTokensRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPersonalAccessTokensRequestBuilderGetRequestConfiguration)([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationProgrammaticAccessGrantable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrganizationProgrammaticAccessGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationProgrammaticAccessGrantable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationProgrammaticAccessGrantable)
        }
    }
    return val, nil
}
// Post updates the access organization members have to organization resources via fine-grained personal access tokens. Limited to revoking a token's existing access. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// Deprecated: This method is obsolete. Use PostAsPersonalAccessTokensPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#update-the-access-to-organization-resources-via-fine-grained-personal-access-tokens
func (m *ItemPersonalAccessTokensRequestBuilder) Post(ctx context.Context, body ItemPersonalAccessTokensPostRequestBodyable, requestConfiguration *ItemPersonalAccessTokensRequestBuilderPostRequestConfiguration)(ItemPersonalAccessTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPersonalAccessTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPersonalAccessTokensResponseable), nil
}
// PostAsPersonalAccessTokensPostResponse updates the access organization members have to organization resources via fine-grained personal access tokens. Limited to revoking a token's existing access. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#update-the-access-to-organization-resources-via-fine-grained-personal-access-tokens
func (m *ItemPersonalAccessTokensRequestBuilder) PostAsPersonalAccessTokensPostResponse(ctx context.Context, body ItemPersonalAccessTokensPostRequestBodyable, requestConfiguration *ItemPersonalAccessTokensRequestBuilderPostRequestConfiguration)(ItemPersonalAccessTokensPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPersonalAccessTokensPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPersonalAccessTokensPostResponseable), nil
}
// ToGetRequestInformation lists approved fine-grained personal access tokens owned by organization members that can access organization resources. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
func (m *ItemPersonalAccessTokensRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPersonalAccessTokensRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation updates the access organization members have to organization resources via fine-grained personal access tokens. Limited to revoking a token's existing access. Only GitHub Apps can call this API,using the `organization_personal_access_tokens: write` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
func (m *ItemPersonalAccessTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPersonalAccessTokensPostRequestBodyable, requestConfiguration *ItemPersonalAccessTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPersonalAccessTokensRequestBuilder) WithUrl(rawUrl string)(*ItemPersonalAccessTokensRequestBuilder) {
    return NewItemPersonalAccessTokensRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
