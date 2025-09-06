PROJECT := $(shell cat go.mod | grep module | awk -F ' ' '{print $$2}' | awk -F '/' '{print $$NF}')

BUILD := build

help:
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Настроить рабочее окружение
	go mod tidy
	go mod vendor

clean: ## Очистить рабочее окружение
	rm -rf vendor
	go clean -r -i -testcache -modcache

build-windows: setup ## Собрать проект под Windows
	GOOS=windows GOARCH=amd64 go build -o $(BUILD)/$(PROJECT).exe main.go

build-linux: setup ## Собрать проект под Linux
	GOOS=linux GOARCH=amd64 go build -o $(BUILD)/$(PROJECT) main.go

build-darwin: setup ## Собрать проект под macOS
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD)/$(PROJECT) main.go
