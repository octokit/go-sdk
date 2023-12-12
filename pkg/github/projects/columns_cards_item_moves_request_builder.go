package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ColumnsCardsItemMovesRequestBuilder builds and executes requests for operations under \projects\columns\cards\{card_id}\moves
type ColumnsCardsItemMovesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewColumnsCardsItemMovesRequestBuilderInternal instantiates a new MovesRequestBuilder and sets the default values.
func NewColumnsCardsItemMovesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsCardsItemMovesRequestBuilder) {
    m := &ColumnsCardsItemMovesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/columns/cards/{card_id}/moves", pathParameters),
    }
    return m
}
// NewColumnsCardsItemMovesRequestBuilder instantiates a new MovesRequestBuilder and sets the default values.
func NewColumnsCardsItemMovesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsCardsItemMovesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsCardsItemMovesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post move a project card
// Deprecated: This method is obsolete. Use PostAsMovesPostResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#move-a-project-card
func (m *ColumnsCardsItemMovesRequestBuilder) Post(ctx context.Context, body ColumnsCardsItemMovesPostRequestBodyable, requestConfiguration *ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration)(ColumnsCardsItemMovesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": CreateColumnsCardsItemMoves403ErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "503": CreateColumnsCardsItemMoves503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateColumnsCardsItemMovesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ColumnsCardsItemMovesResponseable), nil
}
// PostAsMovesPostResponse move a project card
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#move-a-project-card
func (m *ColumnsCardsItemMovesRequestBuilder) PostAsMovesPostResponse(ctx context.Context, body ColumnsCardsItemMovesPostRequestBodyable, requestConfiguration *ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration)(ColumnsCardsItemMovesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": CreateColumnsCardsItemMoves403ErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "503": CreateColumnsCardsItemMoves503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateColumnsCardsItemMovesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ColumnsCardsItemMovesPostResponseable), nil
}
func (m *ColumnsCardsItemMovesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ColumnsCardsItemMovesPostRequestBodyable, requestConfiguration *ColumnsCardsItemMovesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ColumnsCardsItemMovesRequestBuilder) WithUrl(rawUrl string)(*ColumnsCardsItemMovesRequestBuilder) {
    return NewColumnsCardsItemMovesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
