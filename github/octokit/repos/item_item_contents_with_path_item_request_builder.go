package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035 "github.com/octokit/go-sdk/github/octokit/models"
)

// ItemItemContentsWithPathItemRequestBuilder builds and executes requests for operations under \repos\{repos-id}\{Owner-id}\contents\{path}
type ItemItemContentsWithPathItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemContentsWithPathItemRequestBuilderGetQueryParameters gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/overview/media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/overview/media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/git/trees#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download. Size limits:If the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported. If the content is a directory:The response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule". If the content is a symlink: If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself. If the content is a submodule:The `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
type ItemItemContentsWithPathItemRequestBuilderGetQueryParameters struct {
    // The name of the commit/branch/tag. Default: the repository’s default branch.
    Ref *string `uriparametername:"ref"`
}
// ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemContentsWithPathItemRequestBuilderGetQueryParameters
}
// ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithPathGetResponse composed type wrapper for classes contentFile, contentSubmodule, contentSymlink, WithPathGetResponseMember1
type WithPathGetResponse struct {
    // Composed type representation for type contentFile
    contentFile i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable
    // Composed type representation for type contentSubmodule
    contentSubmodule i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable
    // Composed type representation for type contentSymlink
    contentSymlink i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable
    // Composed type representation for type WithPathGetResponseMember1
    withPathGetResponseMember1 []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able
}
// NewWithPathGetResponse instantiates a new WithPathGetResponse and sets the default values.
func NewWithPathGetResponse()(*WithPathGetResponse) {
    m := &WithPathGetResponse{
    }
    return m
}
// CreateWithPathGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithPathGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithPathGetResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateWithPathGetResponseMember1FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)
            }
        }
        result.SetWithPathGetResponseMember1(cast)
    }
    return result, nil
}
// GetContentFile gets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathGetResponse) GetContentFile()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable) {
    return m.contentFile
}
// GetContentSubmodule gets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathGetResponse) GetContentSubmodule()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable) {
    return m.contentSubmodule
}
// GetContentSymlink gets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathGetResponse) GetContentSymlink()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable) {
    return m.contentSymlink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WithPathGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *WithPathGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetWithPathGetResponseMember1 gets the WithPathGetResponseMember1 property value. Composed type representation for type WithPathGetResponseMember1
func (m *WithPathGetResponse) GetWithPathGetResponseMember1()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able) {
    return m.withPathGetResponseMember1
}
// Serialize serializes information the current object
func (m *WithPathGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentFile() != nil {
        err := writer.WriteObjectValue("", m.GetContentFile())
        if err != nil {
            return err
        }
    } else if m.GetContentSubmodule() != nil {
        err := writer.WriteObjectValue("", m.GetContentSubmodule())
        if err != nil {
            return err
        }
    } else if m.GetContentSymlink() != nil {
        err := writer.WriteObjectValue("", m.GetContentSymlink())
        if err != nil {
            return err
        }
    } else if m.GetWithPathGetResponseMember1() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWithPathGetResponseMember1()))
        for i, v := range m.GetWithPathGetResponseMember1() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentFile sets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathGetResponse) SetContentFile(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)() {
    m.contentFile = value
}
// SetContentSubmodule sets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathGetResponse) SetContentSubmodule(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)() {
    m.contentSubmodule = value
}
// SetContentSymlink sets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathGetResponse) SetContentSymlink(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)() {
    m.contentSymlink = value
}
// SetWithPathGetResponseMember1 sets the WithPathGetResponseMember1 property value. Composed type representation for type WithPathGetResponseMember1
func (m *WithPathGetResponse) SetWithPathGetResponseMember1(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)() {
    m.withPathGetResponseMember1 = value
}
// WithPathResponse composed type wrapper for classes contentFile, contentSubmodule, contentSymlink, WithPathGetResponseMember1
type WithPathResponse struct {
    // Composed type representation for type contentFile
    contentFile i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable
    // Composed type representation for type contentSubmodule
    contentSubmodule i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable
    // Composed type representation for type contentSymlink
    contentSymlink i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable
    // Composed type representation for type WithPathGetResponseMember1
    withPathGetResponseMember1 []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able
}
// NewWithPathResponse instantiates a new WithPathResponse and sets the default values.
func NewWithPathResponse()(*WithPathResponse) {
    m := &WithPathResponse{
    }
    return m
}
// CreateWithPathResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWithPathResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithPathResponse()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetCollectionOfObjectValues(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateWithPathGetResponseMember1FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)
            }
        }
        result.SetWithPathGetResponseMember1(cast)
    }
    return result, nil
}
// GetContentFile gets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathResponse) GetContentFile()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable) {
    return m.contentFile
}
// GetContentSubmodule gets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathResponse) GetContentSubmodule()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable) {
    return m.contentSubmodule
}
// GetContentSymlink gets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathResponse) GetContentSymlink()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable) {
    return m.contentSymlink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WithPathResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
