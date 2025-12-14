package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsObjectBindingElementPropertyName04(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    /*0*/property1: number;
    property2: string;
}

function f({ /*1*/property1: p1 }: I,
           { /*2*/property1 }: I,
           { property1: p2 }) {

    return /*3*/property1 + 1;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "0", "1", "2", "3")
}
