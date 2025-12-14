package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOccurrencesClassExpressionPrivate(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `let A = class Foo {
    [|private|] foo;
    [|private|] private;
    constructor([|private|] y: string, public x: string) {
    }
    [|private|] method() { }
    public method2() { }
    [|private|] static static() { }
}

let B = class D {
    constructor(private x: number) {
    }
    private test() {}
    public test2() {}
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
