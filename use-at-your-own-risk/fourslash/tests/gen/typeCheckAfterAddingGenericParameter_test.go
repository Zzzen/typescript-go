package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTypeCheckAfterAddingGenericParameter(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f<x, x>() { }
function f2<X, X>(b: X): X { return null; }
class C<X> {
    public f<x, x>() {}
f2<X>(b): X { return null; }
}

interface I<X, X> {
    f<X/*addTypeParam*/>();
    f2<X>(/*addParam*/a: X): X;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "addParam")
	f.Insert(t, ", X")
	f.GoToMarker(t, "addTypeParam")
	f.Insert(t, ", X")
}
