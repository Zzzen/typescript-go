package ls

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/compiler"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs"
)

type Host interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	NewLine() string
	Trace(msg string)
	GetProjectVersion() int
	// GetRootFileNames was called GetScriptFileNames in the original code.
	GetRootFileNames() []string
	// GetCompilerOptions was called GetCompilationSettings in the original code.
	GetCompilerOptions() *core.CompilerOptions
	GetSourceFile(fileName string, path tspath.Path, languageVersion core.ScriptTarget) *ast.SourceFile
	// This responsibility was moved from the language service to the project,
	// because they were bidirectionally interdependent.
	GetProgram() *compiler.Program
	GetDefaultLibraryPath() string
	GetPositionEncoding() lsproto.PositionEncodingKind
	GetScriptInfo(fileName string) ScriptInfo
}
