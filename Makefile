.PHONY: build, optimize, all

build:
	tinygo build -o main.wasm -no-debug -target wasm ./main.go

optimize:
	wasm2wat main.wasm -o wasm.wat
	wat2wasm wasm.wat -o main.wasm
	rm ./wasm.wat

all:
	make build
	make optimize
