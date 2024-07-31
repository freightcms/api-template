package web

import "github.com/graphql-go/graphql"

var (
	// EntityObject is the GraphQL schema used for the Generic Entity of the
	// API to fetch. This can be replaced with your real Entity Implemntation.
	EntityObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity",
		Fields: graphql.Fields{
			"prop1": &graphql.Field{
				Type: graphql.String,
				Name: "Prop1",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "hello", nil
				},
			},
			"prop2": &graphql.Field{
				Type: graphql.Int,
				Name: "Prop2",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
		},
	})
	// PaginatedEntitiesObject reveals the standardized response for a paginated list of entities
	// fetched from another web service, or data system. This can be re-used across one or many
	// different types of Paginated Entity types.
	PaginatedEntitiesObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Entities",
		Fields: graphql.Fields{
			"etag": &graphql.Field{
				Type: graphql.String,
				Name: "Etag",
			},
			"entities": &graphql.Field{
				Type: graphql.NewList(EntityObject),
				Name: "Entities",
			},
			"page": &graphql.Field{
				Type: graphql.Int,
				Name: "Page",
			},
			"count": &graphql.Field{
				Type: graphql.Int,
				Name: "Count",
			},
			"pageSize": &graphql.Field{
				Type: graphql.Int,
				Name: "pageSize",
			},
			"nextUrl": &graphql.Field{
				Type: graphql.String,
				Name: "NextURL",
			},
			"previousUrl": &graphql.Field{
				Type: graphql.String,
				Name: "PreviousURL",
			},
		},
	})
)
