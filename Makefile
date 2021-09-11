.PHONY: help
help:
	@echo "Usage:"
	@echo "    install              Install protoc programs"
	@echo "    generate             Generate proto files"
	@echo "    build                Build this project locally"
	@echo "    test                 Perform go testing"
	@echo "    docker               Build all services docker images"
	@echo "    clean                Clean this project and database docker volume"

.PHONY: install
install:
	./script/proto-install.sh

.PHONY: generate
generate:
	./script/proto-gen.sh

# build the source to native OS and platform
.PHONY: build
build:
	go build -ldflags '-extldflags "-static"' -o gateway cmd/gateway/*.go
	go build -ldflags '-extldflags "-static"' -o user cmd/user/*.go
	go build -ldflags '-extldflags "-static"' -o company cmd/company/*.go

.PHONY: test
test:
	# nothing to do yet

.PHONY: docker
docker:
	# individual docker for each service
	docker-compose -f docker-compose.yml build

.PHONY: clean
clean:
	rm -rf lib/*
	rm -f gateway
	rm -f user
	rm -f company
