package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsOfConstructor2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class A {
    /*a*/constructor(s: string) {}
}
class B extends A {
    /*b*/constructor() { super(""); }
}
class C extends B {
    /*c*/constructor() {
        super();
    }
}
class D extends B { }
const a = new A("a");
const b = new B();
const c = new C();
const d = new D();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoErrors(t)
	f.VerifyBaselineFindAllReferences(t, "a", "b", "c")
}
