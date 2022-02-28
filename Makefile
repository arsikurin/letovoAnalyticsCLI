.PHONY: build build-mac-arm64 build-mac-amd64 build-linux-amd64 build-win-amd64

MAIN_SRC = ./src/
OUTPUT_SRC = ./bin/letovo

# build for current OS
build:
	go build -o $(OUTPUT_SRC) $(MAIN_SRC)

# build for windows AMD64
build-win-amd64:
	env GOOS=windows GOARCH=amd64 CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build -o $(OUTPUT_SRC).exe $(MAIN_SRC)

# build for macOS ARM64
build-mac-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_SRC) $(MAIN_SRC)

# build for macOS AMD64
build-mac-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_SRC) $(MAIN_SRC)

# build for linux AMD64
build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_SRC) $(MAIN_SRC)

# build for linux ARM64
build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o $(OUTPUT_SRC) $(MAIN_SRC)
