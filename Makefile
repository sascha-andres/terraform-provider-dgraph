.PHONY: linux

linux: ## build linux plugin
	go build -o terraform-provider-dgraph
