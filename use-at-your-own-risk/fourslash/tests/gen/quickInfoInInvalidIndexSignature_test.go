package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoInInvalidIndexSignature(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function method() { var /**/dictionary = <{ [index]: string; }>{}; }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(local var) dictionary: {\n    [x: number]: string;\n}", "")
}
