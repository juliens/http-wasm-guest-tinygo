//go:build tinygo.wasm

package handler

import (
	"github.com/juliens/wasm-goexport/guest"
)

func GetExports() []*guest.Function {
	return nil
}
