default: all
TARGETLIST=$(shell ls cmd | grep -v tools)
PKG_LIST := $(shell go list ./... | grep -v infrastructure | grep -v test | grep -v cmd | grep -v app/controller )
all: ${TARGETLIST}

init:
	@echo GOROOT=$(GOROOT)
	@echo GOPATH=$(GOPATH)

${TARGETLIST}: %:
	go build -mod=vendor  -o bin/$@ cmd/$@/main.go;
module_msg:
	go build -mod=vendor  -o bin/msg cmd/msg/main.go;

test: init ## Run unittests
	go test -v ${PKG_LIST}
coverage: init ## Generate global code coverage report
	./test/coverage.sh ${PKG_LIST};
