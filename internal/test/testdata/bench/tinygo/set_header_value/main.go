package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

func main() {
	handler.HandleRequestFn = setHeader
	guest.SetExports(handler.GetExports())
}

func setHeader(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Headers().Set("Content-Type", "text/plain")
	return // this is a benchmark, so skip the next handler.
}
