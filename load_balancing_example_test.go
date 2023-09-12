package khulnasoft_test

import (
	context "context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

func ExampleAPI_ListLoadBalancers() {
	// Construct a new API object.
	api, err := khulnasoft.New("deadbeef", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// List LBs configured in zone.
	lbList, err := api.ListLoadBalancers(context.Background(), khulnasoft.ZoneIdentifier("d56084adb405e0b7e32c52321bf07be6"), khulnasoft.ListLoadBalancerParams{})
	if err != nil {
		log.Fatal(err)
	}

	for _, lb := range lbList {
		fmt.Println(lb)
	}
}

func ExampleAPI_GetLoadBalancerPoolHealth() {
	// Construct a new API object.
	api, err := khulnasoft.New("deadbeef", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch pool health details.
	healthInfo, err := api.GetLoadBalancerPoolHealth(context.Background(), khulnasoft.AccountIdentifier("01a7362d577a6c3019a474fd6f485823"), "example-pool-id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(healthInfo)
}
