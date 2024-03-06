package authentication_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	auth "github.com/octokit/go-sdk/pkg/authentication"
	"github.com/octokit/go-sdk/pkg/github"
	"github.com/octokit/go-sdk/pkg/github/octocat"
	"github.com/octokit/go-sdk/pkg/handlers"
	"github.com/octokit/go-sdk/pkg/transport"
)

// ExampleApiClient_Octocat shows how to initialize an unauthenticated client
// and make a simple API request.
func ExampleApiClient_Octocat() {
	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)
	adapter, err := kiotaHttp.NewNetHttpRequestAdapter(tokenProvider)
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

func Example_ResponseHandlerForRateLimiting() {
	// this responseHandler function panics:
	// panic: interface conversion: *http.Response is not serialization.Parsable: missing method GetFieldDeserializers
	// not sure if i'm mishandling something in this totally non-implemented function or what
	responseHandler := func(response interface{}, errorMappings abstractions.ErrorMappings) (interface{}, error) {
		// handle rate-limited responses here
		fmt.Printf("%v\n", response)
		if response == nil {
			return nil, nil
		}
		return response, nil
	}
	responseHandlerOption := abstractions.NewRequestHandlerOption()
	responseHandlerOption.SetResponseHandler(responseHandler)

	options := make([]abstractions.RequestOption, 0)
	options = append(options, responseHandlerOption)

	retryHandlerOption := &kiotaHttp.RetryHandlerOptions{
		MaxRetries: 3,
	}
	options = append(options, retryHandlerOption)

	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)

	adapter, err := kiotaHttp.NewNetHttpRequestAdapter(tokenProvider)
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
		Options: options,
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

func Example_HttpClientForRateLimiting() {
	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)

	transport := transport.NewRateLimitTransport(http.DefaultTransport,
		transport.WithWriteDelay(1*time.Second),
		transport.WithReadDelay(0*time.Second),
		transport.WithParallelRequests(false))
	httpClient := &http.Client{
		Transport: transport,
	}

	adapter, err := kiotaHttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(tokenProvider, nil, nil, httpClient)
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

func Example_RateLimitHandler() {
	options := make([]abstractions.RequestOption, 0)

	rateLimitHandlerOption := &handlers.RateLimitHandlerOptions{
		ReadDelay:        0 * time.Second,
		WriteDelay:       1 * time.Second,
		ParallelRequests: false,
	}

	retryHandlerOption := &kiotaHttp.RetryHandlerOptions{
		MaxRetries: 3,
	}
	options = append(options, retryHandlerOption)
	options = append(options, rateLimitHandlerOption)

	tokenProvider := auth.NewTokenProvider(
		// to create an authenticated provider, uncomment the below line and pass in your token
		// auth.WithAuthorizationToken("ghp_your_token"),
		auth.WithUserAgent("octokit/go-sdk.example-functions"),
	)

	adapter, err := kiotaHttp.NewNetHttpRequestAdapter(tokenProvider)
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
		Options: options,
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
