package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-json"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

func ExampleAPI_AccessAuditLogs() {
	api, err := khulnasoft.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	filterOpts := khulnasoft.AccessAuditLogFilterOptions{}
	results, _ := api.AccessAuditLogs(context.Background(), "someaccountid", filterOpts)

	for _, record := range results {
		b, _ := json.Marshal(record)
		fmt.Println(string(b))
	}
}
