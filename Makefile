CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/sfomuseum/go-sfomuseum-airlines
	cp *.go src/github.com/sfomuseum/go-sfomuseum-airlines/
	cp -r flysfo src/github.com/sfomuseum/go-sfomuseum-airlines/
	cp -r sfomuseum src/github.com/sfomuseum/go-sfomuseum-airlines/
	cp -r wikipedia src/github.com/sfomuseum/go-sfomuseum-airlines/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-csv"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt flysfo/*.go
	go fmt sfomuseum/*.go
	go fmt wikipedia/*.go

bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/lookup cmd/lookup.go

test-flysfo:
	./bin/lookup -source flysfo B6 DI EI BF HA IG JL LH MH NZ OZ QF SK SN SQ AA AV DL NH AM HX KE FJ PR AY LX CA SU AZ UX CZ AF KL RJ KE CX TG
