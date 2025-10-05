package ls

import "github.com/Zzzen/typescript-go/use-at-your-own-risk/sourcemap"

type Host interface {
	UseCaseSensitiveFileNames() bool
	ReadFile(path string) (contents string, ok bool)
	Converters() *Converters
	GetECMALineInfo(fileName string) *sourcemap.ECMALineInfo
}
