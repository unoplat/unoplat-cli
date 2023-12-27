# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOINSTALL = $(GOCMD) install
GOCLEAN = $(GOCMD) clean

# Build and install directory
BUILD_DIR = bin
INSTALL_DIR = $(GOPATH)/bin

# Main package file
MAIN_FILE = main.go

# Binary name
BINARY_NAME = unoplat-cli

# Test package
TEST_PACKAGE = ./...

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)

test:
	@echo "Running tests..."
	$(GOTEST) $(TEST_PACKAGE) -v

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

.PHONY: all build install test clean
