package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

func main() {
	handler.HandleRequestFn = getHeaderValues
	guest.SetExports(handler.GetExports())
}

func getHeaderValues(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	_ = req.Headers().GetAll("Accept")
	return // this is a benchmark, so skip the next handler.
}
