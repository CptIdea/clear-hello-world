install-dependencies:
	go install github.com/golang/mock/mockgen@v1.6.0

generate:
	mockgen -source=./pkg/domain/hello_world.go -destination=./pkg/mock/hello_world_mock.go

test:
	go test -v ./pkg/domain

test-with-coverage:
	go test -v -cover ./pkg/domain

run:
	go run ./cmd/print

build:
	go build -o ./bin/print ./cmd/print
