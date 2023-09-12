package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/khulnasoft-go"
)

var exampleNewPageRule = khulnasoft.PageRule{
	Actions: []khulnasoft.PageRuleAction{
		{
			ID:    "always_online",
			Value: "on",
		},
		{
			ID:    "ssl",
			Value: "flexible",
		},
	},
	Targets: []khulnasoft.PageRuleTarget{
		{
			Target: "url",
			Constraint: struct {
				Operator string "json:\"operator\""
				Value    string "json:\"value\""
			}{Operator: "matches", Value: fmt.Sprintf("example.%s", domain)},
		},
	},
	Priority: 1,
	Status:   "active",
}

func ExampleAPI_CreatePageRule() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	pageRule, err := api.CreatePageRule(context.Background(), zoneID, exampleNewPageRule)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", pageRule)
}

func ExampleAPI_ListPageRules() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	pageRules, err := api.ListPageRules(context.Background(), zoneID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", pageRules)
	for _, r := range pageRules {
		fmt.Printf("%+v\n", r)
	}
}

func ExampleAPI_PageRule() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	pageRules, err := api.PageRule(context.Background(), zoneID, "my_page_rule_id")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", pageRules)
}

func ExampleAPI_DeletePageRule() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeletePageRule(context.Background(), zoneID, "my_page_rule_id")
	if err != nil {
		log.Fatal(err)
	}
}
