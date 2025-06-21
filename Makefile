# Install dependencies
tidy:
	@echo -e "Running go mod tidy..."
	@go mod tidy

# Run the application in development mode
run:
	@echo -e "Running the application..."
	@dotenv -e .env -- go run ./cmd/main.go

# Test the application
test:
	@echo -e "Running tests..."
	@dotenv -e .env -- go test -v ./tests/...


.PHONY: tidy run test