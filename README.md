# Barcode API

Get product detail using barcode

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

## Migration

```sh
# Install migration cli
go install github.com/pressly/goose/v3/cmd/goose@latest

# Make sure you have the correct env
DB_DRIVER=sqlite3
DB_URL=./.tmp/sqlite.db
DB_MIGRATION_DIR=./database/migration

GOOSE_DRIVER=$DB_DRIVER
GOOSE_DBSTRING=$DB_URL
GOOSE_MIGRATION_DIR=$DB_MIGRATION_DIR

# Run migration
goose up
```
