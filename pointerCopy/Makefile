# Basic go commands
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

BIN_NAME = $(shell basename $(CURDIR))

build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BIN_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f bin/$(BIN_NAME)
	
