package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsPrivateNameProperties(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
    /*1*/#foo = 10;
    constructor() {
        this./*2*/#foo = 20;
        /*3*/#foo in this;
    }
}
class D extends C {
    constructor() {
        super()
        this.#foo = 20;
    }
}
class E {
    /*4*/#foo: number;
    constructor() {
        this./*5*/#foo = 20;
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4", "5")
}
