.PHONY: help
help:
	@echo "Usage:"
	@echo "Protobuf commands"
	@echo "    install         Install protoc programs"
	@echo "    generate        Generate proto files"
	@echo "Git commands"
	@echo "    push            Push changes to github"
	@echo "    pull            Pull in commits from github"
	@echo "Makefile commands"
	@echo "    cert            Generate all the necessary cert files"
	@echo "    build           Build this project locally"
	@echo "    update          go update libraries"
	@echo "    test            Perform go testing"
	@echo "    clean           Clean this project and database docker volume"
	@echo "    docker          Build all services docker images"
	@echo "Docker commands"
	@echo "    docker-up       Docker-compose up"
	@echo "    docker-down     Docker-compose down"

SERVICES = gateway proxy user company

.PHONY: install
install:
	./script/proto-install.sh

.PHONY: generate
generate:
	./script/proto-gen.sh

.PHONY: push
push:
	git push

.PHONY: pull
pull:
	git pull

# generate both server and client cert for mTLS communication
.PHONY: cert
cert:
	./script/cert.sh -c -v -o certs/localhost localhost 127.0.0.1

# build the source to native OS and platform
.PHONY: build
build:
	@for service in $(SERVICES) ; do \
		echo "Building $$service..." ; \
		go build -ldflags '-extldflags "-static"' -o $$service cmd/$$service/*.go ; \
	done

# go update libraries
.PHONY: update
update:
	@for service in $(SERVICES) ; do \
		echo "Updating $$service..." ; \
		go get -u ./...
		go mod tidy
	done

.PHONY: test
test:
	@for service in $(SERVICES) ; do \
		echo "Testing $$service..." ; \
	done

.PHONY: clean
clean:
	rm -rf lib/*
	@for service in $(SERVICES) ; do \
		echo "Testing $$service..." ; \
		rm -f $$service ; \
	done

# build the docker image
.PHONY: docker
docker:
	docker-compose -f docker-compose.yml build

# docker-compose up
.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yml up

# docker-compose down
.PHONY: docker-down
docker-down:
	docker-compose -f docker-compose.yml down
