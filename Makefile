# Makefile for go-search-tool (2025)

# Name of the final executable
BINARY_NAME=findit

build:
	@echo "Building the search utility..."
	@go build -o $(BINARY_NAME) ./cmd/findit/main.go
	@echo "Build complete! Use ./$(BINARY_NAME) to run."

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Cleaned!"

# Installs the tool to your system path
install:
	@go install ./cmd/findit
