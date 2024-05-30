package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/installation"
)

func main() {
	installationID, err := strconv.ParseInt(os.Getenv("INSTALLATION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("error parsing installation ID from string to int64: %v", err)
	}

	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("my-user-agent"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithGitHubAppAuthentication(os.Getenv("PATH_TO_PEM_FILE"), os.Getenv("CLIENT_ID"), installationID),
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
