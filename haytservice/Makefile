PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/serjplus/cosmos-sdk-tutorials/version.Name=HaytService \
	-X github.com/serjplus/cosmos-sdk-tutorials/version.ServerName=nsd \
	-X github.com/serjplus/cosmos-sdk-tutorials/version.ClientName=nscli \
	-X github.com/serjplus/cosmos-sdk-tutorials/version.Version=$(VERSION) \
	-X github.com/serjplus/cosmos-sdk-tutorials/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

include Makefile.ledger
all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/nsd
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/nscli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
