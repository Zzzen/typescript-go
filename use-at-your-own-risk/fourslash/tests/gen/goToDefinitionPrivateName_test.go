package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionPrivateName(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class A {
    [|/*pnMethodDecl*/#method|]() { }
    [|/*pnFieldDecl*/#foo|] = 3;
    get [|/*pnPropGetDecl*/#prop|]() { return ""; }
    set [|/*pnPropSetDecl*/#prop|](value: string) {  }
    constructor() {
        this.[|/*pnFieldUse*/#foo|]
        this.[|/*pnMethodUse*/#method|]
        this.[|/*pnPropUse*/#prop|]
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, true, "pnFieldUse", "pnMethodUse", "pnPropUse")
}
