
air:
	@air --build.cmd "go build -o bin/api ./cmd/main.go" --build.bin "./bin/api"

run:
	@go run ./cmd/main.go

build:
	@go build -o ./bin/app ./cmd/main.go