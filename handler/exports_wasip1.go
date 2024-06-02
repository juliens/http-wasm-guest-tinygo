//go:build wasip1

package handler

import (
	"github.com/juliens/wasm-goexport/guest"
	wazeroapi "github.com/tetratelabs/wazero/api"
)

func GetExports() []*guest.Function {
	return []*guest.Function{
		{
			ModuleName: "",
			Name:       "handle_request",
			Fn:         handleRequest,
			Results:    []wazeroapi.ValueType{wazeroapi.ValueTypeI64},
		},
		{
			ModuleName: "",
			Name:       "handle_response",
			Fn:         handleResponse,
			Params:     []wazeroapi.ValueType{wazeroapi.ValueTypeI32, wazeroapi.ValueTypeI32},
		},
	}
}
