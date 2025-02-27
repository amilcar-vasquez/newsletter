.PHONY: run-app
run-app:
	@echo "Running the app..."
	@go run ./cmd/web -addr=":3000"

.PHONY: run-test
run-test:
	@echo "Running the tests..."
	@go test -v ./...