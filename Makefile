NAME_API=shortener-api
VERSION=dev
OS ?= linux
PROJECT_PATH ?= github.com/alexmourapb/url-shortener
PKG ?= github.com/alexmourapb/url-shortener/cmd
REGISTRY ?= alexmourapb
TERM=xterm-256color
CLICOLOR_FORCE=true
RICHGO_FORCE_COLOR=1

.PHONY: test
test:
	@echo "==> Running Tests"
	go test -v ./...

.PHONY: compile
compile: clean
	@echo "==> Go Building API"
	@env GOOS=${OS} GOARCH=amd64 go build -v -o build/${NAME_API} ${PKG}/${NAME_API}

.PHONY: build
build: compile
	@echo "==> Building Docker API image"
	@docker build -t ${REGISTRY}/${NAME_API}:${VERSION} build -f build/Dockerfile

.PHONY:
push:
	@echo "==> Pushing to registry"
	@docker push ${REGISTRY}/${NAME_API}:${VERSION}

.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@GOOS=${OS} go clean -i -x ./...
	@rm -f build/${NAME_API}

.PHONY: metalint
metalint:
	@echo "==> installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	go install ./...
	go test ./...
	$$(go env GOPATH)/bin/golangci-lint run -c ./.golangci.yml ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running test's"
	@richgo test -failfast -coverprofile=coverage.out ./...
	go get github.com/matryer/moq@v0.1.3
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: generate
generate:
	@echo "Go Generating"
	go get -u github.com/swaggo/swag/cmd/swag@latest
	go generate ./...
	go mod tidy
