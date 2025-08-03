package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsDocPropertyDescription12(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type SymbolAlias = {
    /** Something generic */
    [p: symbol]: string;
}
function symbolAlias(e: SymbolAlias) {
    console.log(e./*symbolAlias*/anything);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "symbolAlias", "any", "")
}
