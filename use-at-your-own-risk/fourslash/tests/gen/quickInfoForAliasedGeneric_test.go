package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoForAliasedGeneric(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module M {
    export module N {
        export class C<T> { }
        export class D { }
    }
}
import d = M.N;
var /*1*/aa: d.C<number>;
var /*2*/bb: d.D;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "var aa: d.C<number>", "")
	f.VerifyQuickInfoAt(t, "2", "var bb: d.D", "")
}
