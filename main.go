package main

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
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
		ID:           "123",
		CompanyName:  "Stark Industries",
		NumEmployees: 5215,
		Country:      "USA",
	}

	result, _ := client.Collection("companies").Documents().Create(document)
	fmt.Println(result)

	searchParameters := &api.SearchCollectionParams{
		Q: "stark",
	}

	result1, _ := client.Collection("companies").Documents().Search(searchParameters)
	fmt.Println(result1)
}
