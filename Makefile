EXECUTABLE=en-learner
WINDOWS=bin/$(EXECUTABLE)_windows_amd64.exe
LINUX=bin/$(EXECUTABLE)_linux_amd64
DARWIN=bin/$(EXECUTABLE)_darwin_amd64
DARWIN_ARM=bin/$(EXECUTABLE)_darwin_arm64

.PHONY: all test clean

all: test build ## Build and run tests

bin:
	mkdir -p bin/

test: ## Run unit tests
	go test ./...

build: windows linux darwin darwin-arm## Build binaries

windows: bin $(WINDOWS) ## Build for Windows

linux: bin $(LINUX) ## Build for Linux

darwin: bin $(DARWIN) ## Build for Darwin (macOS)

darwin-arm: bin $(DARWIN_ARM) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -extldflags "-static""  ./cmd/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -extldflags "-static""  ./cmd/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -extldflags "-static""  ./cmd/main.go

$(DARWIN_ARM):
	env GOOS=darwin GOARCH=arm64 go build -v -o $(DARWIN_ARM) -ldflags="-s -w -extldflags "-static""  ./cmd/main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN) $(DARWIN_ARM)