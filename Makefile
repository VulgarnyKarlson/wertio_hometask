GOCMD=go
GOBUILD=$(GOCMD) build
CMDPATH=cmd/main.go
BINPATH=app
PROJECTNAME=wertio_hometask

.PHONY: build-linux
build-linux: ## build-linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w" -o ${BINPATH} ${CMDPATH}


.PHONY: build
build: ## build - service binary
	CGO_ENABLED=0 $(GOBUILD) -o ${BINPATH} ${CMDPATH}


.PHONY: clean
clean: ## clean - remove binary file
	rm -f ${BINPATH}

.PHONY: run
run: ## run - service locally
	$(GOCMD) run ${CMDPATH}

.PHONY: run-race
run-race: ## run-race - with race detector
	$(GOCMD) run -race ${CMDPATH}


.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
