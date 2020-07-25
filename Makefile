# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build --trimpath
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINPATH=bin/
BINARYNAME=grgd

all: test build
build:
	GO111MODULE=on $(GOBUILD) -o $(BINPATH)$(BINARYNAME)
test:
	$(GOTEST) -v ./...
run:
	./$(BINPATH)$(BINARYNAME)

linux: docker-build docker-run
docker-run:
	docker run -it -v $(PWD)/bin/grgd-linux:/usr/local/bin/grgd -v $(HOME)/.grgd/:/root/.grgd/ -v $(HOME)/.grgd/plugins/binaries-linux/:/root/.grgd/plugins/binaries/ golang sh
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -v $(PWD):/src/ -w /src/ golang:latest $(GOBUILD) -o ./bin/grgd-linux
