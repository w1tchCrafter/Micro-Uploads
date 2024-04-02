.PHONY: all build run test clean

# set the go executable
GO := go

# set the main package path
MAIN_PACKAGE := ./cmd/micro_uploads

# set the name of the output binary
BINARY := micro_uploads

# Build the project
build:
	@echo "building $(BINARY)..."
	@$(GO) build -o $(BINARY) $(MAIN_PACKAGE)

# Run the project
run: build
	@echo "running $(BINARY)..."
	@./$(BINARY)

# Run tests
test:
	@echo "running tests..."
	@$(GO) test ./...
# Run in development mode
dev:
	@air
# Clean generated files
clean:
	@echo "cleaning up..."
	@rm -f $(BINARY)

# Default target
all: build test run