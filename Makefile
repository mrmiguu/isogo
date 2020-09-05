.PHONY: build

build:
	tinygo build -o main.wasm -no-debug -target wasm ./main.go

