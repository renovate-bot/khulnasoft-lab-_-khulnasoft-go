package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

func ExampleAPI_ListDNSRecords_all() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all records for a zone
	recs, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}

func ExampleAPI_ListDNSRecords_filterByContent() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	recs, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{Content: "198.51.100.1"})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}

func ExampleAPI_ListDNSRecords_filterByName() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	recs, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{Name: "foo.example.com"})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}

func ExampleAPI_ListDNSRecords_filterByType() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	recs, _, err := api.ListDNSRecords(context.Background(), khulnasoft.ZoneIdentifier(zoneID), khulnasoft.ListDNSRecordsParams{Type: "AAAA"})
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}
