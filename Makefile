.PHONY: linux

linux: ## build linux plugin
	go build -o terraform-provider-dgraph
	cp terraform-provider-dgraph ~/.terraform.d/plugins/

start: ## run a local dgraph server
	docker run --name dgraph --rm -p 8000:8000 -p 8080:8080 -p 9080:9080 dgraph/standalone:latest || /bin/true

stop: ## stop docker images
	docker stop dgraph
