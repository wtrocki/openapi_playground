package main

import (
	"context"
	"fmt"
	"os"

	sw "github.com/wtrocki/openapi_playground/contenttype"
)

func main() {
	configuration := sw.NewConfiguration()
	configuration.Debug = true
	configuration.Host = "127.0.0.1:4010"
	configuration.Scheme = "http"
	api_client := sw.NewAPIClient(configuration)
	request := api_client.TestApi.GetRoot(context.Background())
	resp, r, err := request.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestApi.GetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoot`: DataObj
	fmt.Fprintf(os.Stdout, "Response from `TestApi.GetRoot`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Headers: %v\n", r.Header)
}
