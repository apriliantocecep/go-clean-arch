MIGRATION_DIR=database/migrations
MIGRATION_EXT=sql
MIGRATION_COMMAND=migrate
DATABASE_URL=mysql://root:admin@tcp(localhost:3306)/restful

create_table:
	@echo "Creating table $(table)"
	@$(MIGRATION_COMMAND) create -ext $(MIGRATION_EXT) -dir $(MIGRATION_DIR) create_$(table)_table

migrate_up:
	@echo "Running database migrations"
	@$(MIGRATION_COMMAND) -database "$(DATABASE_URL)" -path $(MIGRATION_DIR) up

run:
	go run cmd/web/main.go

wire:
	wire ./internal/server/

build:
	go build -C cmd/web/ -o ../../bin/web

test:
	go test ./... -v

.PHONY: create_table migrate_up wire run build