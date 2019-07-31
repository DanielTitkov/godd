BUILDPATH=${CURDIR}
PACKAGE="godd"


.PHONY: setup
setup:
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1


.PHONY: test
test: 
	@echo 'Running tests'
	@go test -covermode=atomic -v -race -timeout=30s


.PHONY: lint
lint: 
	@echo 'Running linters'
	@golangci-lint run --enable-all -D gochecknoinits -D gochecknoglobals


.PHONY: build
build: test lint
	@echo 'Building executable'
	@go build -v -o ${BUILDPATH}/bin/${PACKAGE}
	@echo 'Executable at ${BUILDPATH}/bin'


.PHONY: install
install: test lint
	@echo 'Running install'
	@go install


.PHONY: clean
clean:
	@echo 'Cleaning'
	@go clean
	@rm -rf ${BUILDPATH}/bin


.DEFAULT_GOAL := build