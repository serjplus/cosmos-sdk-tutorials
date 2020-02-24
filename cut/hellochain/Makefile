VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/serjplus/cosmos-sdk/version.Name=HelloChain \
	-X github.com/serjplus/cosmos-sdk/version.ServerName=hcd \
	-X github.com/serjplus/cosmos-sdk/version.ClientName=hccli \
	-X github.com/serjplus/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/serjplus/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hcd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hccli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		go mod verify
