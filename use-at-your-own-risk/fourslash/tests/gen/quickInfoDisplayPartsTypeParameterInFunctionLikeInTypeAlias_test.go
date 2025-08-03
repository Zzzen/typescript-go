package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoDisplayPartsTypeParameterInFunctionLikeInTypeAlias(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type MixinCtor<A> = new () => /*0*/A & { constructor: MixinCtor</*1*/A> };
type MixinCtor<A> = new () => A & { constructor: { constructor: MixinCtor</*2*/A> } };`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
