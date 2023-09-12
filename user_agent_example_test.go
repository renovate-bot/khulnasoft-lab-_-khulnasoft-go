package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

func ExampleAPI_ListUserAgentRules_all() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all Zone Lockdown rules for a zone, by page.
	rules, err := api.ListUserAgentRules(context.Background(), zoneID, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rules.Result {
		fmt.Printf("%s: %s\n", r.Configuration.Target, r.Configuration.Value)
	}
}
