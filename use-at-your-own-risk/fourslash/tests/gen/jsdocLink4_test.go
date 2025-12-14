package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsdocLink4(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare class I {
  /** {@link I} */
  bar/*1*/(): void
}
/** {@link I} */
var n/*2*/ = 1
/**
 * A real, very serious {@link I to an interface}. Right there.
 * @param x one {@link Pos here too}
 */
function f(x) {
}
f/*3*/()
type Pos = [number, number]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}
