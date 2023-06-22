GO := go

EXECUTABLE := ./build/servette

build-linux:
	GOOS=linux GOARCH=amd64
	$(GO) build -o $(EXECUTABLE)-linux ./cmd/main.go

build-windows-x86:
	GOOS=windows GOARCH=386
	$(GO) build -o $(EXECUTABLE)-win-x86.exe ./cmd/main.go

build-windows-x86-64:
	GOOS=windows GOARCH=amd64
	$(GO) build -o $(EXECUTABLE)-win-x86-64.exe ./cmd/main.go

build-macos:
	GOOS=darwin GOARCH=amd64
	$(GO) build -o $(EXECUTABLE)-darwin ./cmd/main.go

build-all:
	make build-linux
	make build-windows-x86
	make build-windows-x86-64
	make build-macos
	cp lsconfig.json ./build

run:
	$(GO) run ./cmd/main.go

clean-all:
	rm -rf ./build
