migration-create:
	migrate create -ext sql -dir migrations -seq ${NAME}

migration-up:
	migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path migrations up

