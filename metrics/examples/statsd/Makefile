IMAGE_NAME=golang:1.10.0

HOST_PKG_DIR=$(abspath $(shell pwd)/../../..)
PKG_DIR=/go/src/github.com/dcos/dcos-go
MAIN_DIR=$(PKG_DIR)/metrics/examples/statsd
BUILD_DIR=$(MAIN_DIR)/build
BINARY=statsd-generator

windows_EXE=.exe

.PHONY: default
default: fmt
	@make $(shell uname | tr [A-Z] [a-z])

.PHONY: darwin linux windows
darwin linux windows:
	$(call run_docker,GOOS=$(@) go build -o $(BUILD_DIR)/$(@)/$(BINARY)$($(@)_EXE) .)

.PHONY: fmt
fmt:
	$(call run_docker,cd $(PKG_DIR) && go fmt ./...)

.PHONY: vendor
vendor:
	$(call run_docker,cd $(PKG_DIR) && go get -u github.com/golang/dep/... && dep ensure)

.PHONY: clean
clean:
	rm -rf build

define run_docker
	docker run \
		-v $(HOST_PKG_DIR):$(PKG_DIR) \
		-w $(MAIN_DIR) \
		--rm \
		$(IMAGE_NAME) \
		bash -c "$(1)"
endef
