package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavto_excludeLib3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /index.ts
[|function parseInt(s: string): number {}|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "parseInt",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "parseInt",
					Kind:     lsproto.SymbolKindFunction,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
}
