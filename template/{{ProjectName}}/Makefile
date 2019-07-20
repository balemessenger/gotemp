PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

GOOS := $(shell go env GOOS)
GOOSALT ?= 'linux'
ifeq ($(GOOS),'darwin')
  GOOSALT = 'mac'
endif

PKG := {{ProjectName}}
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: build

checks:
	@echo "Checking dependencies"
	@(env bash $(PWD)/buildscripts/checkdeps.sh)

getdeps:
	@mkdir -p ${GOPATH}/bin
	@which golint 1>/dev/null || (echo "Installing golint" && go get -u golang.org/x/lint/golint)
	@which staticcheck 1>/dev/null || (echo "Installing staticcheck" && wget --quiet -O ${GOPATH}/bin/staticcheck https://github.com/dominikh/go-tools/releases/download/2019.1/staticcheck_${GOOS}_amd64 && chmod +x ${GOPATH}/bin/staticcheck)
	@which misspell 1>/dev/null || (echo "Installing misspell" && wget --quiet https://github.com/client9/misspell/releases/download/v0.3.4/misspell_0.3.4_${GOOSALT}_64bit.tar.gz && tar xf misspell_0.3.4_${GOOSALT}_64bit.tar.gz && mv misspell ${GOPATH}/bin/misspell && chmod +x ${GOPATH}/bin/misspell && rm -f misspell_0.3.4_${GOOSALT}_64bit.tar.gz)

verifiers: getdeps vet fmt lint staticcheck spelling

vet:
	@echo "Running $@"
	@GO111MODULE=on go vet ./...

fmt:
	@echo "Running $@"
	@GO111MODULE=on gofmt -w -d api/
	@GO111MODULE=on gofmt -w -d cmd/
	@GO111MODULE=on gofmt -w -d internal/
	@GO111MODULE=on gofmt -w -d pkg/

lint:
	@echo "Running $@"
	@GO111MODULE=on ${GOPATH}/bin/golint -set_exit_status cmd/...
	@GO111MODULE=on ${GOPATH}/bin/golint -set_exit_status internal/...
	@GO111MODULE=on ${GOPATH}/bin/golint -set_exit_status pkg/...
	@GO111MODULE=on ${GOPATH}/bin/golint -set_exit_status test/...

staticcheck:
	@echo "Running $@"
	@GO111MODULE=on ${GOPATH}/bin/staticcheck api/...
	@GO111MODULE=on ${GOPATH}/bin/staticcheck cmd/...
	@GO111MODULE=on ${GOPATH}/bin/staticcheck internal/...
	@GO111MODULE=on ${GOPATH}/bin/staticcheck pkg/...
	@GO111MODULE=on ${GOPATH}/bin/staticcheck test/...

spelling:
	@GO111MODULE=on ${GOPATH}/bin/misspell -locale US -error `find cmd/`
	@GO111MODULE=on ${GOPATH}/bin/misspell -locale US -error `find internal/`
	@GO111MODULE=on ${GOPATH}/bin/misspell -locale US -error `find pkg/`
	@GO111MODULE=on ${GOPATH}/bin/misspell -locale US -error `find test/`
	@GO111MODULE=on ${GOPATH}/bin/misspell -locale US -error `find docs/`

init:
	@echo "Running $@"
	@GO111MODULE=on go mod init
	@GO111MODULE=on go mod vendor

test:
	@echo "Running unit tests"
	# @go test -count=1 pkg/
	# @go test -count=1 internal/
	# @go test -count=1 api/
	@go test -count=1 test/

coverage: build
	@echo "Running all coverage"
	@GO111MODULE=on CGO_ENABLED=0 go test -v -coverprofile=coverage.txt -covermode=atomic ./...

run:
	@echo "Running $@"
	@GO111MODULE=on go run main.go

docker: test fmt
	@echo "Running $@"
	@(env bash $(PWD)/scripts/docker.sh)

deploy: docker
	@echo "Running $@"
	@(env bash $(PWD)/scripts/awx.sh)

proto:
	@echo "Running $@"
	@(env bash $(PWD)/scripts/genproto.sh)

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.pb.go' | xargs rm -fv
	@rm -rvf build/debug
	@rm -rvf build/release

.PHONY: all test clean