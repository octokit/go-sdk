package octokit

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
    i006ac19b46f7860b7359e6cc742a2213da1a512ec1c7d0aae6cf413f77880e18 "github.com/octokit/go-sdk/github/octokit/orgs"
    i036f8c76ca722afbbb72ebcdb51244c927aff7bca1cfebb950994aac3db6a206 "github.com/octokit/go-sdk/github/octokit/user"
    i03a3ace4ab67b9bd12c318876b591a42401445b28fc0d15af81ca9d68f84afef "github.com/octokit/go-sdk/github/octokit/codes_of_conduct"
    i0a41a67c67557b401d607bec1017ba369edee449be6203262889bd4f6fe8018c "github.com/octokit/go-sdk/github/octokit/zen"
    i0c0a073bba79314de6b968ee3ac963288d9a717dea40b7ef084afe188ea55d45 "github.com/octokit/go-sdk/github/octokit/versions"
    i1790f60a6d24b15240ea65ca9183e125346ae1d30cb03d8b67ebb65a5e9af90a "github.com/octokit/go-sdk/github/octokit/issues"
    i198693a66d8516e0b95f63eba53f5bd9faa4b866b296f65f70a160d20cd00869 "github.com/octokit/go-sdk/github/octokit/assignments"
    i226c6e0c3eb27f11a036f911a20cd78714540d256a30ef62bb5c07932ebbf1ad "github.com/octokit/go-sdk/github/octokit/applications"
    i24a6da0dafb64b99af6054f1b79b91fb4c90b5a5223b1117b3329c6a4df15184 "github.com/octokit/go-sdk/github/octokit/enterprises"
    i32b2c35b91803243459a19797935c406e08a7ac5f2e535b29c5578423b31588f "github.com/octokit/go-sdk/github/octokit/rate_limit"
    i42234227323bb9085c39a8953a9cb5e6a70dc990771162276a9ca718fbba7d5e "github.com/octokit/go-sdk/github/octokit/markdown"
    i42ad804983ceb46bd8a65545b25d3ef4fceb8c76cb775ee031e99d72d7501ea1 "github.com/octokit/go-sdk/github/octokit/advisories"
    i4a11dbf1fac8370be6f5106b2269092133cf4ec314afd089fab27dc7dc1f3241 "github.com/octokit/go-sdk/github/octokit/licenses"
    i5406244bdf8038feac942b019f5d863c71d17da72a9cee044e224423a517481f "github.com/octokit/go-sdk/github/octokit/projects"
    i5754c6de91a6e89e2993d008ea667142335a8fa0a5e45a6ee5b29f29c02f34a3 "github.com/octokit/go-sdk/github/octokit/feeds"
    i57dea5e4bfbe19008da64b99b06c98a9e0e8ed54cf56a5d62df6b10a5b352236 "github.com/octokit/go-sdk/github/octokit/app"
    i58d6850d1b0c5fc1569fd86453afdff88f2e933fdcb4cd121b1e20d9d527a0b8 "github.com/octokit/go-sdk/github/octokit/users"
    i5f31d8076704b2a33cf0d9496caffcad199649c2b9693da84305db274106d6c9 "github.com/octokit/go-sdk/github/octokit/networks"
    i60d210be466537b7cfe7e84a79365b690725bffe8bdaff504696996406b87a8b "github.com/octokit/go-sdk/github/octokit/meta"
    i71e3d6b9d82a10e0c351b2793a30816189ce7b7e3809438d9637fddf19f2e4de "github.com/octokit/go-sdk/github/octokit/teams"
    i7238c553ea63f232858d04abc7aa42cf1c6703df7c6b4a89ed97696c881b1b86 "github.com/octokit/go-sdk/github/octokit/octocat"
    i892fd18bf3f479d508c18121aaa79a47af092fb22bf2c016696f4ef49ac0f2df "github.com/octokit/go-sdk/github/octokit/installation"
    i8eb05c4a8c40e9a66f7c3bd436ffaf5869db0e78726339a6c29d711921e3c8ca "github.com/octokit/go-sdk/github/octokit/marketplace_listing"
    i934389078e2d9c63343c6b357201b85c703e6fc5dcfeec1d578a807bea450520 "github.com/octokit/go-sdk/github/octokit/classrooms"
    i95fc277b98fb5717711ae14a55e97eccdd5277a1df93e57f237ea54e3c23e865 "github.com/octokit/go-sdk/github/octokit/appmanifests"
    i97d6f896d44068be5b5e21e5d1fa2c5aaa75945ea97ac569a00dd04274a95d25 "github.com/octokit/go-sdk/github/octokit/gists"
    ia818d425c90b7b6ad2d2240b3a6039727c3a3adb3b856b83297abdb649bf7b5a "github.com/octokit/go-sdk/github/octokit/emojis"
    ibcd8a3e35b6ab970afb6dc75815c8ca8f5a5873760b244cfb82ad1ada7d36271 "github.com/octokit/go-sdk/github/octokit/search"
    ic5c472fad8099124bbac547779804efde5daea6ed48cfd69bad1391db8d32eaf "github.com/octokit/go-sdk/github/octokit/repositories"
    icac773a834232ff9261f69ba5960a9a9c9a9a52a6372d61a99ae834174d581c8 "github.com/octokit/go-sdk/github/octokit/gitignore"
    id498c92982d2f948ace86e8cde2d5fff55ba6cb718f448e213df19b59968ecd3 "github.com/octokit/go-sdk/github/octokit/organizations"
    id93ee658913eeb4bcf71cd66de862383ff3cdd07db79c518b345fc9dd4e57406 "github.com/octokit/go-sdk/github/octokit/notifications"
    iddc31237199800bb0c9999de06596e2979bdd4a3f890a253e6c5893742ea8d64 "github.com/octokit/go-sdk/github/octokit/repos"
    idf7cb172d723d671741ceccc919f4de4b6893a1576f2bde93e7d1a7f1a5260e1 "github.com/octokit/go-sdk/github/octokit/events"
    ie7bb3097c3ea6f5b2ba5bdff96293a9912b3e0f1448edbf76d4581f1985766da "github.com/octokit/go-sdk/github/octokit/apps"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApiClientApiClientGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApiClientApiClientGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Advisories the advisories property
