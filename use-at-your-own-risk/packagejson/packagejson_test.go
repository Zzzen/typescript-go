package packagejson_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	json2 "github.com/go-json-experiment/json"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ast"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/packagejson"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/parser"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/repo"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil/filefixture"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/tspath"
)

var packageJsonFixtures = []filefixture.Fixture{
	filefixture.FromFile("package.json", filepath.Join(repo.RootPath, "package.json")),
	filefixture.FromFile("date-fns.json", filepath.Join(repo.TestDataPath, "fixtures", "packagejson", "date-fns.json")),
}

func BenchmarkPackageJSON(b *testing.B) {
	for _, f := range packageJsonFixtures {
		f.SkipIfNotExist(b)
		content := []byte(f.ReadFile(b))
		b.Run("UnmarshalJSON", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				for b.Loop() {
					var p packagejson.Fields
					if err := json.Unmarshal(content, &p); err != nil {
						b.Fatal(err)
					}
				}
			})
		})

		b.Run("UnmarshalJSONV2", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				for b.Loop() {
					var p packagejson.Fields
					if err := json2.Unmarshal(content, &p); err != nil {
						b.Fatal(err)
					}
				}
			})
		})

		b.Run("ParseJSONText", func(b *testing.B) {
			b.Run(f.Name(), func(b *testing.B) {
				fileName := "/" + f.Name()
				for b.Loop() {
					parser.ParseSourceFile(ast.SourceFileParseOptions{
						FileName: fileName,
						Path:     tspath.Path(fileName),
					}, string(content), core.ScriptKindJSON)
				}
			})
		})
	}
}
