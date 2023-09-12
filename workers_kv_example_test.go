package khulnasoft_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	khulnasoft "github.com/khulnasoft-lab/go-api"
)

const (
	namespace = "xxxxxx96ee002e8xxxxxx665354c0449"
	accountID = "xxxxxx10ee002e8xxxxxx665354c0410"
)

func ExampleAPI_CreateWorkersKVNamespace() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	req := khulnasoft.CreateWorkersKVNamespaceParams{Title: "test_namespace2"}
	response, err := api.CreateWorkersKVNamespace(context.Background(), khulnasoft.AccountIdentifier(accountID), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func ExampleAPI_ListWorkersKVNamespaces() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	lsr, _, err := api.ListWorkersKVNamespaces(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.ListWorkersKVNamespacesParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lsr)

	resp, _, err := api.ListWorkersKVNamespaces(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.ListWorkersKVNamespacesParams{ResultInfo: khulnasoft.ResultInfo{
		PerPage: 10,
	}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_DeleteWorkersKVNamespace() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	response, err := api.DeleteWorkersKVNamespace(context.Background(), khulnasoft.AccountIdentifier(accountID), namespace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

func ExampleAPI_UpdateWorkersKVNamespace() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := api.UpdateWorkersKVNamespace(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.UpdateWorkersKVNamespaceParams{
		NamespaceID: namespace,
		Title:       "test_title",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_WriteWorkersKVEntry() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	payload := []byte("test payload")
	key := "test_key"

	resp, err := api.WriteWorkersKVEntry(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.WriteWorkersKVEntryParams{
		NamespaceID: namespace,
		Key:         key,
		Value:       payload,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_WriteWorkersKVEntries() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	payload := []*khulnasoft.WorkersKVPair{
		{
			Key:   "key1",
			Value: "value1",
		},
		{
			Key:      "key2",
			Value:    base64.StdEncoding.EncodeToString([]byte("value2")),
			Base64:   true,
			Metadata: "key2's value will be decoded in base64 before it is stored",
		},
	}

	resp, err := api.WriteWorkersKVEntries(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.WriteWorkersKVEntriesParams{
		NamespaceID: namespace,
		KVs:         payload,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_GetWorkersKV() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.GetWorkersKV(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.GetWorkersKVParams{NamespaceID: namespace, Key: key})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resp)
}

func ExampleAPI_DeleteWorkersKVEntry() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	key := "test_key"
	resp, err := api.DeleteWorkersKVEntry(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.DeleteWorkersKVEntryParams{
		NamespaceID: namespace,
		Key:         key,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func ExampleAPI_DeleteWorkersKVEntries() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	keys := []string{"key1", "key2", "key3"}

	resp, err := api.DeleteWorkersKVEntries(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.DeleteWorkersKVEntriesParams{
		NamespaceID: namespace,
		Keys:        keys,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func ExampleAPI_ListWorkersKVKeys() {
	api, err := khulnasoft.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	limit := 50
	prefix := "my-prefix"
	cursor := "AArAbNSOuYcr4HmzGH02-cfDN8Ck9ejOwkn_Ai5rsn7S9NEqVJBenU9-gYRlrsziyjKLx48hNDLvtYzBAmkPsLGdye8ECr5PqFYcIOfUITdhkyTc1x6bV8nmyjz5DO-XaZH4kYY1KfqT8NRBIe5sic6yYt3FUDttGjafy0ivi-Up-TkVdRB0OxCf3O3OB-svG6DXheV5XTdDNrNx1o_CVqy2l2j0F4iKV1qFe_KhdkjC7Y6QjhUZ1MOb3J_uznNYVCoxZ-bVAAsJmXA"

	resp, err := api.ListWorkersKVKeys(context.Background(), khulnasoft.AccountIdentifier(accountID), khulnasoft.ListWorkersKVsParams{
		NamespaceID: namespace,
		Prefix:      prefix,
		Limit:       limit,
		Cursor:      cursor,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
