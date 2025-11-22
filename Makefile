# Variables
BINARY_NAME=flow
GO_FILES=./cmd/flow

# Default command (runs when you just type 'make')
all: audit build

# 1. Audit: Runs formatting, vetting, and testing (The "Guard Dog")
audit:
	@echo "🔍 Checking formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then echo "❌ Format error. Run 'make fmt'"; exit 1; fi
	@echo "🔍 Vetting code..."
	@go vet ./...
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Audit passed!"

# 2. Build: Compiles the binary
build:
	@echo "🏗️  Building..."
	@go build -o bin/$(BINARY_NAME) $(GO_FILES)
	@echo "✅ Build complete: ./bin/$(BINARY_NAME)"

# 3. Run: Builds and runs the app
run: build
	@./bin/$(BINARY_NAME)

# 4. Format: Fixes indentation automatically
fmt:
	@go fmt ./...
	@echo "✨ Code formatted"

# 5. Clean: Removes binaries
clean:
	@rm -rf bin
