wasm-build:
	GOOS=js GOARCH=wasm go build -o ./public/main.wasm
	cp $(GOROOT)/misc/wasm/wasm_exec.js ./public
	cp -r assets ./public

wasm-run:
	make wasm-build
	npx serve ./public

run:
	go run .
