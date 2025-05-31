migration-create:
	migrate create -ext sql -dir sql/migrations -seq ${NAME}

migration-up:
	migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path sql/migrations up
	sqlc generate

migration-down:
	migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path sql/migrations down
