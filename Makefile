schema:
	@psql $${DB_URL} -f internal/db/schema.sql

gen:
	@cd internal/db/ && sqlc generate

templ:
	@templ generate

build:
	@go build -o bin/infinity-play