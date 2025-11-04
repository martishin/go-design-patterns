tidy:
	go mod tidy

test:
	go test ./...

lint:
	golangci-lint run ./...

run-factorymethod:
	go run ./patterns/creational/factorymethod/cmd
