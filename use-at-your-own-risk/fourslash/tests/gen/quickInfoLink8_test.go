package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoLink8(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const A = 123;
/**
 * See {@link A | constant A} instead
 */
const /**/B = 456;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
