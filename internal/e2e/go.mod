module github.com/http-wasm/http-wasm-guest-tinygo/e2e

go 1.22.3

require (
	github.com/http-wasm/http-wasm-guest-tinygo v0.0.0
	github.com/http-wasm/http-wasm-host-go v0.5.1
	github.com/stretchr/testify v1.8.4
	github.com/tetratelabs/wazero v1.7.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/juliens/wasm-goexport v0.0.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stealthrocket/wasi-go v0.8.0 // indirect
	github.com/stealthrocket/wazergo v0.19.1 // indirect
	golang.org/x/sys v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/http-wasm/http-wasm-guest-tinygo => ../../

replace github.com/http-wasm/http-wasm-host-go => ../../../../plugins-sandbox/http-wasm-host-go
