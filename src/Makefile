# Build
build:
	GODS=linux GOARCH=amd64 go build -v -o bin/person-service ./cmd/main.go

# Test
test:
	go test -v -count=1 -tags=testing ./...

# Migrations
## Create
migrate-create: # migrate create NAME=<name>
	migrate create -ext sql -dir ./migrations -seq $(NAME)

## Up
migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@postgres:5432/lab01?sslmode=disable' up

## Down
migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@postgres:5432/lab01?sslmode=disable' down
