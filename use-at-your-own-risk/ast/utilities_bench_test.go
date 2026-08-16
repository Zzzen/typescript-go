package ast_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/parser"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil/fixtures"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs/osvfs"
)

func BenchmarkGetCombinedFlags(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)
			scriptKind := core.GetScriptKindFromFileName(fileName)

			sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
				FileName: fileName,
				Path:     path,
			}, sourceText, scriptKind)

			var decls []*ast.Node
			var collect ast.Visitor
			collect = func(n *ast.Node) bool {
				if ast.IsDeclaration(n) {
					decls = append(decls, n)
				}
				n.ForEachChild(collect)
				return false
			}
			sourceFile.AsNode().ForEachChild(collect)

			for b.Loop() {
				for _, n := range decls {
					_ = ast.GetCombinedNodeFlags(n)
					_ = ast.GetCombinedModifierFlags(n)
				}
			}
		})
	}
}
