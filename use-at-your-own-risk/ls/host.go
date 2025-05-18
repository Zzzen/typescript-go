package ls

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/compiler"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
)

type Host interface {
	GetProgram() *compiler.Program
	GetPositionEncoding() lsproto.PositionEncodingKind
	GetLineMap(fileName string) *LineMap
}
