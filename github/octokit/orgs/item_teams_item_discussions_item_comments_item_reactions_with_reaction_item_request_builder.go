package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\teams\{team_slug}\discussions\{discussion_number}\comments\{comment_number}\reactions\{reaction_id}
type ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    m := &ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions/{reaction_id}", pathParameters),
    }
    return m
}
// NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder instantiates a new WithReaction_ItemRequestBuilder and sets the default values.
func NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete **Note:** You can also specify a team or organization with `team_id` and `org_id` using the route `DELETE /organizations/:org_id/team/:team_id/discussions/:discussion_number/comments/:comment_number/reactions/:reaction_id`.Delete a reaction to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/reactions/reactions#delete-team-discussion-comment-reaction
func (m *ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation **Note:** You can also specify a team or organization with `team_id` and `org_id` using the route `DELETE /organizations/:org_id/team/:team_id/discussions/:discussion_number/comments/:comment_number/reactions/:reaction_id`.Delete a reaction to a [team discussion comment](https://docs.github.com/rest/teams/discussion-comments#get-a-discussion-comment). OAuth access tokens require the `write:discussion` [scope](https://docs.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/).
func (m *ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder) {
    return NewItemTeamsItemDiscussionsItemCommentsItemReactionsWithReaction_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
