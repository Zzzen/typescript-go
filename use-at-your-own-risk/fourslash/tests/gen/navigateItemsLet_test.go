package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavigateItemsLet(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true
let [|c = 10|];
function foo() {
    let [|d = 10|];
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "c",
			Preferences: nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:     "c",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		}, {
			Pattern:     "d",
			Preferences: nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:          "d",
					Kind:          lsproto.SymbolKindVariable,
					Location:      f.Ranges()[1].LSLocation(),
					ContainerName: PtrTo("foo"),
				},
			}),
		},
	})
}
