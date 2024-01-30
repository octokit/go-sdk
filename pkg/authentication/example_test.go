package authentication_test

import (
	"context"
	"fmt"
	"log"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	http "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/octocat"
)

// ExampleApiClient_Octocat shows how to initialize an unauthenticated client
// and make a simple API request.
func ExampleApiClient_Octocat() {
	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)
	adapter, err := http.NewNetHttpRequestAdapter(tokenProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v", err)
	}

	client := github.NewApiClient(adapter)

	s := "Salutations"

	// create headers that accept json back; GitHub's OpenAPI definition says
	// octet-stream but that's not actually what the API returns in this case
	headers := abstractions.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")

	octocatRequestConfig := &abstractions.RequestConfiguration[octocat.OctocatRequestBuilderGetQueryParameters]{
		QueryParameters: &octocat.OctocatRequestBuilderGetQueryParameters{
			S: &s,
		},
		Headers: headers,
	}

	cat, err := client.Octocat().Get(context.Background(), octocatRequestConfig)
	if err != nil {
		log.Fatalf("error getting octocat: %v", err)
	}
	fmt.Printf("%v\n", string(cat))
	// Output:
	// MMM.           .MMM
	//                MMMMMMMMMMMMMMMMMMM
	//                MMMMMMMMMMMMMMMMMMM      _____________
	//               MMMMMMMMMMMMMMMMMMMMM    |             |
	//              MMMMMMMMMMMMMMMMMMMMMMM   | Salutations |
	//             MMMMMMMMMMMMMMMMMMMMMMMM   |_   _________|
	//             MMMM::- -:::::::- -::MMMM    |/
	//              MM~:~ 00~:::::~ 00~:~MM
	//         .. MMMMM::.00:::+:::.00::MMMMM ..
	//               .MM::::: ._. :::::MM.
	//                  MMMM;:::::;MMMM
	//           -MM        MMMMMMM
	//           ^  M+     MMMMMMMMM
	//               MMMMMMM MM MM MM
	//                    MM MM MM MM
	//                    MM MM MM MM
	//                 .~~MM~MM~MM~MM~~.
	//              ~~~~MM:~MM~~~MM~:MM~~~~
	//             ~~~~~~==~==~~~==~==~~~~~~
	//              ~~~~~~==~==~==~==~~~~~~
	//                  :~==~==~==~==~~
}
