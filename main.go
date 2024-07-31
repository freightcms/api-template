package main

import (
	"net/http"

	"github.com/freightcms/webservice-template/schemas"
	"github.com/freightcms/webservice-template/web"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"entityList": &graphql.Field{
				Description: "Get All entities",
				Type:        web.PaginatedEntitiesObject,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return &struct {
						ETag        string                 `json:"etag"`
						Entities    []schemas.EntitySchema `json:"entities"`
						Count       int                    `json:"count"`
						PageSize    int                    `json:"pageSize"`
						Page        int                    `json:"page"`
						NextUrl     string                 `json:"nextURL"`
						PreviousUrl string                 `json:"previousURL"`
					}{
						ETag:        "",
						Entities:    []schemas.EntitySchema{},
						Count:       0,
						PageSize:    10,
						Page:        0,
						NextUrl:     "",
						PreviousUrl: "",
					}, nil
				},
			},
		},
	})
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
