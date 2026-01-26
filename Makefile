tidy:
	go mod tidy

test:
	go test ./...

lint:
	golangci-lint run ./...

run-factorymethod:
	go run ./patterns/creational/factorymethod/cmd

run-abstractfactory:
	go run ./patterns/creational/abstractfactory/cmd

run-builder:
	go run ./patterns/creational/builder/cmd

run-prototype:
	go run ./patterns/creational/prototype/cmd
