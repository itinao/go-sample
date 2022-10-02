curdir := `pwd`

.DEFAULT_GOAL := help

build: ## build
		docker build . -t go-sample

run: ## run
		docker run -p 8080:8080 go-sample

help: ## Self-documented Makefile
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
			| sort \
			| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
