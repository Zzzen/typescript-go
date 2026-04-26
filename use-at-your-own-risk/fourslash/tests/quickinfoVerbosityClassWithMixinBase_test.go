package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickinfoVerbosityClassWithMixinBase(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `
class Base {}

declare const Mixin: new () => Base & { mixed: string };

class Derived/*1*/ extends Mixin {}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"1": {0, 1}})
}
