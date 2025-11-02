package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionsImport_weirdDefaultSynthesis(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @esModuleInterop: false
// @allowSyntheticDefaultImports: false
// @Filename: /collection.ts
class Collection {
  public static readonly default: typeof Collection = Collection;
}
export = Collection as typeof Collection & { default: typeof Collection };
// @Filename: /index.ts
Colle/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:        "Collection",
		Source:      "./collection",
		Description: "Add import from \"./collection\"",
		NewFileContent: PtrTo(`import Collection = require("./collection");

Colle`),
	})
}
