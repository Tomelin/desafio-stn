build:
	go build -o server main.go

swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/main.go --output docs/swagger

tidy:
	go mod tidy

run-swagger:
	make swag
	go run cmd/*

run:
	go run cmd/*

test:
	go test ./... -cover -v