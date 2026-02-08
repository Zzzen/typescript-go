package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestExtendsTArray(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: false
interface I1<T> {
    (a: T): T;
}
interface I2<T> extends I1<T[]> {
    b: T;
}
var x: I2<Date>;
var /**/y = x(undefined); // Typeof y should be Date[]
y.length;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "var y: Date[]", "")
	f.VerifyNoErrors(t)
}
