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

run-singleton:
	go run ./patterns/creational/singleton/cmd

run-adapter:
	go run ./patterns/structural/adapter/cmd

run-bridge:
	go run ./patterns/structural/bridge/cmd

run-composite:
	go run ./patterns/structural/composite/cmd

run-decorator:
	go run ./patterns/structural/decorator/cmd

run-facade:
	go run ./patterns/structural/facade/cmd

run-flyweight:
	go run ./patterns/structural/flyweight/cmd

run-proxy:
	go run ./patterns/structural/proxy/cmd

run-chainofresponsibility:
	go run ./patterns/behavioral/chainofresponsibility/cmd

run-command:
	go run ./patterns/behavioral/command/cmd
