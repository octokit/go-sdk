package authentication_test

import (
	"context"
	"log"
	"os"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/user"
	"github.com/octokit/go-sdk/pkg/github/user/repos"
	"github.com/octokit/go-sdk/pkg/handlers"
	"golang.org/x/sync/errgroup"
)

// Work in progress script to trigger rate limits
func ExampleApiClient_User_rateLimits() {
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
	adapter.SetBaseUrl("http://api.github.localhost:1024")

	client := github.NewApiClient(adapter)
	errGroup := &errgroup.Group{}
	for i := 0; i < 10000; i++ {
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
