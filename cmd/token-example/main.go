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
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("my-user-agent"),
		pkg.WithRequestTimeout(5*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
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
		log.Fatalf("error getting repositories: %v", err)
	}
	fmt.Printf("GitHub Zen principle: %v\n", *zen)
}
