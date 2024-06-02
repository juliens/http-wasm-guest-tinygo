package test

import (
	_ "embed"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed testdata/bench/tinygo/log/main.wasm
var BinBenchLogTinyGo []byte

//go:embed testdata/bench/tinygo/log/main_go.wasm
var BinBenchLogGo []byte

//go:embed testdata/bench/wat/log.wasm
var BinBenchLogWat []byte

//go:embed testdata/bench/tinygo/get_uri/main.wasm
var BinBenchGetURITinyGo []byte

//go:embed testdata/bench/tinygo/get_uri/main_go.wasm
var BinBenchGetURIGo []byte

//go:embed testdata/bench/wat/get_uri.wasm
var BinBenchGetURIWat []byte

//go:embed testdata/bench/tinygo/set_uri/main.wasm
var BinBenchSetURITinyGo []byte

//go:embed testdata/bench/tinygo/set_uri/main_go.wasm
var BinBenchSetURIGo []byte

//go:embed testdata/bench/wat/set_uri.wasm
var BinBenchSetURIWat []byte

//go:embed testdata/bench/tinygo/get_header_names/main.wasm
var BinBenchGetHeaderValuesNamesTinyGo []byte

//go:embed testdata/bench/tinygo/get_header_names/main_go.wasm
var BinBenchGetHeaderValuesNamesGo []byte

//go:embed testdata/bench/wat/get_header_names.wasm
var BinBenchGetHeaderValuesNamesWat []byte

//go:embed testdata/bench/tinygo/get_header_values/main.wasm
var BinBenchGetHeaderValuesTinyGo []byte

//go:embed testdata/bench/tinygo/get_header_values/main_go.wasm
var BinBenchGetHeaderValuesGo []byte

//go:embed testdata/bench/wat/get_header_values.wasm
var BinBenchGetHeaderValuesWat []byte

//go:embed testdata/bench/tinygo/set_header_value/main.wasm
var BinBenchSetHeaderValueTinyGo []byte

//go:embed testdata/bench/tinygo/set_header_value/main_go.wasm
var BinBenchSetHeaderValueGo []byte

//go:embed testdata/bench/wat/set_header_value.wasm
var BinBenchSetHeaderValueWat []byte

//go:embed testdata/bench/tinygo/add_header_value/main.wasm
var BinBenchAddHeaderValueTinyGo []byte

//go:embed testdata/bench/tinygo/add_header_value/main_go.wasm
var BinBenchAddHeaderValueGo []byte

//go:embed testdata/bench/wat/add_header_value.wasm
var BinBenchAddHeaderValueWat []byte

//go:embed testdata/bench/tinygo/remove_header/main.wasm
var BinBenchRemoveHeaderTinyGo []byte

//go:embed testdata/bench/tinygo/remove_header/main_go.wasm
var BinBenchRemoveHeaderGo []byte

//go:embed testdata/bench/wat/remove_header.wasm
var BinBenchRemoveHeaderWat []byte

//go:embed testdata/bench/tinygo/read_body/main.wasm
var BinBenchReadBodyTinyGo []byte

//go:embed testdata/bench/tinygo/read_body/main_go.wasm
var BinBenchReadBodyGo []byte

//go:embed testdata/bench/tinygo/write_body/main.wasm
var BinBenchWriteBodyTinyGo []byte

//go:embed testdata/bench/tinygo/write_body/main_go.wasm
var BinBenchWriteBodyGo []byte

//go:embed testdata/bench/wat/write_body.wasm
var BinBenchWriteBodyWat []byte

//go:embed testdata/bench/wat/read_body.wasm
var BinBenchReadBodyWat []byte

//go:embed testdata/bench/tinygo/read_body_stream/main.wasm
var BinBenchReadBodyStreamTinyGo []byte

//go:embed testdata/bench/tinygo/read_body_stream/main_go.wasm
var BinBenchReadBodyStreamGo []byte

//go:embed testdata/bench/wat/read_body_stream.wasm
var BinBenchReadBodyStreamWat []byte

//go:embed testdata/bench/tinygo/set_status_code/main.wasm
var BinBenchSetStatusCodeTinyGo []byte

//go:embed testdata/bench/tinygo/set_status_code/main_go.wasm
var BinBenchSetStatusCodeGo []byte

//go:embed testdata/bench/wat/set_status_code.wasm
var BinBenchSetStatusCodeWat []byte

var BinExampleRouter = func() []byte {
	return binExample("router")
}()

var BinExampleRouterGo = func() []byte {
	return binGoExample("router")
}()

var BinExampleHTTPCallGo = func() []byte {
	return binGoExample("httpcall")
}()

//go:embed testdata/e2e/wat/router.wasm
var BinExampleRouterWat []byte

var BinExampleWASI = func() []byte {
	return binExample("wasi")
}()

var BinExampleWASIGo = func() []byte {
	return binGoExample("wasi")
}()

//go:embed testdata/e2e/wat/wasi.wasm
var BinExampleWASIWat []byte

func RequestExampleWASI(url string) (req *http.Request) {
	body := strings.NewReader(`{"hello": "panda"}`)
	req, _ = http.NewRequest(http.MethodPost, url+"/v1.0/hi?name=panda", body)
	req.Header.Set("Content-Type", "application/json")
	req.Host = "localhost"
	return
}

var HandlerExampleWASI = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Set-Cookie", "a=b") // router of multiple headers
	w.Header().Add("Set-Cookie", "c=d")

	// Use chunked encoding so we can set a test trailer
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Trailer", "grpc-status")
	w.Header().Set(http.TrailerPrefix+"grpc-status", "1")
	w.Write([]byte(`{"hello": "world"}`)) // nolint
})

//go:embed testdata/e2e/tinygo/features/main.wasm
var BinE2EFeaturesTinyGo []byte

//go:embed testdata/e2e/tinygo/features/main_go.wasm
var BinE2EFeaturesGo []byte

//go:embed testdata/e2e/tinygo/handle_response/main.wasm
var BinHandleResponse []byte

//go:embed testdata/e2e/tinygo/handle_response/main_go.wasm
var BinHandleResponseGo []byte

//go:embed testdata/e2e/tinygo/log/main.wasm
var BinE2ELog []byte

//go:embed testdata/e2e/tinygo/log/main_go.wasm
var BinE2ELogGo []byte

// binExample instead of go:embed as files aren't relative to this directory.
func binExample(name string) []byte {
	p := path.Join("..", "..", "examples", name, "main.wasm")
	if wasm, err := os.ReadFile(p); err != nil {
		log.Panicln(err)
		return nil
	} else {
		return wasm
	}
}

// binExample instead of go:embed as files aren't relative to this directory.
func binGoExample(name string) []byte {
	p := path.Join("..", "..", "examples", name, "main_go.wasm")
	if wasm, err := os.ReadFile(p); err != nil {
		log.Panicln(err)
		return nil
	} else {
		return wasm
	}
}
