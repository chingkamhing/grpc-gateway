.PHONY: help
help:
	@echo "Usage:"
	@echo "    install         Install protoc programs"
	@echo "    generate        Generate proto files"
	@echo "    gomod           Perform go mod init or update"
	@echo "    build           Build this project"
	@echo "    clean           Clean all generated proto files"

.PHONY: install
install:
	./proto-install.sh

.PHONY: generate
generate:
	./proto-gen.sh

.PHONY: gomod
gomod:
	./go-mod.sh

# build the source to native OS and platform
.PHONY: build
build:
	go build -ldflags '-extldflags "-static"' -o gateway cmd/gateway/*.go
	go build -ldflags '-extldflags "-static"' -o user cmd/user/*.go
	go build -ldflags '-extldflags "-static"' -o company cmd/company/*.go

.PHONY: clean
clean:
	rm -rf lib/*
	rm -f gateway
	rm -f user
	rm -f company
