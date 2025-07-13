package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestReferencesForStatic(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: referencesOnStatic_1.ts
var n = 43;

class foo {
    /*1*/static /*2*/n = '';

    public bar() {
        foo./*3*/n = "'";
        if(foo./*4*/n) {
            var x = foo./*5*/n;
        }
    }
}

class foo2 {
    private x = foo./*6*/n;
    constructor() {
        foo./*7*/n = x;
    }

    function b(n) {
        n = foo./*8*/n;
    }
}
// @Filename: referencesOnStatic_2.ts
var q = foo./*9*/n;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4", "5", "6", "7", "8", "9")
}