func (m *ApiClient) Advisories()(*i42ad804983ceb46bd8a65545b25d3ef4fceb8c76cb775ee031e99d72d7501ea1.AdvisoriesRequestBuilder) {
    return i42ad804983ceb46bd8a65545b25d3ef4fceb8c76cb775ee031e99d72d7501ea1.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
func (m *ApiClient) App()(*i57dea5e4bfbe19008da64b99b06c98a9e0e8ed54cf56a5d62df6b10a5b352236.AppRequestBuilder) {
    return i57dea5e4bfbe19008da64b99b06c98a9e0e8ed54cf56a5d62df6b10a5b352236.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
func (m *ApiClient) Applications()(*i226c6e0c3eb27f11a036f911a20cd78714540d256a30ef62bb5c07932ebbf1ad.ApplicationsRequestBuilder) {
    return i226c6e0c3eb27f11a036f911a20cd78714540d256a30ef62bb5c07932ebbf1ad.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
func (m *ApiClient) AppManifests()(*i95fc277b98fb5717711ae14a55e97eccdd5277a1df93e57f237ea54e3c23e865.AppManifestsRequestBuilder) {
    return i95fc277b98fb5717711ae14a55e97eccdd5277a1df93e57f237ea54e3c23e865.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *ApiClient) Apps()(*ie7bb3097c3ea6f5b2ba5bdff96293a9912b3e0f1448edbf76d4581f1985766da.AppsRequestBuilder) {
    return ie7bb3097c3ea6f5b2ba5bdff96293a9912b3e0f1448edbf76d4581f1985766da.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments the assignments property
func (m *ApiClient) Assignments()(*i198693a66d8516e0b95f63eba53f5bd9faa4b866b296f65f70a160d20cd00869.AssignmentsRequestBuilder) {
    return i198693a66d8516e0b95f63eba53f5bd9faa4b866b296f65f70a160d20cd00869.NewAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classrooms the classrooms property
func (m *ApiClient) Classrooms()(*i934389078e2d9c63343c6b357201b85c703e6fc5dcfeec1d578a807bea450520.ClassroomsRequestBuilder) {
    return i934389078e2d9c63343c6b357201b85c703e6fc5dcfeec1d578a807bea450520.NewClassroomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
func (m *ApiClient) Codes_of_conduct()(*i03a3ace4ab67b9bd12c318876b591a42401445b28fc0d15af81ca9d68f84afef.Codes_of_conductRequestBuilder) {
    return i03a3ace4ab67b9bd12c318876b591a42401445b28fc0d15af81ca9d68f84afef.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ApiClient) Emojis()(*ia818d425c90b7b6ad2d2240b3a6039727c3a3adb3b856b83297abdb649bf7b5a.EmojisRequestBuilder) {
    return ia818d425c90b7b6ad2d2240b3a6039727c3a3adb3b856b83297abdb649bf7b5a.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
func (m *ApiClient) Enterprises()(*i24a6da0dafb64b99af6054f1b79b91fb4c90b5a5223b1117b3329c6a4df15184.EnterprisesRequestBuilder) {
    return i24a6da0dafb64b99af6054f1b79b91fb4c90b5a5223b1117b3329c6a4df15184.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *ApiClient) Events()(*idf7cb172d723d671741ceccc919f4de4b6893a1576f2bde93e7d1a7f1a5260e1.EventsRequestBuilder) {
    return idf7cb172d723d671741ceccc919f4de4b6893a1576f2bde93e7d1a7f1a5260e1.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
func (m *ApiClient) Feeds()(*i5754c6de91a6e89e2993d008ea667142335a8fa0a5e45a6ee5b29f29c02f34a3.FeedsRequestBuilder) {
    return i5754c6de91a6e89e2993d008ea667142335a8fa0a5e45a6ee5b29f29c02f34a3.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *ApiClientApiClientGetRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.Rootable), nil
}
// Gists the gists property
func (m *ApiClient) Gists()(*i97d6f896d44068be5b5e21e5d1fa2c5aaa75945ea97ac569a00dd04274a95d25.GistsRequestBuilder) {
    return i97d6f896d44068be5b5e21e5d1fa2c5aaa75945ea97ac569a00dd04274a95d25.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
func (m *ApiClient) Gitignore()(*icac773a834232ff9261f69ba5960a9a9c9a9a52a6372d61a99ae834174d581c8.GitignoreRequestBuilder) {
    return icac773a834232ff9261f69ba5960a9a9c9a9a52a6372d61a99ae834174d581c8.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *ApiClient) Installation()(*i892fd18bf3f479d508c18121aaa79a47af092fb22bf2c016696f4ef49ac0f2df.InstallationRequestBuilder) {
    return i892fd18bf3f479d508c18121aaa79a47af092fb22bf2c016696f4ef49ac0f2df.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *ApiClient) Issues()(*i1790f60a6d24b15240ea65ca9183e125346ae1d30cb03d8b67ebb65a5e9af90a.IssuesRequestBuilder) {
    return i1790f60a6d24b15240ea65ca9183e125346ae1d30cb03d8b67ebb65a5e9af90a.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
func (m *ApiClient) Licenses()(*i4a11dbf1fac8370be6f5106b2269092133cf4ec314afd089fab27dc7dc1f3241.LicensesRequestBuilder) {
    return i4a11dbf1fac8370be6f5106b2269092133cf4ec314afd089fab27dc7dc1f3241.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
func (m *ApiClient) Markdown()(*i42234227323bb9085c39a8953a9cb5e6a70dc990771162276a9ca718fbba7d5e.MarkdownRequestBuilder) {
    return i42234227323bb9085c39a8953a9cb5e6a70dc990771162276a9ca718fbba7d5e.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_listing the marketplace_listing property
func (m *ApiClient) Marketplace_listing()(*i8eb05c4a8c40e9a66f7c3bd436ffaf5869db0e78726339a6c29d711921e3c8ca.Marketplace_listingRequestBuilder) {
    return i8eb05c4a8c40e9a66f7c3bd436ffaf5869db0e78726339a6c29d711921e3c8ca.NewMarketplace_listingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
func (m *ApiClient) Meta()(*i60d210be466537b7cfe7e84a79365b690725bffe8bdaff504696996406b87a8b.MetaRequestBuilder) {
    return i60d210be466537b7cfe7e84a79365b690725bffe8bdaff504696996406b87a8b.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
func (m *ApiClient) Networks()(*i5f31d8076704b2a33cf0d9496caffcad199649c2b9693da84305db274106d6c9.NetworksRequestBuilder) {
    return i5f31d8076704b2a33cf0d9496caffcad199649c2b9693da84305db274106d6c9.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
func (m *ApiClient) Notifications()(*id93ee658913eeb4bcf71cd66de862383ff3cdd07db79c518b345fc9dd4e57406.NotificationsRequestBuilder) {
    return id93ee658913eeb4bcf71cd66de862383ff3cdd07db79c518b345fc9dd4e57406.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
func (m *ApiClient) Octocat()(*i7238c553ea63f232858d04abc7aa42cf1c6703df7c6b4a89ed97696c881b1b86.OctocatRequestBuilder) {
    return i7238c553ea63f232858d04abc7aa42cf1c6703df7c6b4a89ed97696c881b1b86.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
func (m *ApiClient) Organizations()(*id498c92982d2f948ace86e8cde2d5fff55ba6cb718f448e213df19b59968ecd3.OrganizationsRequestBuilder) {
    return id498c92982d2f948ace86e8cde2d5fff55ba6cb718f448e213df19b59968ecd3.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
func (m *ApiClient) Orgs()(*i006ac19b46f7860b7359e6cc742a2213da1a512ec1c7d0aae6cf413f77880e18.OrgsRequestBuilder) {
    return i006ac19b46f7860b7359e6cc742a2213da1a512ec1c7d0aae6cf413f77880e18.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *ApiClient) Projects()(*i5406244bdf8038feac942b019f5d863c71d17da72a9cee044e224423a517481f.ProjectsRequestBuilder) {
    return i5406244bdf8038feac942b019f5d863c71d17da72a9cee044e224423a517481f.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
func (m *ApiClient) Rate_limit()(*i32b2c35b91803243459a19797935c406e08a7ac5f2e535b29c5578423b31588f.Rate_limitRequestBuilder) {
    return i32b2c35b91803243459a19797935c406e08a7ac5f2e535b29c5578423b31588f.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *ApiClient) Repos()(*iddc31237199800bb0c9999de06596e2979bdd4a3f890a253e6c5893742ea8d64.ReposRequestBuilder) {
    return iddc31237199800bb0c9999de06596e2979bdd4a3f890a253e6c5893742ea8d64.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
func (m *ApiClient) Repositories()(*ic5c472fad8099124bbac547779804efde5daea6ed48cfd69bad1391db8d32eaf.RepositoriesRequestBuilder) {
    return ic5c472fad8099124bbac547779804efde5daea6ed48cfd69bad1391db8d32eaf.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
func (m *ApiClient) Search()(*ibcd8a3e35b6ab970afb6dc75815c8ca8f5a5873760b244cfb82ad1ada7d36271.SearchRequestBuilder) {
    return ibcd8a3e35b6ab970afb6dc75815c8ca8f5a5873760b244cfb82ad1ada7d36271.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *ApiClient) Teams()(*i71e3d6b9d82a10e0c351b2793a30816189ce7b7e3809438d9637fddf19f2e4de.TeamsRequestBuilder) {
    return i71e3d6b9d82a10e0c351b2793a30816189ce7b7e3809438d9637fddf19f2e4de.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApiClientApiClientGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// User the user property
func (m *ApiClient) User()(*i036f8c76ca722afbbb72ebcdb51244c927aff7bca1cfebb950994aac3db6a206.UserRequestBuilder) {
    return i036f8c76ca722afbbb72ebcdb51244c927aff7bca1cfebb950994aac3db6a206.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
func (m *ApiClient) Users()(*i58d6850d1b0c5fc1569fd86453afdff88f2e933fdcb4cd121b1e20d9d527a0b8.UsersRequestBuilder) {
    return i58d6850d1b0c5fc1569fd86453afdff88f2e933fdcb4cd121b1e20d9d527a0b8.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
func (m *ApiClient) Versions()(*i0c0a073bba79314de6b968ee3ac963288d9a717dea40b7ef084afe188ea55d45.VersionsRequestBuilder) {
    return i0c0a073bba79314de6b968ee3ac963288d9a717dea40b7ef084afe188ea55d45.NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
func (m *ApiClient) Zen()(*i0a41a67c67557b401d607bec1017ba369edee449be6203262889bd4f6fe8018c.ZenRequestBuilder) {
    return i0a41a67c67557b401d607bec1017ba369edee449be6203262889bd4f6fe8018c.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
