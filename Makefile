SHELL := /bin/bash
.PHONY: test run help

help: ## Show this help message.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

make run: ## Run main.go file
	@go run main.go

test: ## Execute all test
	@go test -v

make docker-build: ## Build docker image locally
	@docker build -t juliopedrosa/webapp:latest

make docker-run:
	@docker run -p 8080:3000 -d juliopedrosa/webapp:latest

release: ## Create release
	./scripts/create_release.sh

hotfix: ## Create hotfix
	./scripts/create_hotfix.sh
