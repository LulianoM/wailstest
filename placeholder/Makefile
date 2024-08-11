##########################################################################################
# Metadata
##########################################################################################

BIN_NAME={{app_name}}

##########################################################################################
# Directories
##########################################################################################

CMD_DIR := ./cmd
CMD_API_DIR := $(CMD_DIR)/api
CMD_CONSUMER_DIR := $(CMD_DIR)/consumer

MIGRATION_DIR := ./migrations

TOOLS_DIR := ./tools
TOOLS_DEPENDECIES_DIR = $(TOOLS_DIR)/deps
TOOLS_BUILD_DIR ?= $(TOOLS_DIR)/builds
TOOLS_MOCK_DIR ?= $(TOOLS_DIR)/mocks

DEPLOYMENTS_DIR := ./deployments
DEPLOYMENTS_COMMON_DIR := $(DEPLOYMENTS_DIR)/common
DEPLOYMENTS_TEST_DIR := $(DEPLOYMENTS_DIR)/test
DEPLOYMENTS_LOCAL_DIR := $(DEPLOYMENTS_DIR)/local
DEPLOYMENTS_ACTION_DIR := $(DEPLOYMENTS_DIR)/action

COVER_DIR ?=  ./.cover
COVER_ENTRY ?= ./...
COVER_IGNORED_FILES ?= "src/tests\|mocks/\|cmd/\|migrations/"


##########################################################################################
# Files
##########################################################################################

FILE_COVER_UNIT ?= unit.out
FILE_COVER_INTEGRATION ?= integration.out

##########################################################################################
# Commands
##########################################################################################

GO ?= go
GET ?= $(GO) get
GORUN ?= $(GO) run
GOTEST ?= go test
GOTOOL ?= $(GO) tool

##########################################################################################
# Build
##########################################################################################

.PHONY: build/docker
build/docker:
	@docker build \
		--tag {{app_name}}:local \
		--build-arg ssh_prv_key="`cat ~/.ssh/id_rsa`" \
		--build-arg ssh_pub_key="`cat ~/.ssh/id_rsa.pub`" \
		-f ${DEPLOYMENTS_COMMON_DIR}/Dockerfile .

##########################################################################################
# Setup
##########################################################################################

.PHONY: setup/local
setup/local:
	@$(GO) get -u github.com/rakyll/gotest &&\
	$(GO) get -u github.com/axw/gocov/gocov &&\
	$(GO) install github.com/golang/mock/mockgen@v1.6.0 &&\
    $(GO) mod tidy

.PHONY: setup/cover
setup/cover:
	@rm -rf ${COVER_DIR}
	@mkdir -p ${COVER_DIR}


##########################################################################################
# Simplifiers
##########################################################################################

.PHONY: run/api
run/api: 
	@$(GORUN) ${CMD_API_DIR}/main.go

.PHONY: deps-up
deps-up: deps/local/up

.PHONY: deps-down
deps-down: deps/local/down

.PHONY: test-unit
test-unit: test/local/unit

.PHONY: test-integration
test-integration: test/local/live-integration