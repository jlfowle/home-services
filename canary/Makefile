APP := canary
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

APPBIN := bin/$(APP)-$(GOOS)-$(GOARCH)

$(APPBIN):
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APPBIN) cmd/$(APP)/$(APP).go

.PHONY: build
build: $(APPBIN)

.PHONY: run
run: $(APPBIN)
	$(APPBIN)