.PHONY: clean build run deploy unit_test

UNAME := $(shell uname -s)
curDIR := $(shell basename `pwd`)
app := $(curDIR).$(UNAME)
SOURCES := $(shell find * -type f -name "*.go")

default: run

$(app): $(SOURCES) go.mod go.sum
	go build -v -o $(app)

build: $(app)

run: build
	./$(app)

deploy: build
	./dtools.sh

unit_test:
	go test -v ./... -test.coverprofile cover.out
	go tool cover -func=cover.out

cloc:
	cloc .

clean:
	rm -rf $(app)