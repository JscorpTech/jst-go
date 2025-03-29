
air:
	@air --build.cmd "go build -o bin/api ./cmd/api/main.go" --build.bin "./bin/api"

run:
	@go run ./cmd/api/main.go

build:
	@go build -o ./bin/app ./cmd/api/main.go

migrate:
	@go run ./cmd/api/migrate.go