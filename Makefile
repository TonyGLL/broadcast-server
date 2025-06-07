# Variables
BINARY_NAME=broadcast-server

# Comando para construir el proyecto
build:
	@echo "Building the project..."
	@go build -o $(BINARY_NAME) cmd/main.go
	@echo "Build complete."

# Comando para ejecutar el proyecto
start: build
	@./$(BINARY_NAME) start

connect: build
	@./$(BINARY_NAME) connect