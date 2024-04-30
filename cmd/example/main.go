package main

import (
	"context"
	"log"
	"time"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/user"
	"github.com/octokit/go-sdk/pkg/github/user/repos"
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

	viz := repos.ALL_GETVISIBILITYQUERYPARAMETERTYPE
	queryParams := &user.ReposRequestBuilderGetQueryParameters{
		Visibility: &viz,
	}
	requestConfig := &abstractions.RequestConfiguration[user.ReposRequestBuilderGetQueryParameters]{
		QueryParameters: queryParams,
	}
	repos, err := client.User().Repos().Get(context.Background(), requestConfig)
	if err != nil {
		log.Fatalf("error getting repositories: %v", err)
		// return err
	}

	if len(repos) > 0 {
		log.Printf("Repositories:\n")
		for _, repo := range repos {
			log.Printf("%v\n", *repo.GetFullName())
		}
	}

	// queryParams := &abs.DefaultQueryParameters{}
	// requestConfig := &abs.RequestConfiguration[abs.DefaultQueryParameters]{
	// 	QueryParameters: queryParams,
	// }
	// zen, err := client.Zen().Get(context.Background(), requestConfig)
	// if err != nil {
	// 	fmt.Printf("error getting Zen: %v\n", err)
	// 	return
	// }
	// fmt.Printf("%v\n", *zen)
}
