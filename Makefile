APP ?= items-api
BUILD_DIR ?= bin
CONFIG ?= etc/items-api.yaml
GO ?= go

.PHONY: help build run test fmt vet tidy clean docker-build docker-run k8s-apply

help:
	@printf "Targets:\n"
	@printf "  build        Build the binary into %s/%s\n" "$(BUILD_DIR)" "$(APP)"
	@printf "  run          Run locally with CONFIG=%s\n" "$(CONFIG)"
	@printf "  test         Run all tests\n"
	@printf "  fmt          Format Go code\n"
	@printf "  vet          Run go vet\n"
	@printf "  tidy         Tidy go.mod/go.sum\n"
	@printf "  clean        Remove build output\n"
	@printf "  docker-build Build the Docker image\n"
	@printf "  docker-run   Run the Docker image (port 8888)\n"
	@printf "  k8s-apply    Apply K8s manifests\n"

build:
	@mkdir -p "$(BUILD_DIR)"
	"$(GO)" build -o "$(BUILD_DIR)/$(APP)" ./main.go

run:
	"$(GO)" run ./main.go -f "$(CONFIG)"

test:
	"$(GO)" test ./...

fmt:
	"$(GO)" fmt ./...

vet:
	"$(GO)" vet ./...

tidy:
	"$(GO)" mod tidy

clean:
	rm -rf "$(BUILD_DIR)"

docker-build:
	docker build -t "$(APP):latest" .

docker-run:
	docker run --rm -p 8888:8888 "$(APP):latest"

k8s-apply:
	kubectl apply -f k8s/configmap.yaml
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl apply -f k8s/hpa.yaml
