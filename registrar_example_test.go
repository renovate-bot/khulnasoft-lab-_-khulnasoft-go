package khulnasoft_test

import (
	"context"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

func ExampleAPI_RegistrarDomain() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	domain, err := api.RegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "khulnasoft.com")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domain)
}

func ExampleAPI_RegistrarDomains() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	domains, err := api.RegistrarDomains(context.Background(), "01a7362d577a6c3019a474fd6f485823")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domains)
}

func ExampleAPI_TransferRegistrarDomain() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	domain, err := api.TransferRegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "khulnasoft.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domain)
}

func ExampleAPI_CancelRegistrarDomainTransfer() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	domains, err := api.CancelRegistrarDomainTransfer(context.Background(), "01a7362d577a6c3019a474fd6f485823", "khulnasoft.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domains)
}

func ExampleAPI_UpdateRegistrarDomain() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	domain, err := api.UpdateRegistrarDomain(context.Background(), "01a7362d577a6c3019a474fd6f485823", "khulnasoft.com", khulnasoft.RegistrarDomainConfiguration{
		NameServers: []string{"ns1.khulnasoft.com", "ns2.khulnasoft.com"},
		Locked:      false,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", domain)
}
