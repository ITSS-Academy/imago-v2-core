ENVIROMENT=staging

dev:
	@echo "Starting development server"
	@echo "Copy environment.yaml to current.yaml"
	go run ./cmd/server/main.go --env-file dev.env up

build-bin:
	go build -o out/$(BINARY_NAME) ./cmd/server/main.go
