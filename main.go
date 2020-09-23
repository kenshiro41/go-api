package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	mydb "github.com/kenshiro41/go_app/db"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kenshiro41/go_app/gql"
)

const (
	defaultPort = "7890"
)

func main() {
	mode := os.Getenv("MODE")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := mydb.DB

	if mode != "production" {
		err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Panicln(err)
		}
		AccessKeyID := os.Getenv("AccessKeyID")
		SecretAccessKey := os.Getenv("SecretAccessKey")
		fmt.Println(AccessKeyID, SecretAccessKey)
	}

	//Websokcet
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: true,
	}
	options := []handler.Option{
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			fmt.Println(err)
			return fmt.Errorf("Internel server error")
		}),
		handler.WebsocketUpgrader(upgrader),
	}

	r := mux.NewRouter()

	// Graphql PlayGround
	if mode != "production" {
		r.HandleFunc("/", playground.Handler("GraphQL Playground", "/graphql"))
	}

	r.HandleFunc("/graphql", handler.GraphQL(
		gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{DB: db}}),
		options...,
	))

	handler := cors.AllowAll().Handler(r)

	log.Printf("start http://localhost:%s/ ", port)
	http.ListenAndServe(":7890", handler)
}
