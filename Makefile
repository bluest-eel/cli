VERSION = $(shell cat VERSION)
GO ?= go
GOFMT ?= $(GO)fmt
GO111MODULE = on

FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
DEFAULT_GOPATH = $(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN = $(DEFAULT_GOPATH)/bin
export PATH := $(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo

DVCS_HOST = github.com
ORG = bluest-eel
PROJ = cli
BIN = $(ORG)
FQ_PROJ = $(DVCS_HOST)/$(ORG)/$(PROJ)

LD_VERSION = -X $(FQ_PROJ)/common.Version=$(VERSION)
LD_BUILDDATE = -X $(FQ_PROJ)/common.BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_GITCOMMIT = -X $(FQ_PROJ)/common.GitCommit=$(shell git rev-parse --short HEAD)
LD_GITBRANCH = -X $(FQ_PROJ)/common.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)
LD_GITSUMMARY = -X $(FQ_PROJ)/common.GitSummary=$(shell git describe --tags --dirty --always)

LDFLAGS = -w -s $(LD_VERSION) $(LD_BUILDDATE) $(LD_GITBRANCH) $(LD_GITSUMMARY) $(LD_GITCOMMIT)

#############################################################################
###   Core   ################################################################
#############################################################################

default: all

all-common: clean check-deps lint build
all: all-common test
all-cicd: all-common test-nocolor

#############################################################################
###   Build   ###############################################################
#############################################################################

version:
	@echo $(VERSION)

clean:
	@echo '>> Removing project binaries ...'
	@rm -f bin/$(BIN)

clean-cache:
	@echo '>> Purging Go mod cahce ...'
	@GO111MODULE=$(GO111MODULE) $(GO) clean -cache

clean-all: clean clean-cache

deps:
	@echo '>> Downloading deps ...'
	@GO111MODULE=$(GO111MODULE) $(GO) get -v -d ./...

bin:
	@mkdir ./bin

bin/$(BIN): bin
	@echo '>> Building CLI binary'
	@GO111MODULE=$(GO111MODULE) $(GO) build -ldflags "$(LDFLAGS)" -o bin/$(BIN) ./cmd/$(BIN)

build-client: | bin/$(BIN)
build: bin build-client
build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(MAKE) build

rebuild: clean build

#############################################################################
###   Linting & Testing   ###################################################
#############################################################################

$(GOLANGCI_LINT):
	@echo ">> Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.15.0

show-linter:
	@echo $(GOLANGCI_LINT)

lint: $(GOLANGCI_LINT)
	@echo '>> Linting source code'
	@GL_DEBUG=linters_output GOPACKAGESPRINTGOLISTERRORS=1 GO111MODULE=$(GO111MODULE) \
	$(GOLANGCI_LINT) \
	--enable=golint \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	run ./...

$(RICH_GO):
	@echo ">> Couldn't find $(RICH_GO); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=$(GO111MODULE) \
	$(GO) get -u github.com/kyoh86/richgo

test: $(RICH_GO)
	@echo '>> Running all tests'
	@GO111MODULE=$(GO111MODULE) $(RICH_GO) test ./... -v

test-nocolor:
	@echo '>> Running all tests'
	@GO111MODULE=$(GO111MODULE) $(GO) test ./... -v

#############################################################################
###   Release Process   #####################################################
#############################################################################

tag:
	@echo "Tags:"
	@git tag|tail -5
	@git tag "v$(VERSION)"
	@echo "New tag list:"
	@git tag|tail -6

tag-and-push: tag
	@git push --tags

tag-delete: VERSION ?= 0.0.0
tag-delete:
	@git tag --delete v$(VERSION)
	@git push --delete origin v$(VERSION)

#############################################################################
###   Misc   ################################################################
#############################################################################

# If the make target has `run` in the name
ifeq (run,$(findstring run,$(MAKECMDGOALS)))
  # use the rest as arguments for `*run*` ...
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ... and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

run-debug:
	@BLUEST_EEL_LOGGING_REPORT_CALLER=true \
	BLUEST_EEL_LOGGING_COLORED=true \
	BLUEST_EEL_LOGGING_LEVEL=debug \
	./bin/$(BIN) $(RUN_ARGS)

run-trace:
	@GRPC_TRACE=all \
	GRPC_VERBOSITY=DEBUG \
	GRPC_GO_LOG_VERBOSITY_LEVEL=2 \
	GRPC_GO_LOG_SEVERITY_LEVEL=info \
	BLUEST_EEL_LOGGING_REPORT_CALLER=true \
	BLUEST_EEL_LOGGING_COLORED=true \
	BLUEST_EEL_LOGGING_LEVEL=trace \
	./bin/$(BIN) $(RUN_ARGS)

show-ldflags:
	@echo $(LDFLAGS)

check-deps:
	@GO111MODULE=on $(GO) mod tidy
	@#@GO111MODULE=on $(GO) mod verify

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | \
	awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
	sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