func (m *WithPathResponse) GetIsComposedType()(bool) {
    return true
}
// GetWithPathGetResponseMember1 gets the WithPathGetResponseMember1 property value. Composed type representation for type WithPathGetResponseMember1
func (m *WithPathResponse) GetWithPathGetResponseMember1()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able) {
    return m.withPathGetResponseMember1
}
// Serialize serializes information the current object
func (m *WithPathResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentFile() != nil {
        err := writer.WriteObjectValue("", m.GetContentFile())
        if err != nil {
            return err
        }
    } else if m.GetContentSubmodule() != nil {
        err := writer.WriteObjectValue("", m.GetContentSubmodule())
        if err != nil {
            return err
        }
    } else if m.GetContentSymlink() != nil {
        err := writer.WriteObjectValue("", m.GetContentSymlink())
        if err != nil {
            return err
        }
    } else if m.GetWithPathGetResponseMember1() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWithPathGetResponseMember1()))
        for i, v := range m.GetWithPathGetResponseMember1() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentFile sets the contentFile property value. Composed type representation for type contentFile
func (m *WithPathResponse) SetContentFile(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)() {
    m.contentFile = value
}
// SetContentSubmodule sets the contentSubmodule property value. Composed type representation for type contentSubmodule
func (m *WithPathResponse) SetContentSubmodule(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)() {
    m.contentSubmodule = value
}
// SetContentSymlink sets the contentSymlink property value. Composed type representation for type contentSymlink
func (m *WithPathResponse) SetContentSymlink(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)() {
    m.contentSymlink = value
}
// SetWithPathGetResponseMember1 sets the WithPathGetResponseMember1 property value. Composed type representation for type WithPathGetResponseMember1
func (m *WithPathResponse) SetWithPathGetResponseMember1(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)() {
    m.withPathGetResponseMember1 = value
}
// WithPathGetResponseable 
type WithPathGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFile()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)
    GetContentSubmodule()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)
    GetContentSymlink()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)
    GetWithPathGetResponseMember1()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)
    SetContentFile(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)()
    SetContentSubmodule(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)()
    SetContentSymlink(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)()
    SetWithPathGetResponseMember1(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)()
}
// WithPathResponseable 
type WithPathResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFile()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)
    GetContentSubmodule()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)
    GetContentSymlink()(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)
    GetWithPathGetResponseMember1()([]i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)
    SetContentFile(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentFileable)()
    SetContentSubmodule(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSubmoduleable)()
    SetContentSymlink(value i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.ContentSymlinkable)()
    SetWithPathGetResponseMember1(value []i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.WithPathGetResponseMember1able)()
}
// NewItemItemContentsWithPathItemRequestBuilderInternal instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    m := &ItemItemContentsWithPathItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{repos%2Did}/{Owner%2Did}/contents/{path}{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemContentsWithPathItemRequestBuilder instantiates a new WithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemContentsWithPathItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/repos/contents/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#delete-a-file
func (m *ItemItemContentsWithPathItemRequestBuilder) Delete(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.FileCommitable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "409": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
        "503": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateFileCommit503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.FileCommitable), nil
}
// Get gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/overview/media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/overview/media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/git/trees#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download. Size limits:If the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported. If the content is a directory:The response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule". If the content is a symlink: If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself. If the content is a submodule:The `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
// Deprecated: This method is obsolete. Use GetAsWithPathGetResponse instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#get-repository-content
func (m *ItemItemContentsWithPathItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration)(WithPathResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithPathResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithPathResponseable), nil
}
// GetAsWithPathGetResponse gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/overview/media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/overview/media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/git/trees#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download. Size limits:If the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported. If the content is a directory:The response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule". If the content is a symlink: If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself. If the content is a submodule:The `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#get-repository-content
func (m *ItemItemContentsWithPathItemRequestBuilder) GetAsWithPathGetResponse(ctx context.Context, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration)(WithPathGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithPathGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithPathGetResponseable), nil
}
// Put creates a new file or replaces an existing file in a repository. You must authenticate using an access token with the `repo` scope to use this endpoint. If you want to modify files in the `.github/workflows` directory, you must authenticate using an access token with the `workflow` scope.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/repos/contents/#delete-a-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#create-or-update-file-contents
func (m *ItemItemContentsWithPathItemRequestBuilder) Put(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration)(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.FileCommitable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "409": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateBasicErrorFromDiscriminatorValue,
        "422": i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i000736ae6dd74f01081193e4f903216bc2bd2954ed818433b986f45d581ed035.FileCommitable), nil
}
// ToDeleteRequestInformation deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/repos/contents/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation gets the contents of a file or directory in a repository. Specify the file path or directory in `:path`. If you omit`:path`, you will receive the contents of the repository's root directory. See the description below regarding what the API response includes for directories. Files and symlinks support [a custom media type](https://docs.github.com/rest/overview/media-types) forretrieving the raw content or rendered HTML (when supported). All content types support [a custom mediatype](https://docs.github.com/rest/overview/media-types) to ensure the content is returned in a consistentobject format.**Notes**:*   To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).*   This API has an upper limit of 1,000 files for a directory. If you need to retrieve more files, use the [Git TreesAPI](https://docs.github.com/rest/git/trees#get-a-tree). *  Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download. Size limits:If the requested file's size is:* 1 MB or smaller: All features of this endpoint are supported.* Between 1-100 MB: Only the `raw` or `object` [custom media types](https://docs.github.com/rest/repos/contents#custom-media-types-for-repository-contents) are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an empty string and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type. * Greater than 100 MB: This endpoint is not supported. If the content is a directory:The response will be an array of objects, one object for each item in the directory.When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value_should_ be "submodule". This behavior exists in API v3 [for backwards compatibility purposes](https://git.io/v1YCW).In the next major version of the API, the type will be returned as "submodule". If the content is a symlink: If the requested `:path` points to a symlink, and the symlink's target is a normal file in the repository, then theAPI responds with the content of the file (in the format shown in the example. Otherwise, the API responds with an object describing the symlink itself. If the content is a submodule:The `submodule_git_url` identifies the location of the submodule repository, and the `sha` identifies a specificcommit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks outthe submodule at that specific commit.If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and thegithub.com URLs (`html_url` and `_links["html"]`) will have null values.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation creates a new file or replaces an existing file in a repository. You must authenticate using an access token with the `repo` scope to use this endpoint. If you want to modify files in the `.github/workflows` directory, you must authenticate using an access token with the `workflow` scope.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/repos/contents/#delete-a-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
func (m *ItemItemContentsWithPathItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *ItemItemContentsWithPathItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemContentsWithPathItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemContentsWithPathItemRequestBuilder) {
    return NewItemItemContentsWithPathItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
