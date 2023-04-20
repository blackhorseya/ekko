//go:generate wire
//go:build wireinject

package tokenx

import (
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewJwtx)

func CreateTokenizer(opts *Options) Tokenizer {
	panic(wire.Build(testProviderSet))
}
