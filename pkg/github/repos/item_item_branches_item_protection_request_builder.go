package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemBranchesItemProtectionRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\branches\{branch}\protection
type ItemItemBranchesItemProtectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBranchesItemProtectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemBranchesItemProtectionRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemBranchesItemProtectionRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemBranchesItemProtectionRequestBuilderInternal instantiates a new ProtectionRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/branches/{branch}/protection", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRequestBuilder instantiates a new ProtectionRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#delete-branch-protection
func (m *ItemItemBranchesItemProtectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Enforce_admins the enforce_admins property
func (m *ItemItemBranchesItemProtectionRequestBuilder) Enforce_admins()(*ItemItemBranchesItemProtectionEnforce_adminsRequestBuilder) {
    return NewItemItemBranchesItemProtectionEnforce_adminsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-branch-protection
func (m *ItemItemBranchesItemProtectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderGetRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.BranchProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBranchProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.BranchProtectionable), nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Protecting a branch requires admin or owner permissions to the repository.**Note**: Passing new arrays of `users` and `teams` replaces their previous values.**Note**: The list of users, apps, and teams in total is limited to 100 items.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#update-branch-protection
func (m *ItemItemBranchesItemProtectionRequestBuilder) Put(ctx context.Context, body ItemItemBranchesItemProtectionPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderPutRequestConfiguration)(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProtectedBranchable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateProtectedBranchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProtectedBranchable), nil
}
// Required_pull_request_reviews the required_pull_request_reviews property
func (m *ItemItemBranchesItemProtectionRequestBuilder) Required_pull_request_reviews()(*ItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequired_pull_request_reviewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Required_signatures the required_signatures property
func (m *ItemItemBranchesItemProtectionRequestBuilder) Required_signatures()(*ItemItemBranchesItemProtectionRequired_signaturesRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequired_signaturesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Required_status_checks the required_status_checks property
func (m *ItemItemBranchesItemProtectionRequestBuilder) Required_status_checks()(*ItemItemBranchesItemProtectionRequired_status_checksRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequired_status_checksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restrictions the restrictions property
func (m *ItemItemBranchesItemProtectionRequestBuilder) Restrictions()(*ItemItemBranchesItemProtectionRestrictionsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRestrictionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
func (m *ItemItemBranchesItemProtectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Protecting a branch requires admin or owner permissions to the repository.**Note**: Passing new arrays of `users` and `teams` replaces their previous values.**Note**: The list of users, apps, and teams in total is limited to 100 items.
func (m *ItemItemBranchesItemProtectionRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemBranchesItemProtectionPutRequestBodyable, requestConfiguration *ItemItemBranchesItemProtectionRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *ItemItemBranchesItemProtectionRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
