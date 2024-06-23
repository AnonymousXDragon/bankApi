migrate_up:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/bankdb?sslmode=disable" -verbose down

migrate_create:
	migrate -ext postgres -dir db/migration -seq init_schema

generate_sqlc:
	sqlc generate

test:
	go run test ./...

.PHONY: migrate_up migrate_down migrate_create generate_sqlc test