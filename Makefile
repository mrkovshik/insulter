GO=go
export GOPRIVATE=gitlab.com/tabby.ai/*
TARGET_DIR?=$(PWD)/.build
M=$(shell printf "\033[34;1m>>\033[0m")


.PHONY: build
build: 
	$(info $(M) building service...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET_DIR)/service ./cmd/*.go

.PHONY: run
run: build 
	$(TARGET_DIR)/service


.PHONY: evans
evans:
	evans grpc/proto/calculus.proto -p 8080

.PHONY: tidy
tidy:
	go mod tidy
%:
	@:

