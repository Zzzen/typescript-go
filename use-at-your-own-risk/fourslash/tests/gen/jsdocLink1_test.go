package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsdocLink1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
}
/**
 * {@link C}
 * @wat Makes a {@link C}. A default one.
 * {@link C()}
 * {@link C|postfix text}
 * {@link unformatted postfix text}
 * @see {@link C} its great
 */
function /**/CC() {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
