export GO111MODULE=on
BINARY_NAME=balkonygardener
BINARY_FILE_PATH="bin/$(BINARY_NAME)"

setup: ## Install all the build and lint dependencies
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s v1.27.0
	go get -u golang.org/x/tools/cmd/goimports

test-with-coverage: ## Run all the tests
	rm -f coverage.tmp && rm -f coverage.txt
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -race -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

test: ## Run all the tests
	go version
	go env
	go list ./... | xargs -n1 -I{} sh -c 'go test -race {}'

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt-all: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

fmt: ## gofmt and goimports all uncommited go files
	 git diff --name-only | grep .go | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done


lint: ## Run the linters
	golangci-lint run

build: ## build binary to .build folder
	rm -f $(BINARY_FILE_PATH) 
	go build -o $(BINARY_FILE_PATH) app/main.go

install-service:
	sudo service balkonygardener stop || true
	sudo rm /etc/systemd/system/balkonygardener.service || true
	sudo ln -s /home/pi/go-workspace/src/github.com/Oppodelldog/balkonygardener/balkonygardener.service /etc/systemd/system/balkonygardener.service
	sudo systemctl enable /etc/systemd/system/balkonygardener.service
	sudo service balkonygardener start

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help