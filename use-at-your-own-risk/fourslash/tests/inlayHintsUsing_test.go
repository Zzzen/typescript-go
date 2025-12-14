package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestInlayHintsUsing(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @target: esnext
using _defer = {
	[Symbol.dispose]() {},
};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{
		IncludeInlayVariableTypeHints: true,
	}})
}
