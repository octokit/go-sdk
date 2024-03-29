package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"golang.org/x/sync/errgroup"

	// import "pkg/client" locally
	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/user"
	"github.com/octokit/go-sdk/pkg/github/user/repos"
	"github.com/octokit/go-sdk/pkg/handlers"
)

func main() {
	// as a consumer, how do i want to use the default API client and constructor?
	// tenets:
	// - minimal chaining
	// - functional options
	// - defaults include rate-limiting middleware and other sensible values (included without specification)
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("my-user-agent"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithRateLimitingMiddleware(),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithAuthorizationToken(os.Getenv("GITHUB_TOKEN")),
	)
	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	queryParams := &abstractions.DefaultQueryParameters{}
	requestConfig := &abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]{
		QueryParameters: queryParams,
	}
	zen, err := client.Zen().Get(context.Background(), requestConfig)
	if err != nil {
		fmt.Printf("error getting Zen: %v\n", err)
		return
	}
	fmt.Printf("%v\n", *zen)
}

// rateLimitTest can be used to hammer a local instance of GitHub
// to validate rate-limiting code
func rateLimitTest() {

	rateLimitHandler := handlers.NewRateLimitHandler()
	middlewares := kiotaHttp.GetDefaultMiddlewares()
	middlewares = append(middlewares, rateLimitHandler)
	netHttpClient := kiotaHttp.GetDefaultClient(middlewares...)

	tokenProvider := auth.NewTokenProvider(
		auth.WithAuthorizationToken(os.Getenv("GITHUB_TOKEN")),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)

	adapter, err := kiotaHttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(tokenProvider, nil, nil, netHttpClient)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}
	// adapter.SetBaseUrl("http://api.github.localhost:1024")

	client := github.NewApiClient(adapter)
	errGroup := &errgroup.Group{}
	for i := 0; i < 1500; i++ {
		errGroup.Go(func() error {
			viz := repos.ALL_GETVISIBILITYQUERYPARAMETERTYPE
			queryParams := &user.ReposRequestBuilderGetQueryParameters{
				Visibility: &viz,
			}
			requestConfig := &abstractions.RequestConfiguration[user.ReposRequestBuilderGetQueryParameters]{
				QueryParameters: queryParams,
			}
			_, err := client.User().Repos().Get(context.Background(), requestConfig)
			if err != nil {
				log.Fatalf("error getting repositories: %v", err)
				return err
			}

			// if len(repos) > 0 {
			// 	log.Printf("Repositories:\n")
			// 	for _, repo := range repos {
			// 		log.Printf("%v\n", *repo.GetFullName())
			// 	}
			// }
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		log.Fatalf("error from errgroup getting repositories: %v", err)
	} else {
		log.Printf("ran into no errors.")
	}
	// Output:
}
