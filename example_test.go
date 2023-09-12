package khulnasoft_test

import (
	"context"
	"fmt"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

const (
	user   = "khulnasoft@example.org"
	domain = "example.com"
	apiKey = "deadbeef"
)

func Example() {
	api, err := khulnasoft.New("deadbeef", "khulnasoft@example.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch the zone ID for zone example.org
	zoneID, err := api.ZoneIDByName("example.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch all DNS records for example.org
	records, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range records {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}
