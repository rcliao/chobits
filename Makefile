# Project meta
BINARY_NAME=chobits
BINARY_FOLDER=bin
MAIN_FILE=cmd/main.go
BINARY_UNIX=$(BINARY_NAME)_unix

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build
build: 
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v $(MAIN_FILE)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_FOLDER)
run:
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v $(MAIN_FILE)
	./$(BINARY_FOLDER)/$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_UNIX) -v $(MAIN_FILE)
