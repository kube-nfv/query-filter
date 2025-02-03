NAME := query-filter

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

CONTAINER_TOOL ?= docker

# Antlr options
ANTLR_IMAGE ?= ghcr.io/kube-nfv/antlr
ANTLR_IMAGE_TAG ?= v4.13.2


QUERY_FILTER_GRAMMAR_NAME ?= Filter.g4
QUERY_FILTER_GRAMMAR_DIR ?= antlr/grammar/
QUERY_FILTER_GRAMMAR = $(QUERY_FILTER_GRAMMAR_DIR)/$(QUERY_FILTER_GRAMMAR_NAME)
ANTLR_GEN_DIR ?= antlr/generate

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: all
all: build

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: generate
generate: antlr-generate ## Generate golang files
	go generate ./...

.PHONY: mod-tidy
mod-tidy: ## Run go mod tidy against code.
	go mod tidy

.PHONY: test
test:
	go test -v

.PHONY: antlr-generate
antlr-generate: $(ANTLR_GEN_DIR) $(QUERY_FILTER_GRAMMAR)
	$(CONTAINER_TOOL) run --rm -v "$(PWD)/antlr:/work" $(ANTLR_IMAGE):$(ANTLR_IMAGE_TAG) -o generate -Dlanguage=Go grammar/$(QUERY_FILTER_GRAMMAR_NAME)
	mv $(ANTLR_GEN_DIR)/grammar/*.go $(ANTLR_GEN_DIR)
	rm -rf $(ANTLR_GEN_DIR)/grammar

$(ANTLR_GEN_DIR):
	mkdir -p $(ANTLR_GEN_DIR)

.PHONY: clean
clean:
	@rm -rf $(ANTLR_GEN_DIR)

