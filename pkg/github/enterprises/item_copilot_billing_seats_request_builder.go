package enterprises

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemCopilotBillingSeatsRequestBuilder builds and executes requests for operations under \enterprises\{enterprise}\copilot\billing\seats
type ItemCopilotBillingSeatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCopilotBillingSeatsRequestBuilderGetQueryParameters > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all Copilot seats currently being billed for across organizations or enterprise teams for an enterprise with a Copilot Business or Copilot Enterprise subscription.Users with access through multiple organizations or enterprise teams will only be counted toward `total_seats` once.For each organization or enterprise team which grants Copilot access to a user, a seat detail object will appear in the `seats` array.Each seat object contains information about the assigned user's most recent Copilot activity. Users must havetelemetry enabled in their IDE for Copilot in the IDE activity to be reflected in `last_activity_at`. For more information about activity data,see "[Reviewing user activity data for Copilot in your organization](https://docs.github.com/copilot/managing-copilot/managing-github-copilot-in-your-organization/reviewing-activity-related-to-github-copilot-in-your-organization/reviewing-user-activity-data-for-copilot-in-your-organization)."Only enterprise owners and billing managers can view assigned Copilot seats across their child organizations or enterprise teams.Personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
type ItemCopilotBillingSeatsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// NewItemCopilotBillingSeatsRequestBuilderInternal instantiates a new ItemCopilotBillingSeatsRequestBuilder and sets the default values.
func NewItemCopilotBillingSeatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSeatsRequestBuilder) {
    m := &ItemCopilotBillingSeatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/enterprises/{enterprise}/copilot/billing/seats{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemCopilotBillingSeatsRequestBuilder instantiates a new ItemCopilotBillingSeatsRequestBuilder and sets the default values.
func NewItemCopilotBillingSeatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCopilotBillingSeatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCopilotBillingSeatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all Copilot seats currently being billed for across organizations or enterprise teams for an enterprise with a Copilot Business or Copilot Enterprise subscription.Users with access through multiple organizations or enterprise teams will only be counted toward `total_seats` once.For each organization or enterprise team which grants Copilot access to a user, a seat detail object will appear in the `seats` array.Each seat object contains information about the assigned user's most recent Copilot activity. Users must havetelemetry enabled in their IDE for Copilot in the IDE activity to be reflected in `last_activity_at`. For more information about activity data,see "[Reviewing user activity data for Copilot in your organization](https://docs.github.com/copilot/managing-copilot/managing-github-copilot-in-your-organization/reviewing-activity-related-to-github-copilot-in-your-organization/reviewing-user-activity-data-for-copilot-in-your-organization)."Only enterprise owners and billing managers can view assigned Copilot seats across their child organizations or enterprise teams.Personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
// returns a ItemCopilotBillingSeatsGetResponseable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/copilot/copilot-user-management#list-all-copilot-seat-assignments-for-an-enterprise
func (m *ItemCopilotBillingSeatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCopilotBillingSeatsRequestBuilderGetQueryParameters])(ItemCopilotBillingSeatsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "500": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCopilotBillingSeatsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCopilotBillingSeatsGetResponseable), nil
}
// ToGetRequestInformation > [!NOTE]> This endpoint is in public preview and is subject to change.Lists all Copilot seats currently being billed for across organizations or enterprise teams for an enterprise with a Copilot Business or Copilot Enterprise subscription.Users with access through multiple organizations or enterprise teams will only be counted toward `total_seats` once.For each organization or enterprise team which grants Copilot access to a user, a seat detail object will appear in the `seats` array.Each seat object contains information about the assigned user's most recent Copilot activity. Users must havetelemetry enabled in their IDE for Copilot in the IDE activity to be reflected in `last_activity_at`. For more information about activity data,see "[Reviewing user activity data for Copilot in your organization](https://docs.github.com/copilot/managing-copilot/managing-github-copilot-in-your-organization/reviewing-activity-related-to-github-copilot-in-your-organization/reviewing-user-activity-data-for-copilot-in-your-organization)."Only enterprise owners and billing managers can view assigned Copilot seats across their child organizations or enterprise teams.Personal access tokens (classic) need either the `manage_billing:copilot` or `read:enterprise` scopes to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemCopilotBillingSeatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemCopilotBillingSeatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCopilotBillingSeatsRequestBuilder when successful
func (m *ItemCopilotBillingSeatsRequestBuilder) WithUrl(rawUrl string)(*ItemCopilotBillingSeatsRequestBuilder) {
    return NewItemCopilotBillingSeatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
