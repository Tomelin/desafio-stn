build:
	go build -o server main.go

# swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/main.go --output docs/swagger
swag:
	swag init --parseDependency  --parseInternal -g cmd/main.go --output docs/swagger

tidy:
	go mod tidy

run-swagger:
	make swag
	go run cmd/*

run:
	make fmt
	go run cmd/*

test:
	go test ./... -cover -v

fmt:
	go fmt ./... 
