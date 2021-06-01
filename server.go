package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clshu/gqlgen-todos/dbmgm"
	"github.com/clshu/gqlgen-todos/graph"
	"github.com/clshu/gqlgen-todos/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"
const defaulGraphQLPath = "/tvu_graphql"

func main() {
	setUpEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	gqlPath := os.Getenv("GRAPHQL_PATH")
	if gqlPath == "" {
		gqlPath = defaulGraphQLPath
	}

	err := dbmgm.Connect()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/tvu_playground", playground.Handler("GraphQL playground", gqlPath))
	http.Handle(gqlPath, srv)

	log.Printf("connect to http://localhost:%s/tvu_playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Set up development or test environment variables
// Expect files to be in project_root/config
// Set GOGO_ENV first before go run server.go
func setUpEnv() {
	const dir string = "config/"
	var fname string
	gogo := os.Getenv("GOGO_ENV")

	switch gogo {
	case "dev":
		fname = dir + "dev.env"
		break
	case "test":
		fname = dir + "test.env"
		break
	default:
		// production environment
		// Do nothing. Let clound platform environment take over
		return
	}

	envMap, err := godotenv.Read(fname)
	if err != nil {
		fmt.Printf("Reading file %v failed. %v", fname, err.Error())
		return
	}

	for key, value := range envMap {
		os.Setenv(key, value)
		// fmt.Printf("%v=%v\n", key, os.Getenv(key))
	}

}
