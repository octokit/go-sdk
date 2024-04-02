package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg"
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
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithAuthorizationToken(os.Getenv("GITHUB_TOKEN")),
	)

	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	queryParams := &abs.DefaultQueryParameters{}
	requestConfig := &abs.RequestConfiguration[abs.DefaultQueryParameters]{
		QueryParameters: queryParams,
	}
	zen, err := client.Zen().Get(context.Background(), requestConfig)
	if err != nil {
		fmt.Printf("error getting Zen: %v\n", err)
		return
	}
	fmt.Printf("%v\n", *zen)
}
