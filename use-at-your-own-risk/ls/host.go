package ls

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/format"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/autoimport"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsconv"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/sourcemap"
)

type Host interface {
	UseCaseSensitiveFileNames() bool
	ReadFile(path string) (contents string, ok bool)
	Converters() *lsconv.Converters
	UserPreferences() *lsutil.UserPreferences
	FormatOptions() *format.FormatCodeSettings
	GetECMALineInfo(fileName string) *sourcemap.ECMALineInfo
	AutoImportRegistry() *autoimport.Registry
}
