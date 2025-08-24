package ls

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/compiler"
)

type Host interface {
	GetProgram() *compiler.Program
}
