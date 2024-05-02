package main

import (
	"context"
	"log"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/installation"
)

func main() {
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("my-user-agent"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		// pkg.WithAuthorizationToken(os.Getenv("GITHUB_TOKEN")),
		pkg.WithGitHubAppAuthentication("/home/kfcampbell/github/dev/go-sdk/kfcampbell-terraform-provider.2024-04-30.private-key.pem", 131977, 20570954),
	)

	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	queryParams := &installation.RepositoriesRequestBuilderGetQueryParameters{}
	requestConfig := &abs.RequestConfiguration[installation.RepositoriesRequestBuilderGetQueryParameters]{
		QueryParameters: queryParams,
	}
	repos, err := client.Installation().Repositories().Get(context.Background(), requestConfig)
	if err != nil {
		log.Fatalf("error getting repositories: %v", err)
	}

	if len(repos.GetRepositories()) > 0 {
		log.Printf("Repositories:\n")
		for _, repo := range repos.GetRepositories() {
			log.Printf("%v\n", *repo.GetFullName())
		}
	}
}
