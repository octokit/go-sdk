## How to contribute

Please note that this project is released with a [Contributor Code of Conduct](./CODE_OF_CONDUCT.md). By participating you agree to abide by its terms.

We appreciate you taking the time to contribute to Octokit or any of the projects in Octokit's ecosystem. Especially as a new contributor, you have a valuable perspective: you will find things confusing and run into problems that no longer occur to the maintainers. Please share them with us so we can make the experience for future contributors the best it can be.

Thank you 💖

There are two types of contributions: issues and pull requests (PRs). Issues are used to track bugs and feature requests, while PRs contribute new content to the repository.

## Creating an issue

Before you create a new issue:
1. Search the [project's issues](https://github.com/octokit/go-sdk/issues) to see if a similar issue already exists. If so, please add onto that issue rather than creating a new one.
1. If submitting a bug report, include steps and a minimal code sample that will reproduce the issue.
1. If submitting a feature request, please share the motivation for the new feature, what alternatives you considered, and any implementation suggestions.

## Creating a pull request

First, is your code a change to the generated source code present in `pikg/github`? If so, you'll want to go to [octokit/source-generator](https://github.com/octokit/source-generator) to modify the generation process.

If your changes do need to be made here, fork the repository. If you're not sure how to do so, please read [the linked docs](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo). Then clone your fork and `cd` to that directory locally.

Ensure that the `go build ./...` and `go test ./...` commands each succeed and result in no changes to the repository. Make your code changes (adding tests and documentation as necessary) and confirm the above validation steps still pass. If you'd like to debug the project, you can use VSCode's tooling to do so. When you're satisfied with your changes, follow [GitHub's docs](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request) to create your PR.
