package postgres

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/lib/pq"
	"github.com/seg491X-team36/vsn-backend/codegen/db"
)

type Connection struct {
	*db.Queries
}

func NewConnection(credentials Credentials, schemaReader io.Reader) (*Connection, error) {
	// connect to postgres
	conn, err := sql.Open("postgres", credentials.String())
	if err != nil {
		return nil, err
	}

	// read the schema
	schema, _ := io.ReadAll(schemaReader)

	// update the schema
	_, err = conn.Exec(string(schema))
	if err == nil {
		fmt.Println("successfully updated postgres schema")
	}

	return &Connection{
		Queries: db.New(conn),
	}, nil
}
