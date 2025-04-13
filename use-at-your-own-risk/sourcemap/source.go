package sourcemap

import "github.com/Zzzen/typescript-go/use-at-your-own-risk/core"

type Source interface {
	Text() string
	FileName() string
	LineMap() []core.TextPos
}
