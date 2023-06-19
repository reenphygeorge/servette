GO := go

EXECUTABLE := light-server

build-linux:
	GOOS=linux GOARCH=amd64
	$(GO) build -o $(EXECUTABLE) ./cmd/main.go

build-windowsx86:
	GOOS=windows GOARCH=386
	$(GO) build -o $(EXECUTABLE).exe ./cmd/main.go

build-windowsx86-64:
	GOOS=windows GOARCH=64
	$(GO) build -o $(EXECUTABLE).exe ./cmd/main.go

build-macos:
	GOOS=darwin GOARCH=amd64
	$(GO) build -o $(EXECUTABLE) ./cmd/main.go

run:
	$(GO) build -o $(EXECUTABLE) ./cmd/main.go
	./$(EXECUTABLE)

clean:
	rm -f $(EXECUTABLE)
