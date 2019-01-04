GOCMD       = go
GOBUILD     = $(GOCMD) build
GOCLEAN     = $(GOCMD) clean
GOTEST      = $(GOCMD) test
GOGET       = $(GOCMD) get
GOINSTALL   = $(GOCMD) install
GOLINT      = $(GOPATH)/bin/golint

BINARY_NAME=bin

all: lint test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	OCTOSUMMON_DEBUG=false $(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

install:
	$(GOINSTALL) -v ./...

deps:
	$(GOGET) "github.com/caarlos0/env"
	$(GOGET) "github.com/gorilla/mux"
	$(GOGET) "github.com/Sirupsen/logrus"
	$(GOGET) "github.com/stretchr/testify/assert"

lint:
	$(GOGET) "github.com/golang/lint/golint"
	$(GOLINT) -min_confidence 0 -set_exit_status
