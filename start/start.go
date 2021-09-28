package main

import (
	"context"
	"log"
	"myeduate/ent"
	"myeduate/ent/migrate"
	"myeduate/resolvers"
	"myeduate/utils"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Printf("cannot load config: %v", err)
	}

	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.Postgres, config.DbSource)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// create a new serve mux and register the handlers
	route := mux.NewRouter()

	// CORS
	ch := cors()

	// create a new server
	s := http.Server{
		Addr:    config.ServerAddr, // configure the bind address
		Handler: ch(route),         // set the default handler
		//		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	route.Handle("/", playground.Handler("myeduate", "/query"))
	route.Handle("/query", srv)
	log.Printf("listening on : %v", config.ServerAddr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("http server terminated", err)
	}
}

func cors() mux.MiddlewareFunc {
	return gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept"}),
		gohandlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions}),
		gohandlers.AllowCredentials(),
	)
}
