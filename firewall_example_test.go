package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

func ExampleAPI_ListZoneAccessRules_all() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all access rules for a zone
	response, err := api.ListZoneAccessRules(context.Background(), zoneID, khulnasoft.AccessRule{}, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range response.Result {
		fmt.Printf("%s: %s\n", r.Configuration.Value, r.Mode)
	}
}

func ExampleAPI_ListZoneAccessRules_filterByIP() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch only access rules whose target is 198.51.100.1
	localhost := khulnasoft.AccessRule{
		Configuration: khulnasoft.AccessRuleConfiguration{Target: "198.51.100.1"},
	}
	response, err := api.ListZoneAccessRules(context.Background(), zoneID, localhost, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range response.Result {
		fmt.Printf("%s: %s\n", r.Configuration.Value, r.Mode)
	}
}

func ExampleAPI_ListZoneAccessRules_filterByMode() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch access rules with an action of "block"
	foo := khulnasoft.AccessRule{
		Mode: "block",
	}
	response, err := api.ListZoneAccessRules(context.Background(), zoneID, foo, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range response.Result {
		fmt.Printf("%s: %s\n", r.Configuration.Value, r.Mode)
	}
}

func ExampleAPI_ListZoneAccessRules_filterByNote() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch only access rules with notes containing "example"
	foo := khulnasoft.AccessRule{
		Notes: "example",
	}
	response, err := api.ListZoneAccessRules(context.Background(), zoneID, foo, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range response.Result {
		fmt.Printf("%s: %s\n", r.Configuration.Value, r.Mode)
	}
}
