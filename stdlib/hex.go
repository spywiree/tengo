package stdlib

import (
	"encoding/hex"

	"github.com/spywiree/tengo"
)

var hexModule = map[string]tengo.Object{
	"encode": &tengo.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &tengo.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
