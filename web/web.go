package web

import "github.com/graphql-go/graphql"

var (
	entitiesQuery graphql.ObjectConfig
)

func init() {
	entitiesQuery = graphql.NewObjectConfig(
	graphql.NewObject({
	})
}
