package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOutside_collaboratorsWithUsernameItemRequestBuilder builds and executes requests for operations under \orgs\{org}\outside_collaborators\{username}
type ItemOutside_collaboratorsWithUsernameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutside_collaboratorsWithUsernameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutside_collaboratorsWithUsernameItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOutside_collaboratorsWithUsernameItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutside_collaboratorsWithUsernameItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    m := &ItemOutside_collaboratorsWithUsernameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/outside_collaborators/{username}", pathParameters),
    }
    return m
}
// NewItemOutside_collaboratorsWithUsernameItemRequestBuilder instantiates a new WithUsernameItemRequestBuilder and sets the default values.
func NewItemOutside_collaboratorsWithUsernameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutside_collaboratorsWithUsernameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removing a user from this list will remove them from all the organization's repositories.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/outside-collaborators#remove-outside-collaborator-from-an-organization
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOutside_collaboratorsWithUsernameItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": CreateItemOutside_collaboratorsItemWithUsername422ErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Put when an organization member is converted to an outside collaborator, they'll only have access to the repositories that their current team membership allows. The user will no longer be a member of the organization. For more information, see "[Converting an organization member to an outside collaborator](https://docs.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)". Converting an organization member to an outside collaborator may be restricted by enterprise administrators. For more information, see "[Enforcing repository management policies in your enterprise](https://docs.github.com/admin/policies/enforcing-policies-for-your-enterprise/enforcing-repository-management-policies-in-your-enterprise#enforcing-a-policy-for-inviting-outside-collaborators-to-repositories)."
// Deprecated: This method is obsolete. Use PutAsWithUsernamePutResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/outside-collaborators#convert-an-organization-member-to-outside-collaborator
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) Put(ctx context.Context, body ItemOutside_collaboratorsItemWithUsernamePutRequestBodyable, requestConfiguration *ItemOutside_collaboratorsWithUsernameItemRequestBuilderPutRequestConfiguration)(ItemOutside_collaboratorsItemWithUsernameResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutside_collaboratorsItemWithUsernameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutside_collaboratorsItemWithUsernameResponseable), nil
}
// PutAsWithUsernamePutResponse when an organization member is converted to an outside collaborator, they'll only have access to the repositories that their current team membership allows. The user will no longer be a member of the organization. For more information, see "[Converting an organization member to an outside collaborator](https://docs.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)". Converting an organization member to an outside collaborator may be restricted by enterprise administrators. For more information, see "[Enforcing repository management policies in your enterprise](https://docs.github.com/admin/policies/enforcing-policies-for-your-enterprise/enforcing-repository-management-policies-in-your-enterprise#enforcing-a-policy-for-inviting-outside-collaborators-to-repositories)."
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/outside-collaborators#convert-an-organization-member-to-outside-collaborator
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) PutAsWithUsernamePutResponse(ctx context.Context, body ItemOutside_collaboratorsItemWithUsernamePutRequestBodyable, requestConfiguration *ItemOutside_collaboratorsWithUsernameItemRequestBuilderPutRequestConfiguration)(ItemOutside_collaboratorsItemWithUsernamePutResponseable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutside_collaboratorsItemWithUsernamePutResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutside_collaboratorsItemWithUsernamePutResponseable), nil
}
// ToDeleteRequestInformation removing a user from this list will remove them from all the organization's repositories.
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOutside_collaboratorsWithUsernameItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation when an organization member is converted to an outside collaborator, they'll only have access to the repositories that their current team membership allows. The user will no longer be a member of the organization. For more information, see "[Converting an organization member to an outside collaborator](https://docs.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)". Converting an organization member to an outside collaborator may be restricted by enterprise administrators. For more information, see "[Enforcing repository management policies in your enterprise](https://docs.github.com/admin/policies/enforcing-policies-for-your-enterprise/enforcing-repository-management-policies-in-your-enterprise#enforcing-a-policy-for-inviting-outside-collaborators-to-repositories)."
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemOutside_collaboratorsItemWithUsernamePutRequestBodyable, requestConfiguration *ItemOutside_collaboratorsWithUsernameItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOutside_collaboratorsWithUsernameItemRequestBuilder) WithUrl(rawUrl string)(*ItemOutside_collaboratorsWithUsernameItemRequestBuilder) {
    return NewItemOutside_collaboratorsWithUsernameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
