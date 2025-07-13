package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestReferencesForMergedDeclarations3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|class /*class*/[|testClass|] {
    static staticMethod() { }
    method() { }
}|]

[|module /*module*/[|testClass|] {
    export interface Bar {

    }
}|]

var c1: [|testClass|];
var c2: [|testClass|].Bar;
[|testClass|].staticMethod();
[|testClass|].prototype.method();
[|testClass|].bind(this);
new [|testClass|]();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "module", "class")
}
