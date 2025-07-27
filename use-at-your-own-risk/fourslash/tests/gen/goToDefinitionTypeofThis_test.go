package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionTypeofThis(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f(/*fnDecl*/this: number) {
    type X = typeof [|/*fnUse*/this|];
}
class /*cls*/C {
    constructor() { type X = typeof [|/*clsUse*/this|]; }
    get self(/*getterDecl*/this: number) { type X = typeof [|/*getterUse*/this|]; }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "fnUse", "clsUse", "getterUse")
}
