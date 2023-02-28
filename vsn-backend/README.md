# Backend

```
# Run GraphQL code generator
go run github.com/99designs/gqlgen generate --config=./codegen/gqlgen.yaml

# Run SQL code generator
docker run --rm -v "YOUR-PATH/vsn-backend:/src" -w /src/codegen kjconroy/sqlc generate

# Run Postgres
docker run --name vsn-db -p 5432:5432 -e POSTGRES_DB=vsn-db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_PORT=5432 -d postgres
```

