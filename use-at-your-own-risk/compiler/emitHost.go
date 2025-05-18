package compiler

import (
	"context"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/printer"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
)

var _ printer.EmitHost = (*emitHost)(nil)

// NOTE: emitHost operations must be thread-safe
type emitHost struct {
	program *Program
}

func (host *emitHost) Options() *core.CompilerOptions { return host.program.Options() }
func (host *emitHost) SourceFiles() []*ast.SourceFile { return host.program.SourceFiles() }
func (host *emitHost) GetCurrentDirectory() string    { return host.program.host.GetCurrentDirectory() }
func (host *emitHost) CommonSourceDirectory() string  { return host.program.CommonSourceDirectory() }
func (host *emitHost) UseCaseSensitiveFileNames() bool {
	return host.program.host.FS().UseCaseSensitiveFileNames()
}

func (host *emitHost) IsEmitBlocked(file string) bool {
	// !!!
	return false
}

func (host *emitHost) WriteFile(fileName string, text string, writeByteOrderMark bool, _ []*ast.SourceFile, _ *printer.WriteFileData) error {
	return host.program.host.FS().WriteFile(fileName, text, writeByteOrderMark)
}

func (host *emitHost) GetEmitResolver(file *ast.SourceFile, skipDiagnostics bool) printer.EmitResolver {
	// The context and done function don't matter in tsc, currently the only caller of this function.
	// But if this ever gets used by LSP code, we'll need to thread the context properly and pass the
	// done function to the caller to ensure resources are cleaned up at the end of the request.
	checker, done := host.program.GetTypeCheckerForFile(context.TODO(), file)
	defer done()
	return checker.GetEmitResolver(file, skipDiagnostics)
}

func (host *emitHost) GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData {
	return host.program.GetSourceFileMetaData(path)
}
