LINTER_CFG = ./config-dev/.golangci.yml
PKG = ./...

# Docker app
IMAGE = app
CONTAINER = http-highload
PORT = 8088

# DEV
create-env:
	go mod download

run-linter:
	docker run --rm -v $(shell pwd):/app -v ~/.cache/golangci-lint/:/root/.cache -w /app golangci/golangci-lint:latest golangci-lint run $(PKG) --config=$(LINTER_CFG) -v
	go fmt $(PKG)

build:
	rm -f bin/main.bin
	go build -o bin/main.bin cmd/app/main.go

# easyjson
generate:
	go generate ${PKG}

clear:
	sudo rm -rf fullchain.pem privkey.pem bin/*


# Env export POD_UUID=8 && export PORT_APP=8080 && export PORT_METRICS_APP=9000
# INFRA
build-app:
	docker-compose build app

run-app:
	docker-compose run --rm -d --name app_1_8080_9000 -e POD_UUID=8 -e PORT_APP=:8080 -e PORT_METRICS_APP=:9000 -p 8080:8080 -p 9000:9000 app

SERVICE_NAME=app
POD=1
PORT_APP=8080
PORT_METRICS_APP=9000
SIZE_CLUSTER=10

run-cluster-app:
	./scripts/run_cluster_app.sh ${SERVICE_NAME} ${POD} ${PORT_APP} ${PORT_METRICS_APP} ${SIZE_CLUSTER}

get-list-apps:
	 docker-compose ps -a

# Example: make get-logs-app CONTAINER_NAME=app_1_8080_9000
get-logs-app:
	docker logs ${CONTAINER_NAME}

stop:
	docker-compose kill
	docker-compose down