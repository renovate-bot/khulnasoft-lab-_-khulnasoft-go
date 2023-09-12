package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

var exampleNewRateLimit = khulnasoft.RateLimit{
	Description: "test",
	Match: khulnasoft.RateLimitTrafficMatcher{
		Request: khulnasoft.RateLimitRequestMatcher{
			URLPattern: "exampledomain.com/test-rate-limit",
		},
	},
	Threshold: 0,
	Period:    0,
	Action: khulnasoft.RateLimitAction{
		Mode:    "ban",
		Timeout: 60,
	},
	Correlate: &khulnasoft.RateLimitCorrelate{
		By: "nat",
	},
}

func ExampleAPI_CreateRateLimit() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimit, err := api.CreateRateLimit(context.Background(), zoneID, exampleNewRateLimit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimit)
}

func ExampleAPI_ListRateLimits() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	pageOpts := khulnasoft.PaginationOptions{
		PerPage: 5,
		Page:    1,
	}
	rateLimits, _, err := api.ListRateLimits(context.Background(), zoneID, pageOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
	for _, r := range rateLimits {
		fmt.Printf("%+v\n", r)
	}
}

func ExampleAPI_RateLimit() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimits, err := api.RateLimit(context.Background(), zoneID, "my_rate_limit_id")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
}

func ExampleAPI_DeleteRateLimit() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeleteRateLimit(context.Background(), zoneID, "my_rate_limit_id")
	if err != nil {
		log.Fatal(err)
	}
}
