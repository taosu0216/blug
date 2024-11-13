package main

import (
	"blug/internal/data/ent"
	"context"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	client, err := ent.Open("postgres", "postgres://postgres:pass@localhost:5432/blug?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
