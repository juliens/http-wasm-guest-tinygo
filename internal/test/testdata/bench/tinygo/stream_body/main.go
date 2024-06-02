package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

func main() {
	handler.HandleRequestFn = readBody
	guest.SetExports(handler.GetExports())
}

var empty = make([]byte, 0)

func readBody(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	size, _ := req.Body().Read(empty)
	_, _ = req.Body().Read(make([]byte, size))
	return // this is a benchmark, so skip the next handler.
}
