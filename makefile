# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=exchange-cli

MACOS_DIR=bin/macos
LINUX_DIR=bin/linux
WINDOWS_DIR=bin/windows

BUILD_DIR=bin

# Targets for cross-compilation
build-linux-amd64:
	mkdir -p $(LINUX_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_DIR)/$(BINARY_NAME)_amd64

build-linux-arm:
	mkdir -p $(LINUX_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(LINUX_DIR)/$(BINARY_NAME)_arm

build-macos-amd64:
	mkdir -p $(MACOS_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(MACOS_DIR)/$(BINARY_NAME)_amd64

build-macos-arm:
	mkdir -p $(MACOS_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(MACOS_DIR)/$(BINARY_NAME)_arm

build-windows-amd64:
	mkdir -p $(WINDOWS_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_DIR)/$(BINARY_NAME)_amd64.exe

build-windows-arm:
	mkdir -p $(WINDOWS_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(GOBUILD) -o $(WINDOWS_DIR)/$(BINARY_NAME)_arm.exe

# Build all targets
all: build-linux-amd64 build-windows-amd64 build-macos-amd64 build-linux-arm build-windows-arm build-macos-arm

# Default target
default: build-linux