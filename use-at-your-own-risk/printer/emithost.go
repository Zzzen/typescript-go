package printer

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tsoptions"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
)

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	UseCaseSensitiveFileNames() bool
	GetCurrentDirectory() string
	CommonSourceDirectory() string
	IsEmitBlocked(file string) bool
	WriteFile(fileName string, text string, writeByteOrderMark bool) error
	GetEmitModuleFormatOfFile(file ast.HasFileName) core.ModuleKind
	GetEmitResolver() EmitResolver
	GetOutputAndProjectReference(path tspath.Path) *tsoptions.OutputDtsAndProjectReference
	IsSourceFileFromExternalLibrary(file *ast.SourceFile) bool
}
