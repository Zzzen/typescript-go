package compiler

import (
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/parser"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tsoptions"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs/cachedvfs"
)

type CompilerHost interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	Trace(msg string)
	GetSourceFile(opts ast.SourceFileParseOptions) *ast.SourceFile
	GetResolvedProjectReference(fileName string, path tspath.Path) *tsoptions.ParsedCommandLine
}

var _ CompilerHost = (*compilerHost)(nil)

type compilerHost struct {
	currentDirectory    string
	fs                  vfs.FS
	defaultLibraryPath  string
	extendedConfigCache tsoptions.ExtendedConfigCache
	trace               func(msg string)
}

func NewCachedFSCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	trace func(msg string),
) CompilerHost {
	return NewCompilerHost(currentDirectory, cachedvfs.From(fs), defaultLibraryPath, extendedConfigCache, trace)
}

func NewCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	trace func(msg string),
) CompilerHost {
	if trace == nil {
		trace = func(msg string) {}
	}
	return &compilerHost{
		currentDirectory:    currentDirectory,
		fs:                  fs,
		defaultLibraryPath:  defaultLibraryPath,
		extendedConfigCache: extendedConfigCache,
		trace:               trace,
	}
}

func (h *compilerHost) FS() vfs.FS {
	return h.fs
}

func (h *compilerHost) DefaultLibraryPath() string {
	return h.defaultLibraryPath
}

func (h *compilerHost) GetCurrentDirectory() string {
	return h.currentDirectory
}

func (h *compilerHost) Trace(msg string) {
	h.trace(msg)
}

func (h *compilerHost) GetSourceFile(opts ast.SourceFileParseOptions) *ast.SourceFile {
	text, ok := h.FS().ReadFile(opts.FileName)
	if !ok {
		return nil
	}
	return parser.ParseSourceFile(opts, text, core.GetScriptKindFromFileName(opts.FileName))
}

func (h *compilerHost) GetResolvedProjectReference(fileName string, path tspath.Path) *tsoptions.ParsedCommandLine {
	commandLine, _ := tsoptions.GetParsedCommandLineOfConfigFilePath(fileName, path, nil, h, h.extendedConfigCache)
	return commandLine
}
