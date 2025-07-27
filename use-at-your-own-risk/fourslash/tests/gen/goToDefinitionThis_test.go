package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionThis(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f(/*fnDecl*/this: number) {
    return [|/*fnUse*/this|];
}
class /*cls*/C {
    constructor() { return [|/*clsUse*/this|]; }
    get self(/*getterDecl*/this: number) { return [|/*getterUse*/this|]; }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "fnUse", "clsUse", "getterUse")
}
