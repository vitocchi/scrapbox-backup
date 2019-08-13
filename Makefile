## Sample Makefile for go project
##
## Last update
## 	2019-05-28

# define project name
# make use the project name to name binary file and run it

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
PROJECT_NAME=trade-history-server

.PHONY: vendor

rerun:
	$(MAKE) install
	$(MAKE) run 

run:
	$(PROJECT_NAME)

build: vendor
	GOOS=linux GOARCH=amd64 go build -o binary
	zip binary.zip ./binary

vendor:
	dep ensure

test:
	$(GOTEST) .

clean:
	rm $(GOPATH)/bin/$(PROJECT_NAME)
	rm -rf vendor
