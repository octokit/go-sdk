package authentication_test

import (
	"context"
	"log"
	"os"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	http "github.com/microsoft/kiota-http-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/user"
	"github.com/octokit/go-sdk/pkg/github/user/repos"
	"github.com/octokit/go-sdk/pkg/handlers"
)

// Work in progress script to trigger primary rate limits and then stop execution.
func ExampleApiClient_User_rateLimits() {
	rateLimitHandler := handlers.NewRateLimitHandler()
	middlewares := kiotaHttp.GetDefaultMiddlewares()
	middlewares = append(middlewares, rateLimitHandler)
	netHttpClient := kiotaHttp.GetDefaultClient(middlewares...)

	tokenProvider := auth.NewTokenProvider(
		auth.WithAuthorizationToken(os.Getenv("GITHUB_TOKEN")),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)

	adapter, err := http.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(tokenProvider, nil, nil, netHttpClient)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}
	// adapter.SetBaseUrl("https://kfcampbell.some-enterprise-instance.github.com")

	client := github.NewApiClient(adapter)

	viz := repos.ALL_GETVISIBILITYQUERYPARAMETERTYPE
	var page int32 = 0
	queryParams := &user.ReposRequestBuilderGetQueryParameters{
		Visibility: &viz,
		Page:       &page,
	}
	requestConfig := &abstractions.RequestConfiguration[user.ReposRequestBuilderGetQueryParameters]{
		QueryParameters: queryParams,
	}

	// page through all repos
	repos, err := client.User().Repos().Get(context.Background(), requestConfig)
	if err != nil {
		log.Fatalf("error getting repositories: %v", err)
	}

	for len(repos) > 0 && err == nil {
		log.Printf("Repositories:\n")
		for _, repo := range repos {
			log.Printf("%v\n", *repo.GetFullName())
		}
		page++
		queryParams.Page = &page
		requestConfig.QueryParameters = queryParams
		repos, err = client.User().Repos().Get(context.Background(), requestConfig)
	}
	// Output:
}
