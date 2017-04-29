.PHONY: all docker-push

SHELL := /bin/bash

all:
	mkdir -p .gopath/src/github.com/euank/
	[[ -L ./.gopath/src/github.com/euank/alexa-parrot ]] || ln -vsf ../../../.. ./.gopath/src/github.com/euank/alexa-parrot || exit 255
	GOPATH="$(shell pwd)/.gopath" && cd ./.gopath/src/github.com/euank/alexa-parrot && CGO_ENABLED=0 go build -o alexa-parrot .
