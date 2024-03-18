ENVIROMENT=staging

dev:
	@echo "Starting development server"
	@echo "Copy environment.yaml to current.yaml"
	cp ./configs/$(ENVIROMENT).yaml ./current.yaml
	go run ./cmd/server/main.go

build-bin:
	go build -o out/$(BINARY_NAME) ./cmd/server/main.go
