build:
	go build -o singularity .

buildall:
	go build ./...

gen:
	go generate ./replication/internal/proposal110/types.go
	go generate ./replication/internal/proposal120/types.go
	go run handler/datasource/generate/add.go

lint:
	gofmt -s -w .
	golangci-lint run

swag:
	swag init --parseDependency --parseInternal -g singularity.go -d .,./api,./handler -o ./api/docs

test:
	gotestsum --format testname ./...
