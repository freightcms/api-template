package main

import (
	"context"
	"net/http"
	"os"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/freightcms/webservice-template/web"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	return mongo.Connect(context.Background(), opts)
}

func main() {
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"entityList": &graphql.Field{
				Description: "Get All entities",
				Type:        web.PaginatedEntitiesObject,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
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
	srvr := &http.ServeMux{}
	srvr.Handle("/graphql", h)

	http.ListenAndServe(":3000", srvr)
}
