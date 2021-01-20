# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -ldflags="-w"
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINPATH=bin/
BINARYNAME=grgd
OS=darwin
PLATFORM=amd64

all: test build-native stats index

test:
	cd interfaces && ./makeMocks.sh
	$(GOTEST) ./...

build-native:
	GO111MODULE=on $(GOBUILD) -o $(BINPATH)$(BINARYNAME)-$(OS)-$(PLATFORM)

run:
	./$(BINPATH)$(BINARYNAME)

stats:
	du -sh $(BINPATH)$(BINARYNAME)-$(OS)-$(PLATFORM)

index:
	./idxer $(BINARYNAME) $(OS) $(PLATFORM)

upload:
	mc mirror --remove --overwrite bin/ minio/public/grgd/

cover:
	$(GOTEST) -coverprofile=coverage.out -cover ./...
	go tool cover -html=coverage.out

docker: docker-build-bin docker-build 

docker-build-bin:
	docker run --rm -it 					\
		-v "$(GOPATH)":/go 					\
		-v $(PWD):/src/ 					\
		-w /src/ 							\
		-e CGO_ENABLED=1 					\
		-e GO111MODULE=on 					\
		-e GOOS=linux 						\
		-e GOARCH=amd64 					\
		golang:latest 						\
		$(GOBUILD) -ldflags="-w -s"			\
		-o ./$(BINPATH)$(BINARYNAME)-linux

docker-build:
	docker build -t registry.gitlab.com/iamdevelopment/iamk3d:latest .
	docker push registry.gitlab.com/iamdevelopment/iamk3d:latest
