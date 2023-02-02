package main

import (
	"context"
	"fmt"
	"os"

	"github.com/seg491X-team36/vsn-backend/features/postgres"
)

func main() {
	schema, _ := os.Open("./codegen/sql/schema.sql")
	defer schema.Close()

	conn, err := postgres.NewConnection(postgres.DefaultCredentials, schema)
	if err != nil {
		panic(err)
	}

	// example get all users
	users, err := conn.GetAllUsers(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("found", len(users), "users")
}
