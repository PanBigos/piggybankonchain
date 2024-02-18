SOLC_VERSION := 0.8.24-alpine
GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

ARCHITECTURE := $(shell uname -m)

ifeq ($(ARCHITECTURE),x86_64)
    DOCKER_COMPOSE_FILE := infra/deployment/docker-compose.dev.x86.yaml
else ifeq ($(ARCHITECTURE),aarch64)
    DOCKER_COMPOSE_FILE := infra/deployment/docker-compose.dev.arm.yaml
else ifeq ($(ARCHITECTURE),arm64)
    DOCKER_COMPOSE_FILE := infra/deployment/docker-compose.dev.arm.yaml
else
    $(error Unsupported architecture: $(ARCHITECTURE))
endif

sqlc:
	docker pull sqlc/sqlc
	docker run -u $(id -u ${USER}):$(id -g ${USER}) --rm -v "${PWD}:/src" -w /src sqlc/sqlc generate

.PHONY: sqlc mock

generate:
	buf generate

generate.contracts.piggy:
	docker pull ethereum/solc:$(SOLC_VERSION)
	docker pull ethereum/client-go:alltools-latest
	docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v ./contracts/piggy:/contracts/piggy/ ethereum/solc:${SOLC_VERSION} --abi /contracts/piggy/fuse/TimeLockedPiggyBankFactory.sol /contracts/piggy/fuse/PiggyBankRouter.sol -o /contracts/piggy/abi/ --overwrite

	docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v ./contracts:/contracts/ ethereum/client-go:alltools-latest abigen --abi=/contracts/piggy/abi/TimeLockedPiggyBankFactory.abi --pkg=factory --out=/contracts/piggy/factory/piggy.go
	docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v ./contracts:/contracts/ ethereum/client-go:alltools-latest abigen --abi=/contracts/piggy/abi/PiggyBankRouter.abi --pkg=router --out=/contracts/piggy/router/piggy.go
docker-image:
	docker compose -f infra/build/docker-compose.yaml build 

docker-up:
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d 

docker-down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

node.forge.start:
	cd ./tools/forge && ./start.sh

node.forge.down:
	cd ./tools/forge && ./down.sh

service:
	go build -o $(GOBIN)/piggy ./cmd/piggy/ 
	@echo "Run \"$(GOBIN)/piggy\" to launch piggy service."

test:
	go test -v -cover -short ./...

mock:
	mockery
	
.PHONY: generate node.forge docker-image docker-up docker-down service mock 