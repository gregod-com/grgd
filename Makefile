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

tdd:
	fswatch -o ../* | xargs -n1 -I{} bash -c 'clear && go test ./...'

mocks:
	@for f in interfaces/*.go; do \
		echo generate $${f}; \
		mockgen -imports interfaces/IConfig.go=github.com/gregod-com/grgd/interfaces,interfaces/ICore.go=github.com/gregod-com/grgd/interfaces,interfaces/IExtractor.go=github.com/gregod-com/grgd/interfaces,interfaces/IHelper.go=github.com/gregod-com/grgd/interfaces,interfaces/INetworker.go=github.com/gregod-com/grgd/interfaces,interfaces/IProfile.go=github.com/gregod-com/grgd/interfaces,interfaces/IProject.go=github.com/gregod-com/grgd/interfaces  \
		--source=$${f} -destination interfaces/mocks/mock`basename $${f}` -package mocks; \
	done

proto:
	protoc  --go_out=../  protobuf/*.proto

