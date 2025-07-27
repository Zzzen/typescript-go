package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionFunctionOverloadsInClass(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class clsInOverload {
    static fnOverload();
    static [|/*staticFunctionOverload*/fnOverload|](foo: string);
    static /*staticFunctionOverloadDefinition*/fnOverload(foo: any) { }
    public [|/*functionOverload*/fnOverload|](): any;
    public fnOverload(foo: string);
    public /*functionOverloadDefinition*/fnOverload(foo: any) { return "foo" }

    constructor() { }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "staticFunctionOverload", "functionOverload")
}
