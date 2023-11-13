package repos

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCheckSuitesPreferencesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\check-suites\preferences
type ItemItemCheckSuitesPreferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCheckSuitesPreferencesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemCheckSuitesPreferencesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemCheckSuitesPreferencesRequestBuilderInternal instantiates a new PreferencesRequestBuilder and sets the default values.
func NewItemItemCheckSuitesPreferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesPreferencesRequestBuilder) {
    m := &ItemItemCheckSuitesPreferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/check-suites/preferences", pathParameters),
    }
    return m
}
// NewItemItemCheckSuitesPreferencesRequestBuilder instantiates a new PreferencesRequestBuilder and sets the default values.
func NewItemItemCheckSuitesPreferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckSuitesPreferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckSuitesPreferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Patch changes the default automatic flow when creating check suites. By default, a check suite is automatically created each time code is pushed to a repository. When you disable the automatic creation of check suites, you can manually [Create a check suite](https://docs.github.com/rest/checks/suites#create-a-check-suite). You must have admin permissions in the repository to set preferences for check suites.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/suites#update-repository-preferences-for-check-suites
func (m *ItemItemCheckSuitesPreferencesRequestBuilder) Patch(ctx context.Context, body ItemItemCheckSuitesPreferencesPatchRequestBodyable, requestConfiguration *ItemItemCheckSuitesPreferencesRequestBuilderPatchRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CheckSuitePreferenceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCheckSuitePreferenceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CheckSuitePreferenceable), nil
}
// ToPatchRequestInformation changes the default automatic flow when creating check suites. By default, a check suite is automatically created each time code is pushed to a repository. When you disable the automatic creation of check suites, you can manually [Create a check suite](https://docs.github.com/rest/checks/suites#create-a-check-suite). You must have admin permissions in the repository to set preferences for check suites.
func (m *ItemItemCheckSuitesPreferencesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemItemCheckSuitesPreferencesPatchRequestBodyable, requestConfiguration *ItemItemCheckSuitesPreferencesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemCheckSuitesPreferencesRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckSuitesPreferencesRequestBuilder) {
    return NewItemItemCheckSuitesPreferencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
