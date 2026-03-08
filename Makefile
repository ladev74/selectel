BINARY_NAME=custom-gcl
GOBIN=$(shell go env GOPATH)/bin
CONFIG_FILE=.golangci.yaml
HOME_DIR=$(HOME)

all: build install copy_config

build:
	golangci-lint custom -v

install:
	cp $(BINARY_NAME) $(GOBIN)/$(BINARY_NAME)

copy_config:
	cp $(CONFIG_FILE) $(HOME_DIR)/$(CONFIG_FILE)
