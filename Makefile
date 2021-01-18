# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -ldflags="-w"
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINPATH=bin/
BINARYNAME=grgd-darwin

all: test build-native stats

test:
	cd interfaces && ./makeMocks.sh
	$(GOTEST) ./...

build-native:
	GO111MODULE=on $(GOBUILD) -o $(BINPATH)$(BINARYNAME)

run:
	./$(BINPATH)$(BINARYNAME)

stats:
	du -sh $(BINPATH)$(BINARYNAME)

upload:
	mc cp $(BINPATH)$(BINARYNAME) minio/public/grgd/$(BINARYNAME)

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
