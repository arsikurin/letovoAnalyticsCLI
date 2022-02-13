.PHONY: build

MAIN_SRC = ./src/main.go
OUTPUT_SRC = ./bin/letovo

build:
	go build -o $(OUTPUT_SRC) $(MAIN_SRC)

