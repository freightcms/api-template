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

func ServeHTTP(h *handler.Handler) http.HandlerFunc {
	// this is a hack around the internal go-graphql ServeHTTP because the API
	// does not appropriately expose the request context or a good way to add
	// middleware into requests.
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := client.StartSession(&options.SessionOptions{})
		if err != nil {
			panic(err)
		}
		ctx := context.WithValue(r.Context(), "dbSession", session)
		h.ContextHandler(ctx, w, r)
	}
}

func main() {
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	uri := os.Getenv("MONGODB_URI")
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
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

	middlewareHandler := struct {
		ServeHTTP http.HandlerFunc
	}{
		ServeHTTP: ServeHTTP(h),
	}
	http.Handle("/graphql", middlewareHandler)
	http.ListenAndServe(":8080", nil)
}
