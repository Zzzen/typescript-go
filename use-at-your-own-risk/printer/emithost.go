package printer

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
)

type WriteFileData struct {
	SourceMapUrlPos int
	// BuildInfo BuildInfo
	Diagnostics      []*ast.Diagnostic
	DiffersOnlyInMap bool
	SkippedDtsWrite  bool
}

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	SourceFileMetaDataProvider
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	UseCaseSensitiveFileNames() bool
	GetCurrentDirectory() string
	CommonSourceDirectory() string
	IsEmitBlocked(file string) bool
	WriteFile(fileName string, text string, writeByteOrderMark bool, relatedSourceFiles []*ast.SourceFile, data *WriteFileData) error
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
	GetEmitResolver(file *ast.SourceFile, skipDiagnostics bool) EmitResolver
}
