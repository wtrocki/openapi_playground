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
	api_client := sw.NewAPIClient(configuration)
	resp, r, err := api_client.TestApi.GetRoot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestApi.GetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoot`: DataObj
	fmt.Fprintf(os.Stdout, "Response from `TestApi.GetRoot`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Headers: %v\n", r.Header)
}
