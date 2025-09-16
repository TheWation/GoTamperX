# TamperX Makefile

.PHONY: build clean test install run help

# Default target
all: build

# Build the application
build:
	@echo "Building TamperX..."
	@go build -o tamperx ./cmd/tamperx
	@echo "Build complete!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f tamperx
	@go clean
	@echo "Clean complete!"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Install globally
install:
	@echo "Installing TamperX..."
	@go install ./cmd/tamperx
	@echo "Installation complete!"

# Run with example
run: build
	@echo "Running TamperX with example..."
	@./tamperx -u https://httpbin.org/get -c 4

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  install  - Install globally"
	@echo "  run      - Build and run with example"
	@echo "  help     - Show this help"
