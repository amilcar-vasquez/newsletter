.PHONY: run-app
run-app:
	@echo "Running the app..."
	@go run ./cmd/web -addr=":3000"