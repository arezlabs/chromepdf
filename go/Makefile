# Makefile for building and running the ChromePDF project

GO_BIN_PATH := go/bin
GO_SRC_PATH := go/src
OUTPUT_BINARY := $(GO_BIN_PATH)/chrome-pdf

.PHONY: all build clean

all: build

# Build the Go binary
build:
	mkdir -p $(GO_BIN_PATH)
	go build -o $(OUTPUT_BINARY) ./$(GO_SRC_PATH)

# Clean up the binaries
clean:
	rm -rf $(GO_BIN_PATH)/*