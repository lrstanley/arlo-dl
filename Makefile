.DEFAULT_GOAL := build-all

export PROJECT := "arlo-dl"
export PACKAGE := "github.com/lrstanley/arlo-dl"

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

build-all: clean go-fetch go-build
	@echo

up: go-upgrade-deps
	@echo

clean:
	/bin/rm -rfv "dist/" "${PROJECT}"

go-prepare: license go-fetch
	go generate -x ./...

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

go-dlv: go-prepare
	dlv debug \
		--headless --listen=:2345 \
		--api-version=2 --log \
		--allow-non-terminal-interactive \
		${PACKAGE} --

go-debug: go-prepare
	go run ${PACKAGE}

go-build: go-prepare go-fetch
	CGO_ENABLED=0 \
	go build \
		-ldflags '-d -s -w -extldflags=-static' \
		-tags=netgo,osusergo,static_build \
		-installsuffix netgo \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}
