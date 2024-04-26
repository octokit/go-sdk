# github.com/octokit/go-sdk

An "alpha" version of a generated Go SDK from [GitHub's OpenAPI spec](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota).

## How do I use it?

See example client instantiations and requests in [example_test.go](pkg/example_test.go).

⚠️ **Note**: This SDK is not yet stable. Breaking changes may occur at any time.

### Authentication

Currently, this SDK supports both [Personal Access Tokens (classic)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#personal-access-tokens-classic) and [fine-grained Personal Access Tokens](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#fine-grained-personal-access-tokens).

Future work is planned for the SDK to support [GitHub Apps](https://docs.github.com/en/apps/overview) authentication as well.

## Why a generated SDK?

We want to...
1.  provide 100% coverage of the API in our SDK
2.  use this as a building block for future SDK tooling.

## Why Go?

We don't currently support a Go SDK and we wanted a narrow scope for the initial effort without worrying about cutting over users of an existing SDK.

## How can I report on my experience or issues with the SDK?

Please use this project's [issues](https://github.com/octokit/go-sdk/issues)!

## Releasing this project

Periodically (based on the frequency of [this workflow](https://github.com/octokit/source-generator/blob/main/.github/workflows/build-go.yml)), the source-generator repository will ingest the latest version of [GitHub's OpenAPI spec](https://github.com/github/rest-api-description) and generate a new version of this SDK. If there is a diff, a PR (similar to [this one](https://github.com/octokit/go-sdk/pull/22)) will be generated.

When reviewing the PR, analyze the diff and determine whether the changes are breaking (for which a major version number must be incremented), feature additions (for which a minor version number must be incremented), or bug fixes or docs changes (for which a patch number must be incremented). For more details about how to select an appropriate semantic version, see [semver.org](https://semver.org/). In many/most cases, due to the scale of GitHub's specification and the rate of change on it, the diff will be large and the changes will be technically breaking. This will mean incrementing a major version number. If a major version is being incremented, it must be changed in the [go.mod](./go.mod) file as described in [these docs](https://go.dev/doc/modules/release-workflow#breaking).

When changes are analyzed, change the PR title appropriately (see [this PR](https://github.com/octokit/go-sdk/pull/40) for an example) and merge it. Then go to [repository releases](https://github.com/octokit/go-sdk/releases), tag the release with the chosen version, title it with the chosen version, use the "Generate release notes" button to see what PRs will be included in the release, and manually edit the release notes grouping the changes under the headings `Features`, `Fixes`, `Maintenance`, and `Documentation` when appropriate. After clicking "Publish Release", the new version will be available for use!
