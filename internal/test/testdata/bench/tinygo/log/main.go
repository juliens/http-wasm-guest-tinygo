package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
)

func main() {
	handler.HandleRequestFn = log
	guest.SetExports(handler.GetExports())
}

func log(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	handler.Host.Log(api.LogLevelInfo, "hello world")
	return // this is a benchmark, so skip the next handler.
}
