WEB_PACKAGE_PATH := ./cmd/webserver
WEB_BINARY_NAME := serverManager

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tailwind: compile tailwind css
.PHONY: tailwind
tailwind:
	npx @tailwindcss/cli -i ./ui/css/input.css -o ./ui/static/app.css --watch

## air: automatically recompile on change
.PHONY: air
air:
	air

## run-web: run the web
.PHONY: run-web
run-web:
	go run ${WEB_PACKAGE_PATH}

## build: build the application
.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/${WEB_BINARY_NAME := serverManager

