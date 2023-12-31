# github.com/octokit/go-sdk

An "alpha" version of a generated Go SDK from [GitHub's OpenAPI spec](https://github.com/github/rest-api-description), built on [Kiota](https://github.com/microsoft/kiota).

## How do I use it?

See example client instantiations and requests in [example_test.go](pkg/authentication/example_test.go).

⚠️ **Note**: This SDK is not yet stable. Breaking changes may occur at any time.

## Why a generated SDK?

We want to...
1.  provide 100% coverage of the API in our SDK
2.  use this as a building block for future SDK tooling.

## Why Go?

We don't currently support a Go SDK and we wanted a narrow scope for the initial effort without worrying about cutting over users of an existing SDK.

## How can I report on my experience or issues with the SDK?

Please use this project's [issues](https://github.com/octokit/go-sdk/issues)!
