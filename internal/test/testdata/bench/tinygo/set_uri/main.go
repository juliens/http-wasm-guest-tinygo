package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

func main() {
	handler.HandleRequestFn = setURI
	guest.SetExports(handler.GetExports())
}

func setURI(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	req.SetURI("/v1.0/hello")
	return // this is a benchmark, so skip the next handler.
}
