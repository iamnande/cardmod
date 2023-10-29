.PHONY: default help

default: help
help: ## help: display make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-20s -> %s\n\033[0m", $$1, $$2}'

# --------------------------------------------------
# Application Information
# --------------------------------------------------
APP_NAME     := cardmod
APP_WORKDIR  := $(shell pwd)
APP_PACKAGES := $(shell go list -f '{{.Dir}}' ./...)
APP_LOG_FMT  := `/bin/date "+%Y-%m-%d %H:%M:%S %z [$(APP_NAME)]"`

# --------------------------------------------------
# Runtime Targets
# --------------------------------------------------

# --------------------------------------------------
# Build Targets
# --------------------------------------------------
BUILD_DIR := $(APP_WORKDIR)/build

.PHONY: build-clean
build-clean: ## build: clean build workspace
	@echo $(APP_LOG_FMT) "cleaning build workspace"
	@rm -rf $(BUILD_DIR)

.PHONY: build-api
build-api: build-clean ## build: build binary file
	@echo $(APP_LOG_FMT) "building binary"
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
		-o $(BUILD_DIR)/$(APP_NAME) -ldflags '-extldflags "-static"' \
		cmd/$(APP_NAME)/main.go

# --------------------------------------------------
# Test Targets
# --------------------------------------------------
COVERAGE_DIR := $(APP_WORKDIR)/coverage

.PHONY: test-lint
test-lint: ## test: check for lint failures
	@echo $(APP_LOG_FMT) "checking for lint failures"
	@golangci-lint run --timeout=5m --verbose

.PHONY: test-unit
test-unit: export UNIT_DIR=$(COVERAGE_DIR)/unit
test-unit: export UNIT_WEBPAGE=$(UNIT_DIR)/index.html
test-unit: export UNIT_COVERAGE=$(UNIT_DIR)/coverage.txt
test-unit: ## test: execute unit test suite
	@echo $(APP_LOG_FMT) "executing unit test suite"
	@mkdir -p $(UNIT_DIR)
	@go test -v \
		-race \
		-covermode=atomic \
		-coverprofile=$(UNIT_COVERAGE) \
		$(APP_PACKAGES)
	@go tool cover -func=$(UNIT_COVERAGE)
	@go tool cover -html=$(UNIT_COVERAGE) -o $(UNIT_WEBPAGE)

.PHONY: test-integration
test-integration: export INTEGRATION_DIR=$(COVERAGE_DIR)/unit
test-integration: export INTEGRATION_WEBPAGE=$(INTEGRATION_DIR)/index.html
test-integration: export INTEGRATION_COVERAGE=$(INTEGRATION_DIR)/coverage.txt
test-integration: ## test: execute integration test suite
	@echo $(APP_LOG_FMT) "executing integration test suite"
	@mkdir -p $(INTEGRATION_DIR)
	@go test -v \
		-race \
		-tags=integration \
		-covermode=atomic \
		-coverprofile=$(INTEGRATION_COVERAGE) \
		$(APP_PACKAGES)
	@go tool cover -func=$(INTEGRATION_COVERAGE)
	@go tool cover -html=$(INTEGRATION_COVERAGE) -o $(INTEGRATION_WEBPAGE)
