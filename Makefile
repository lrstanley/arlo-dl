.DEFAULT_GOAL := build

DIRS=bin dist
BINARY=arlo-dl

VERSION=$(shell git describe --tags --always --abbrev=0 --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)
VERSION_FULL=$(shell git describe --tags --always --dirty --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)

RSRC=README_TPL.md
ROUT=README.md

$(info $(shell mkdir -p $(DIRS)))
BIN=$(CURDIR)/bin
export GOBIN=$(CURDIR)/bin

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

fetch: ## Fetches the necessary dependencies to build.
	which $(BIN)/goreleaser 2>&1 > /dev/null || wget -qO- "https://github.com/goreleaser/goreleaser/releases/download/v0.122.0/goreleaser_Linux_x86_64.tar.gz" | tar -xz -C $(BIN) goreleaser
	go mod download
	go mod tidy
	go mod vendor

upgrade-deps: ## Upgrade all dependencies to the latest version.
	go get -u ./...

upgrade-deps-patch: ## Upgrade all dependencies to the latest patch release.
	go get -u=patch ./...

readme-gen: ## Generates readme from template file.
	cp -av "${RSRC}" "${ROUT}"
	sed -ri -e "s:\[\[tag\]\]:${VERSION}:g" -e "s:\[\[os\]\]:linux:g" -e "s:\[\[arch\]\]:amd64:g" "${ROUT}"

snapshot: clean fetch generate ## Generate a snapshot release.
	$(BIN)/goreleaser --snapshot --skip-validate --skip-publish

release: clean fetch generate ## Generate a release, but don't publish to GitHub.
	$(BIN)/goreleaser --skip-validate --skip-publish

publish: clean fetch generate ## Generate a release, and publish to GitHub.
	$(BIN)/goreleaser

clean: ## Cleans up generated files/folders from the build.
	/bin/rm -rfv "dist/" "${BINARY}-${VERSION_FULL}"

build: fetch clean ## Compile and generate a binary with static assets embedded.
	go build -ldflags '-d -s -w' -tags netgo -installsuffix netgo -v -o "${BINARY}-${VERSION_FULL}"

debug: clean
	go run *.go
