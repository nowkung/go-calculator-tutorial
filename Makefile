PROJECTNAME := $(shell basename "$(PWD)")
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64

## help: helper
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Project: ["$(PROJECTNAME)"]"
	@echo " Please choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## run: execute main go application in local
.PHONY: run
run:
	export APPENV=local; go run -race cmd/main.go

# ## up: start the application using docker compose
# .PHONY: up
# up:
# 	docker-compose -f ./docker-compose.yml up -d

# ## down: down docker compose service
# .PHONY: down
# down:
# 	docker compose -f ./docker-compose.yml down  --volumes

# ## tidy: special go mod tidy without golang database checksum(GOSUMDB)
# .PHONY: tidy
# tidy:
# 	export GOSUMDB=off ; go mod tidy

## test: run go test
.PHONY: test
test:
	go clean -testcache & go test -v -race ./...

# ## set_private_repo_global: set a "gitdev.devops.krungthai.com" to be a private repo in go global environment
# set_private_repo_global:
# 	go env -w GOPRIVATE="gitdev.devops.krungthai.com/*"

# ## update_standard_lib: update standard library (glo-standard-library) with GOPRIVATE option
# update_standard_lib:
# 	GOPRIVATE=gitdev.devops.krungthai.com/glo/glo-standard-library go get gitdev.devops.krungthai.com/glo/glo-standard-library

# ## gosec: run for scan code vulnerability by securego/gosec
# .PHONY: gosec 
# gosec: 
# 	gosec ./... 

# ## govulncheck: run for scan vulnerability package from Go vulnerability database
# .PHONY: govulncheck
# govulncheck: 
# 	govulncheck ./... 

# ## security: run make gosec and make govulncheck
# security: gosec govulncheck

# ## gogen: run all go genrate(mockgen, etc)
# gogen:
# 	go generate ./...
