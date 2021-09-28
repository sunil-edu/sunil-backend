package myeduate

import (
	"context"
	"log"
	"myeduate/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func Example_MstCustomer() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.Postgres, "postgresql://root:root123@localhost:5432/eduate?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:
}
