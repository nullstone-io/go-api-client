NAME := nullstone

.PHONY: setup test

.DEFAULT_GOAL: default

default: setup

setup:
	cd ~ && go get gotest.tools/gotestsum && cd -

test:
	go fmt ./...
	gotestsum
