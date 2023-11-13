package octokit

import (
    "context"
    i0edeb99caa3e40c70798f25e5a27c55869d044689e27cef74192b0225526d437 "github.com/octokit/go-sdk/codes_of_conduct"
    i1eeadbb6fa76303d0ccc908c11379a0f6ddf7956ef2117c416b0ec336a85cfb6 "github.com/octokit/go-sdk/markdown"
    i2b57abce03ae5d519272fa5a4d14d27eabe20a84c1e6505a42f4a0cb290762fe "github.com/octokit/go-sdk/emojis"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i33fcba8ce855e605b8f15ccd7290a946697d8704c9b472df8732cb5f812ca539 "github.com/octokit/go-sdk/classrooms"
    i3632a260d0435a491465a526c1d3232e223f8dfb30f301a0e2a90899986ab22e "github.com/octokit/go-sdk/assignments"
    i37733996f76e69bcabb3c10d427fc3029787364b1cd897751001a0cceecc2718 "github.com/octokit/go-sdk/rate_limit"
    i39c39e899798c58690f07f81cd25bbe7bb6e4a66960a54447088d0e09b5d85b3 "github.com/octokit/go-sdk/teams"
    i3a55b61fd85afb0f5ae02624d4682159527eaaf45baae493908df8c36374fa09 "github.com/octokit/go-sdk/users"
    i3aeb6dfc461351e09262b214e5fa042be2c3992918cf3c5b7fb8bfc84b3da329 "github.com/octokit/go-sdk/meta"
    i40534efa48a42d9d87dd32b84b61d9f2cd02e7c70efc718d07b5dd8c922a3082 "github.com/octokit/go-sdk/orgs"
    i42d654426097e38e6a732208880929626d3e87e130c09e24b6cf97fa75bf5968 "github.com/octokit/go-sdk/gists"
    i483b6fca40418b3c6c53d08c95df0ec498c96bb620d46afd94c34f4f102ef246 "github.com/octokit/go-sdk/projects"
    i4a767acf86cbeeb88b1c1f4fe5773ca8814092f85d606d1a692b9817a7cf904f "github.com/octokit/go-sdk/repos"
    i57854cfd8c9781a5a5b4edb6c1c8d663efd129565ac2ca09b2a94234feb8997a "github.com/octokit/go-sdk/issues"
    i5fc59549ecc1079ab4a3ae2a6fa6fbd7207b10a363ce9406ea04ea1fa352c872 "github.com/octokit/go-sdk/octocat"
    i6367f7419773125a5a89f6350f326661ee2ccf2b841ed7e375d7739bc1f07b14 "github.com/octokit/go-sdk/zen"
    i70ecdf03a4e989003679733378857ea989eb6dc3dd00337d4508842765226332 "github.com/octokit/go-sdk/feeds"
    i7e41b0ab8655fe18a1668f44b7c46301cc2db130163812e55f15d57b873e581e "github.com/octokit/go-sdk/marketplace_listing"
    i8e6d3b2fe0cb883b7f7ef20ebae358bb4bafed3ea491828f989854afc4dcbc5a "github.com/octokit/go-sdk/gitignore"
    i98fca5cb06513f67bfb7932ca04dd724e9ac571fd2541e3575516ee0a6bac94a "github.com/octokit/go-sdk/organizations"
    iad89f483b87d3a1415ae5444ef2e11914c19b687852bbe8a88cd79ba93ab0137 "github.com/octokit/go-sdk/licenses"
    ib0c6601b728651fca33c81c08294313e043ff51a4d0a8a7f2d6eace09476f3b2 "github.com/octokit/go-sdk/versions"
    ib53ac360127b9b47060bcbb2833f21ac287e193384c04791f99c8790c7fea682 "github.com/octokit/go-sdk/enterprises"
    ibb96144d913375b816f10df2565670010dd845658a3185c2c7f30dfabb39f79a "github.com/octokit/go-sdk/events"
    ibcb1fecb770521b4ae2da04e8c1aa0f7eb53d708e4129db421405df84e448ecf "github.com/octokit/go-sdk/user"
    ibd9714d3f38fa52071d409aee7fcde2aa9b29b35bdc267d4a98143a0a8bea274 "github.com/octokit/go-sdk/notifications"
    ic81fee852b4fbc3390bd8f2b8194938f9d003429b08b22bd71a7115f915ceeb6 "github.com/octokit/go-sdk/appmanifests"
    ice7bc448ddf8e1370eb88c850f4d41a14f3e4175da7f2592ffa766122dce805b "github.com/octokit/go-sdk/app"
    idb2207e7edba6a08542a95006445f6dfbcbb98325160c98ddcc7c91932f985b7 "github.com/octokit/go-sdk/networks"
    idbb345e34ad597d225029889ab3f6e9b555499cfe4ef69836105d8ad016089a4 "github.com/octokit/go-sdk/repositories"
    ie1e6ad15aad51dec6733db6b1f69a17a9f265495d48d64b310a597cefbbeb864 "github.com/octokit/go-sdk/installation"
    ieb3b09d8cff55eae0f26db70ed2d78b98a661d5999ceab8303ac472038668df1 "github.com/octokit/go-sdk/apps"
    ifa5916dfcaba67bcc0fa4a0f970d82e3c003d32f3ea02574f77c0e0b67753c33 "github.com/octokit/go-sdk/applications"
    ifac17d65247495c2dda4e88cef58d53bc84259514ecd0db95b6e827d4895ed8f "github.com/octokit/go-sdk/advisories"
    ifc5dd8531126898fddde54fd0202a34b2754fe3118db9f2d59e1dc1e3ce22e48 "github.com/octokit/go-sdk/search"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApiClientApiClientApiClientGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApiClientApiClientApiClientGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Advisories the advisories property
func (m *ApiClient) Advisories()(*ifac17d65247495c2dda4e88cef58d53bc84259514ecd0db95b6e827d4895ed8f.AdvisoriesRequestBuilder) {
    return ifac17d65247495c2dda4e88cef58d53bc84259514ecd0db95b6e827d4895ed8f.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
func (m *ApiClient) App()(*ice7bc448ddf8e1370eb88c850f4d41a14f3e4175da7f2592ffa766122dce805b.AppRequestBuilder) {
    return ice7bc448ddf8e1370eb88c850f4d41a14f3e4175da7f2592ffa766122dce805b.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
func (m *ApiClient) Applications()(*ifa5916dfcaba67bcc0fa4a0f970d82e3c003d32f3ea02574f77c0e0b67753c33.ApplicationsRequestBuilder) {
    return ifa5916dfcaba67bcc0fa4a0f970d82e3c003d32f3ea02574f77c0e0b67753c33.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
func (m *ApiClient) AppManifests()(*ic81fee852b4fbc3390bd8f2b8194938f9d003429b08b22bd71a7115f915ceeb6.AppManifestsRequestBuilder) {
    return ic81fee852b4fbc3390bd8f2b8194938f9d003429b08b22bd71a7115f915ceeb6.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *ApiClient) Apps()(*ieb3b09d8cff55eae0f26db70ed2d78b98a661d5999ceab8303ac472038668df1.AppsRequestBuilder) {
    return ieb3b09d8cff55eae0f26db70ed2d78b98a661d5999ceab8303ac472038668df1.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments the assignments property
func (m *ApiClient) Assignments()(*i3632a260d0435a491465a526c1d3232e223f8dfb30f301a0e2a90899986ab22e.AssignmentsRequestBuilder) {
    return i3632a260d0435a491465a526c1d3232e223f8dfb30f301a0e2a90899986ab22e.NewAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classrooms the classrooms property
func (m *ApiClient) Classrooms()(*i33fcba8ce855e605b8f15ccd7290a946697d8704c9b472df8732cb5f812ca539.ClassroomsRequestBuilder) {
    return i33fcba8ce855e605b8f15ccd7290a946697d8704c9b472df8732cb5f812ca539.NewClassroomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
func (m *ApiClient) Codes_of_conduct()(*i0edeb99caa3e40c70798f25e5a27c55869d044689e27cef74192b0225526d437.Codes_of_conductRequestBuilder) {
    return i0edeb99caa3e40c70798f25e5a27c55869d044689e27cef74192b0225526d437.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://api.github.com")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Emojis the emojis property
func (m *ApiClient) Emojis()(*i2b57abce03ae5d519272fa5a4d14d27eabe20a84c1e6505a42f4a0cb290762fe.EmojisRequestBuilder) {
    return i2b57abce03ae5d519272fa5a4d14d27eabe20a84c1e6505a42f4a0cb290762fe.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
func (m *ApiClient) Enterprises()(*ib53ac360127b9b47060bcbb2833f21ac287e193384c04791f99c8790c7fea682.EnterprisesRequestBuilder) {
    return ib53ac360127b9b47060bcbb2833f21ac287e193384c04791f99c8790c7fea682.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *ApiClient) Events()(*ibb96144d913375b816f10df2565670010dd845658a3185c2c7f30dfabb39f79a.EventsRequestBuilder) {
    return ibb96144d913375b816f10df2565670010dd845658a3185c2c7f30dfabb39f79a.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
func (m *ApiClient) Feeds()(*i70ecdf03a4e989003679733378857ea989eb6dc3dd00337d4508842765226332.FeedsRequestBuilder) {
    return i70ecdf03a4e989003679733378857ea989eb6dc3dd00337d4508842765226332.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Rootable), nil
}
// Gists the gists property
func (m *ApiClient) Gists()(*i42d654426097e38e6a732208880929626d3e87e130c09e24b6cf97fa75bf5968.GistsRequestBuilder) {
    return i42d654426097e38e6a732208880929626d3e87e130c09e24b6cf97fa75bf5968.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
func (m *ApiClient) Gitignore()(*i8e6d3b2fe0cb883b7f7ef20ebae358bb4bafed3ea491828f989854afc4dcbc5a.GitignoreRequestBuilder) {
    return i8e6d3b2fe0cb883b7f7ef20ebae358bb4bafed3ea491828f989854afc4dcbc5a.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *ApiClient) Installation()(*ie1e6ad15aad51dec6733db6b1f69a17a9f265495d48d64b310a597cefbbeb864.InstallationRequestBuilder) {
    return ie1e6ad15aad51dec6733db6b1f69a17a9f265495d48d64b310a597cefbbeb864.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *ApiClient) Issues()(*i57854cfd8c9781a5a5b4edb6c1c8d663efd129565ac2ca09b2a94234feb8997a.IssuesRequestBuilder) {
    return i57854cfd8c9781a5a5b4edb6c1c8d663efd129565ac2ca09b2a94234feb8997a.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
func (m *ApiClient) Licenses()(*iad89f483b87d3a1415ae5444ef2e11914c19b687852bbe8a88cd79ba93ab0137.LicensesRequestBuilder) {
    return iad89f483b87d3a1415ae5444ef2e11914c19b687852bbe8a88cd79ba93ab0137.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
func (m *ApiClient) Markdown()(*i1eeadbb6fa76303d0ccc908c11379a0f6ddf7956ef2117c416b0ec336a85cfb6.MarkdownRequestBuilder) {
    return i1eeadbb6fa76303d0ccc908c11379a0f6ddf7956ef2117c416b0ec336a85cfb6.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_listing the marketplace_listing property
func (m *ApiClient) Marketplace_listing()(*i7e41b0ab8655fe18a1668f44b7c46301cc2db130163812e55f15d57b873e581e.Marketplace_listingRequestBuilder) {
    return i7e41b0ab8655fe18a1668f44b7c46301cc2db130163812e55f15d57b873e581e.NewMarketplace_listingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
func (m *ApiClient) Meta()(*i3aeb6dfc461351e09262b214e5fa042be2c3992918cf3c5b7fb8bfc84b3da329.MetaRequestBuilder) {
    return i3aeb6dfc461351e09262b214e5fa042be2c3992918cf3c5b7fb8bfc84b3da329.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
func (m *ApiClient) Networks()(*idb2207e7edba6a08542a95006445f6dfbcbb98325160c98ddcc7c91932f985b7.NetworksRequestBuilder) {
    return idb2207e7edba6a08542a95006445f6dfbcbb98325160c98ddcc7c91932f985b7.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
func (m *ApiClient) Notifications()(*ibd9714d3f38fa52071d409aee7fcde2aa9b29b35bdc267d4a98143a0a8bea274.NotificationsRequestBuilder) {
    return ibd9714d3f38fa52071d409aee7fcde2aa9b29b35bdc267d4a98143a0a8bea274.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
func (m *ApiClient) Octocat()(*i5fc59549ecc1079ab4a3ae2a6fa6fbd7207b10a363ce9406ea04ea1fa352c872.OctocatRequestBuilder) {
    return i5fc59549ecc1079ab4a3ae2a6fa6fbd7207b10a363ce9406ea04ea1fa352c872.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
func (m *ApiClient) Organizations()(*i98fca5cb06513f67bfb7932ca04dd724e9ac571fd2541e3575516ee0a6bac94a.OrganizationsRequestBuilder) {
    return i98fca5cb06513f67bfb7932ca04dd724e9ac571fd2541e3575516ee0a6bac94a.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
func (m *ApiClient) Orgs()(*i40534efa48a42d9d87dd32b84b61d9f2cd02e7c70efc718d07b5dd8c922a3082.OrgsRequestBuilder) {
    return i40534efa48a42d9d87dd32b84b61d9f2cd02e7c70efc718d07b5dd8c922a3082.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *ApiClient) Projects()(*i483b6fca40418b3c6c53d08c95df0ec498c96bb620d46afd94c34f4f102ef246.ProjectsRequestBuilder) {
    return i483b6fca40418b3c6c53d08c95df0ec498c96bb620d46afd94c34f4f102ef246.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
func (m *ApiClient) Rate_limit()(*i37733996f76e69bcabb3c10d427fc3029787364b1cd897751001a0cceecc2718.Rate_limitRequestBuilder) {
    return i37733996f76e69bcabb3c10d427fc3029787364b1cd897751001a0cceecc2718.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *ApiClient) Repos()(*i4a767acf86cbeeb88b1c1f4fe5773ca8814092f85d606d1a692b9817a7cf904f.ReposRequestBuilder) {
    return i4a767acf86cbeeb88b1c1f4fe5773ca8814092f85d606d1a692b9817a7cf904f.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
func (m *ApiClient) Repositories()(*idbb345e34ad597d225029889ab3f6e9b555499cfe4ef69836105d8ad016089a4.RepositoriesRequestBuilder) {
    return idbb345e34ad597d225029889ab3f6e9b555499cfe4ef69836105d8ad016089a4.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
func (m *ApiClient) Search()(*ifc5dd8531126898fddde54fd0202a34b2754fe3118db9f2d59e1dc1e3ce22e48.SearchRequestBuilder) {
    return ifc5dd8531126898fddde54fd0202a34b2754fe3118db9f2d59e1dc1e3ce22e48.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *ApiClient) Teams()(*i39c39e899798c58690f07f81cd25bbe7bb6e4a66960a54447088d0e09b5d85b3.TeamsRequestBuilder) {
    return i39c39e899798c58690f07f81cd25bbe7bb6e4a66960a54447088d0e09b5d85b3.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// User the user property
func (m *ApiClient) User()(*ibcb1fecb770521b4ae2da04e8c1aa0f7eb53d708e4129db421405df84e448ecf.UserRequestBuilder) {
    return ibcb1fecb770521b4ae2da04e8c1aa0f7eb53d708e4129db421405df84e448ecf.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
func (m *ApiClient) Users()(*i3a55b61fd85afb0f5ae02624d4682159527eaaf45baae493908df8c36374fa09.UsersRequestBuilder) {
    return i3a55b61fd85afb0f5ae02624d4682159527eaaf45baae493908df8c36374fa09.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
func (m *ApiClient) Versions()(*ib0c6601b728651fca33c81c08294313e043ff51a4d0a8a7f2d6eace09476f3b2.VersionsRequestBuilder) {
    return ib0c6601b728651fca33c81c08294313e043ff51a4d0a8a7f2d6eace09476f3b2.NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
func (m *ApiClient) Zen()(*i6367f7419773125a5a89f6350f326661ee2ccf2b841ed7e375d7739bc1f07b14.ZenRequestBuilder) {
    return i6367f7419773125a5a89f6350f326661ee2ccf2b841ed7e375d7739bc1f07b14.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
