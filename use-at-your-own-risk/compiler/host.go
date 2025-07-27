package compiler

import (
	"sync"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/collections"
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
	currentDirectory        string
	fs                      vfs.FS
	defaultLibraryPath      string
	extendedConfigCache     *collections.SyncMap[tspath.Path, *tsoptions.ExtendedConfigCacheEntry]
	extendedConfigCacheOnce sync.Once
}

func NewCachedFSCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache *collections.SyncMap[tspath.Path, *tsoptions.ExtendedConfigCacheEntry],
) CompilerHost {
	return NewCompilerHost(currentDirectory, cachedvfs.From(fs), defaultLibraryPath, extendedConfigCache)
}

func NewCompilerHost(
	currentDirectory string,
	fs vfs.FS,
	defaultLibraryPath string,
	extendedConfigCache *collections.SyncMap[tspath.Path, *tsoptions.ExtendedConfigCacheEntry],
) CompilerHost {
	return &compilerHost{
		currentDirectory:    currentDirectory,
		fs:                  fs,
		defaultLibraryPath:  defaultLibraryPath,
		extendedConfigCache: extendedConfigCache,
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
	//!!! TODO: implement
}

func (h *compilerHost) GetSourceFile(opts ast.SourceFileParseOptions) *ast.SourceFile {
	text, ok := h.FS().ReadFile(opts.FileName)
	if !ok {
		return nil
	}
	return parser.ParseSourceFile(opts, text, core.GetScriptKindFromFileName(opts.FileName))
}

func (h *compilerHost) GetResolvedProjectReference(fileName string, path tspath.Path) *tsoptions.ParsedCommandLine {
	h.extendedConfigCacheOnce.Do(func() {
		if h.extendedConfigCache == nil {
			h.extendedConfigCache = &collections.SyncMap[tspath.Path, *tsoptions.ExtendedConfigCacheEntry]{}
		}
	})
	commandLine, _ := tsoptions.GetParsedCommandLineOfConfigFilePath(fileName, path, nil, h, h.extendedConfigCache)
	return commandLine
}
