package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

var body = []byte("hello world")

func main() {
	handler.HandleRequestFn = writeBody
	guest.SetExports(handler.GetExports())
}

func writeBody(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Body().Write(body)
	return // this is a benchmark, so skip the next handler.
}
