package octokit

import (
    "context"
    i00f1fed9ef7cd71bdf3e75aeef327685883fab64913202cd0139764d217aaff2 "github.com/octokit/go-sdk/github/octokit/gitignore"
    i012e0ce26fed2d41789483a15c5c138e39c09d7138903797fa49ac262560f5df "github.com/octokit/go-sdk/github/octokit/assignments"
    i02ee8c0c88594751b60d64ed7fcdb000fdaad97def006623cbdb4f69f4b32d94 "github.com/octokit/go-sdk/github/octokit/advisories"
    i0630a5ec64f53746aa0214e4e80d70b6c40564fdc4c6b1e5638c35922cba31e1 "github.com/octokit/go-sdk/github/octokit/orgs"
    i082acd51e55dfd68ccfffd788e8960fd8e3bda49ae4cdc128e63d99ce442dee7 "github.com/octokit/go-sdk/github/octokit/networks"
    i181e64ab2934af4de43a294ded7ef34b081a50c35ba40f4cec6a76a5822989d1 "github.com/octokit/go-sdk/github/octokit/organizations"
    i1dfd9332d37df20f4bcd14ad579e3b1cffe9663d10ab09245466823b4f3012b2 "github.com/octokit/go-sdk/github/octokit/installation"
    i262e7bac7d437f187308d33ff0b0896e42bda1aeb406ee592579018ad9438b08 "github.com/octokit/go-sdk/github/octokit/meta"
    i2c35b84e5dfa49f713209a013e80a7385beaf706e6e556aecf2e399fe207c78e "github.com/octokit/go-sdk/github/octokit/appmanifests"
    i3058d03d797c773039f833ab56d4ca057161326ee2f7abed09b74ff73b4ae7eb "github.com/octokit/go-sdk/github/octokit/licenses"
    i38d9bbc7d8ce8ccc6b96b9e050427bbdfb0f0e2c490a5b696de8c1700d5873e4 "github.com/octokit/go-sdk/github/octokit/rate_limit"
    i44f21d2b6915e282dade6494e0dd5d54c66e851206bc3d6538b7c0409e1b068a "github.com/octokit/go-sdk/github/octokit/enterprises"
    i5095dd2639d5bc5f4e034eff6e6d0f6c9abd91b7305e13d5f2b03c3952784b2d "github.com/octokit/go-sdk/github/octokit/codes_of_conduct"
    i5158455fbddcd6c896099ba1b5d12464896ded2373b500b4f608501c9bbf6043 "github.com/octokit/go-sdk/github/octokit/gists"
    i5a2e1bd682f7fd3d317f2f71dad65e31642f2c3320cddf0b774df79d329e8b15 "github.com/octokit/go-sdk/github/octokit/zen"
    i627571c2e22f9314da80ae4c23623e706d31b8776afca5e44a7a102a6d472571 "github.com/octokit/go-sdk/github/octokit/octocat"
    i64a317472850c3e4f5ef3131ab2a9a6f13db1588332a5c1dcd9f1a196a7eb8c7 "github.com/octokit/go-sdk/github/octokit/projects"
    i6c37bff252b47d15034dd4a14748f8549f8828890fe90f4959308d33347e8517 "github.com/octokit/go-sdk/github/octokit/user"
    i6e575246fc177c2ef9715eb0616ab13e30864c150a617d034a0ea090847e4cbe "github.com/octokit/go-sdk/github/octokit/emojis"
    i7efaf73b4de581fad7766e10d6635bf593174afa4b8f1d434b52818d3f587a75 "github.com/octokit/go-sdk/github/octokit/app"
    i833fa9bc6593661f71796b7d84d648c0d65f8408b784ab7cd6eebab510ae1315 "github.com/octokit/go-sdk/github/octokit/markdown"
    i85c5abd6c240949d53522ba93d9bd02e9eb0476edc84b6b5594555c4562d37b8 "github.com/octokit/go-sdk/github/octokit/teams"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i8c195517ba5951b8b25203ea359263da58c5106fe9fc3e298b0fd10433c1b73a "github.com/octokit/go-sdk/github/octokit/search"
    ia5719cf22df54e3f653a789d4bb3cfbf26867f5d852bae050b776a3627a20804 "github.com/octokit/go-sdk/github/octokit/apps"
    ia9291c1caee63132323a299d49592eaca763c9e77ad428426f40cb6259aa6d96 "github.com/octokit/go-sdk/github/octokit/applications"
    iaffade5a83b840a1cef15c7bea8f813cbe645cb139c446bfa96d701a293886ea "github.com/octokit/go-sdk/github/octokit/repositories"
    ibbd3ded5ab35a41a71d583e10c43d0f095f9726b5b82bba8b5a25e9f933474b3 "github.com/octokit/go-sdk/github/octokit/repos"
    ic78807f3d372c667fba52904ce880ea4a1037de1fbe2df9e1063e9824975c7a0 "github.com/octokit/go-sdk/github/octokit/events"
    icbe907cd68e180423ba744b18cd54d6f1a4504ebbabbc646c47cd2e53fe3b710 "github.com/octokit/go-sdk/github/octokit/notifications"
    ice4dff0648d8b9ca210814f95ef4895f05a759326d8f73affda67a6182c86bef "github.com/octokit/go-sdk/github/octokit/classrooms"
    id4564c761335fadb4fd125944822dfe1b721afd322a3fea1696f698e2c45d082 "github.com/octokit/go-sdk/github/octokit/issues"
    id849601c7588aabc5e596130497dfee29f50c45e6c5523e22c139610f6c753ca "github.com/octokit/go-sdk/github/octokit/users"
    idc1a675f538c66fdc49077be1d93fd555148dcf6f7a03f36272f04a40b73ef73 "github.com/octokit/go-sdk/github/octokit/feeds"
    ie577812b0036060606fad9c9b7baa97e00057feccfa313626b08117c252194d7 "github.com/octokit/go-sdk/github/octokit/versions"
    if30a1452149cfdf73a767d93b983bc75ceafe8e5c3a2145b5d091ecd0c1e9ded "github.com/octokit/go-sdk/github/octokit/marketplace_listing"
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
func (m *ApiClient) Advisories()(*i02ee8c0c88594751b60d64ed7fcdb000fdaad97def006623cbdb4f69f4b32d94.AdvisoriesRequestBuilder) {
    return i02ee8c0c88594751b60d64ed7fcdb000fdaad97def006623cbdb4f69f4b32d94.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
func (m *ApiClient) App()(*i7efaf73b4de581fad7766e10d6635bf593174afa4b8f1d434b52818d3f587a75.AppRequestBuilder) {
    return i7efaf73b4de581fad7766e10d6635bf593174afa4b8f1d434b52818d3f587a75.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
func (m *ApiClient) Applications()(*ia9291c1caee63132323a299d49592eaca763c9e77ad428426f40cb6259aa6d96.ApplicationsRequestBuilder) {
    return ia9291c1caee63132323a299d49592eaca763c9e77ad428426f40cb6259aa6d96.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
func (m *ApiClient) AppManifests()(*i2c35b84e5dfa49f713209a013e80a7385beaf706e6e556aecf2e399fe207c78e.AppManifestsRequestBuilder) {
    return i2c35b84e5dfa49f713209a013e80a7385beaf706e6e556aecf2e399fe207c78e.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
func (m *ApiClient) Apps()(*ia5719cf22df54e3f653a789d4bb3cfbf26867f5d852bae050b776a3627a20804.AppsRequestBuilder) {
    return ia5719cf22df54e3f653a789d4bb3cfbf26867f5d852bae050b776a3627a20804.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments the assignments property
func (m *ApiClient) Assignments()(*i012e0ce26fed2d41789483a15c5c138e39c09d7138903797fa49ac262560f5df.AssignmentsRequestBuilder) {
    return i012e0ce26fed2d41789483a15c5c138e39c09d7138903797fa49ac262560f5df.NewAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classrooms the classrooms property
func (m *ApiClient) Classrooms()(*ice4dff0648d8b9ca210814f95ef4895f05a759326d8f73affda67a6182c86bef.ClassroomsRequestBuilder) {
    return ice4dff0648d8b9ca210814f95ef4895f05a759326d8f73affda67a6182c86bef.NewClassroomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
func (m *ApiClient) Codes_of_conduct()(*i5095dd2639d5bc5f4e034eff6e6d0f6c9abd91b7305e13d5f2b03c3952784b2d.Codes_of_conductRequestBuilder) {
    return i5095dd2639d5bc5f4e034eff6e6d0f6c9abd91b7305e13d5f2b03c3952784b2d.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ApiClient) Emojis()(*i6e575246fc177c2ef9715eb0616ab13e30864c150a617d034a0ea090847e4cbe.EmojisRequestBuilder) {
    return i6e575246fc177c2ef9715eb0616ab13e30864c150a617d034a0ea090847e4cbe.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
func (m *ApiClient) Enterprises()(*i44f21d2b6915e282dade6494e0dd5d54c66e851206bc3d6538b7c0409e1b068a.EnterprisesRequestBuilder) {
    return i44f21d2b6915e282dade6494e0dd5d54c66e851206bc3d6538b7c0409e1b068a.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
func (m *ApiClient) Events()(*ic78807f3d372c667fba52904ce880ea4a1037de1fbe2df9e1063e9824975c7a0.EventsRequestBuilder) {
    return ic78807f3d372c667fba52904ce880ea4a1037de1fbe2df9e1063e9824975c7a0.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
func (m *ApiClient) Feeds()(*idc1a675f538c66fdc49077be1d93fd555148dcf6f7a03f36272f04a40b73ef73.FeedsRequestBuilder) {
    return idc1a675f538c66fdc49077be1d93fd555148dcf6f7a03f36272f04a40b73ef73.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *ApiClientApiClientApiClientGetRequestConfiguration)(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.Rootable), nil
}
// Gists the gists property
func (m *ApiClient) Gists()(*i5158455fbddcd6c896099ba1b5d12464896ded2373b500b4f608501c9bbf6043.GistsRequestBuilder) {
    return i5158455fbddcd6c896099ba1b5d12464896ded2373b500b4f608501c9bbf6043.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
func (m *ApiClient) Gitignore()(*i00f1fed9ef7cd71bdf3e75aeef327685883fab64913202cd0139764d217aaff2.GitignoreRequestBuilder) {
    return i00f1fed9ef7cd71bdf3e75aeef327685883fab64913202cd0139764d217aaff2.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
func (m *ApiClient) Installation()(*i1dfd9332d37df20f4bcd14ad579e3b1cffe9663d10ab09245466823b4f3012b2.InstallationRequestBuilder) {
    return i1dfd9332d37df20f4bcd14ad579e3b1cffe9663d10ab09245466823b4f3012b2.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
func (m *ApiClient) Issues()(*id4564c761335fadb4fd125944822dfe1b721afd322a3fea1696f698e2c45d082.IssuesRequestBuilder) {
    return id4564c761335fadb4fd125944822dfe1b721afd322a3fea1696f698e2c45d082.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
func (m *ApiClient) Licenses()(*i3058d03d797c773039f833ab56d4ca057161326ee2f7abed09b74ff73b4ae7eb.LicensesRequestBuilder) {
    return i3058d03d797c773039f833ab56d4ca057161326ee2f7abed09b74ff73b4ae7eb.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
func (m *ApiClient) Markdown()(*i833fa9bc6593661f71796b7d84d648c0d65f8408b784ab7cd6eebab510ae1315.MarkdownRequestBuilder) {
    return i833fa9bc6593661f71796b7d84d648c0d65f8408b784ab7cd6eebab510ae1315.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_listing the marketplace_listing property
func (m *ApiClient) Marketplace_listing()(*if30a1452149cfdf73a767d93b983bc75ceafe8e5c3a2145b5d091ecd0c1e9ded.Marketplace_listingRequestBuilder) {
    return if30a1452149cfdf73a767d93b983bc75ceafe8e5c3a2145b5d091ecd0c1e9ded.NewMarketplace_listingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
func (m *ApiClient) Meta()(*i262e7bac7d437f187308d33ff0b0896e42bda1aeb406ee592579018ad9438b08.MetaRequestBuilder) {
    return i262e7bac7d437f187308d33ff0b0896e42bda1aeb406ee592579018ad9438b08.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
func (m *ApiClient) Networks()(*i082acd51e55dfd68ccfffd788e8960fd8e3bda49ae4cdc128e63d99ce442dee7.NetworksRequestBuilder) {
    return i082acd51e55dfd68ccfffd788e8960fd8e3bda49ae4cdc128e63d99ce442dee7.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
func (m *ApiClient) Notifications()(*icbe907cd68e180423ba744b18cd54d6f1a4504ebbabbc646c47cd2e53fe3b710.NotificationsRequestBuilder) {
    return icbe907cd68e180423ba744b18cd54d6f1a4504ebbabbc646c47cd2e53fe3b710.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
func (m *ApiClient) Octocat()(*i627571c2e22f9314da80ae4c23623e706d31b8776afca5e44a7a102a6d472571.OctocatRequestBuilder) {
    return i627571c2e22f9314da80ae4c23623e706d31b8776afca5e44a7a102a6d472571.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
func (m *ApiClient) Organizations()(*i181e64ab2934af4de43a294ded7ef34b081a50c35ba40f4cec6a76a5822989d1.OrganizationsRequestBuilder) {
    return i181e64ab2934af4de43a294ded7ef34b081a50c35ba40f4cec6a76a5822989d1.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
func (m *ApiClient) Orgs()(*i0630a5ec64f53746aa0214e4e80d70b6c40564fdc4c6b1e5638c35922cba31e1.OrgsRequestBuilder) {
    return i0630a5ec64f53746aa0214e4e80d70b6c40564fdc4c6b1e5638c35922cba31e1.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
func (m *ApiClient) Projects()(*i64a317472850c3e4f5ef3131ab2a9a6f13db1588332a5c1dcd9f1a196a7eb8c7.ProjectsRequestBuilder) {
    return i64a317472850c3e4f5ef3131ab2a9a6f13db1588332a5c1dcd9f1a196a7eb8c7.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
func (m *ApiClient) Rate_limit()(*i38d9bbc7d8ce8ccc6b96b9e050427bbdfb0f0e2c490a5b696de8c1700d5873e4.Rate_limitRequestBuilder) {
    return i38d9bbc7d8ce8ccc6b96b9e050427bbdfb0f0e2c490a5b696de8c1700d5873e4.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
func (m *ApiClient) Repos()(*ibbd3ded5ab35a41a71d583e10c43d0f095f9726b5b82bba8b5a25e9f933474b3.ReposRequestBuilder) {
    return ibbd3ded5ab35a41a71d583e10c43d0f095f9726b5b82bba8b5a25e9f933474b3.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
func (m *ApiClient) Repositories()(*iaffade5a83b840a1cef15c7bea8f813cbe645cb139c446bfa96d701a293886ea.RepositoriesRequestBuilder) {
    return iaffade5a83b840a1cef15c7bea8f813cbe645cb139c446bfa96d701a293886ea.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
func (m *ApiClient) Search()(*i8c195517ba5951b8b25203ea359263da58c5106fe9fc3e298b0fd10433c1b73a.SearchRequestBuilder) {
    return i8c195517ba5951b8b25203ea359263da58c5106fe9fc3e298b0fd10433c1b73a.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
func (m *ApiClient) Teams()(*i85c5abd6c240949d53522ba93d9bd02e9eb0476edc84b6b5594555c4562d37b8.TeamsRequestBuilder) {
    return i85c5abd6c240949d53522ba93d9bd02e9eb0476edc84b6b5594555c4562d37b8.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// User the user property
func (m *ApiClient) User()(*i6c37bff252b47d15034dd4a14748f8549f8828890fe90f4959308d33347e8517.UserRequestBuilder) {
    return i6c37bff252b47d15034dd4a14748f8549f8828890fe90f4959308d33347e8517.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
func (m *ApiClient) Users()(*id849601c7588aabc5e596130497dfee29f50c45e6c5523e22c139610f6c753ca.UsersRequestBuilder) {
    return id849601c7588aabc5e596130497dfee29f50c45e6c5523e22c139610f6c753ca.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
func (m *ApiClient) Versions()(*ie577812b0036060606fad9c9b7baa97e00057feccfa313626b08117c252194d7.VersionsRequestBuilder) {
    return ie577812b0036060606fad9c9b7baa97e00057feccfa313626b08117c252194d7.NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
func (m *ApiClient) Zen()(*i5a2e1bd682f7fd3d317f2f71dad65e31642f2c3320cddf0b774df79d329e8b15.ZenRequestBuilder) {
    return i5a2e1bd682f7fd3d317f2f71dad65e31642f2c3320cddf0b774df79d329e8b15.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
