package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

func ExampleAPI_ListZones_all() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all zones available to this user.
	zones, err := api.ListZones(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Println(z.Name)
	}
}

func ExampleAPI_ListZones_filter() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch a slice of zones example.org and example.net.
	zones, err := api.ListZones(context.Background(), "example.org", "example.net")
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Println(z.Name)
	}
}
