schema:
	@rm internal/db/infinity.db
	@sqlite3 internal/db/infinity.db < internal/db/schema.sql

gen:
	@cd internal/db/ && sqlc generate