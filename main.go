package main

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"os"
)

func main() {
	client := typesense.NewClient(
		typesense.WithServer("http://typesense:8108"),
		typesense.WithAPIKey("test_demo"))

	test := "num_employees"
	schema := &api.CollectionSchema{
		Name: "companies",
		Fields: []api.Field{
			{
				Name: "company_name",
				Type: "string",
			},
			{
				Name: "num_employees",
				Type: "int32",
			},
			{
				Name: "country",
				Type: "string",
			},
		},
		DefaultSortingField: &test,
	}

	data, _ := client.Collections().Create(schema)
	fmt.Println(data)

	document := struct {
		ID           string `json:"id"`
		CompanyName  string `json:"company_name"`
		NumEmployees int    `json:"num_employees"`
		Country      string `json:"country"`
	}{
		ID:           "1234",
		CompanyName:  "Stark Industries 111",
		NumEmployees: 5215,
		Country:      "USA",
	}

	result, _ := client.Collection("companies").Documents().Create(document)
	fmt.Println(result)

	searchParameters := &api.SearchCollectionParams{
		Q:       "stark",
		QueryBy: "company_name",
	}

	// search result
	result1, _ := client.Collection("companies").Documents().Search(searchParameters)
	for _, value := range *result1.Hits {
		document := value.Document
		fmt.Println((*document)["company_name"], (*document)["num_employees"])
	}

	// get document by id
	result2, _ := client.Collection("companies").Document("123").Retrieve()
	fmt.Println(result2)

	// get all collection
	fmt.Println(client.Collection("companies").Retrieve())

	// export collection
	client.Collection("companies").Documents().Export()

	// import
	documents := []interface{}{
		struct {
			ID           string `json:"id"`
			CompanyName  string `json:"company_name"`
			NumEmployees int    `json:"num_employees"`
			Country      string `json:"country"`
		}{
			ID:           "1236",
			CompanyName:  "Stark Industries",
			NumEmployees: 5215,
			Country:      "USA",
		},
	}

	action := "create"
	batchSize := 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batchSize,
	}
	client.Collection("companies").Documents().Import(documents, params)

	fmt.Println("import file json")
	action = "create"
	batchSize = 40
	params = &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batchSize,
	}
	importBody, err := os.Open("documents.jsonl")
	fmt.Println(importBody, err)
	client.Collection("companies").Documents().ImportJsonl(importBody, params)

	// list Collection
	fmt.Println(client.Collections().Retrieve())
}
