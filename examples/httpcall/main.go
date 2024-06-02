//go:build wasip1

package main

import (
	"io"
	"log"
	"net/http"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
	"github.com/juliens/wasm-goexport/guest"
	_ "github.com/stealthrocket/net/http"
)

func main() {
	handler.HandleRequestFn = handleRequest
	guest.SetExports(handler.GetExports())

}

// handleRequest implements a simple HTTP router.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	url := handler.Host.GetConfig()
	response, err := http.Get(string(url))
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatal(response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Serve a static response
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().Write(body)
	return // skip the next handler, as we wrote a response.
}
