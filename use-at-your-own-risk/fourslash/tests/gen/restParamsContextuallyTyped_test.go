package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRestParamsContextuallyTyped(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var foo: Function = function (/*1*/a, /*2*/b, /*3*/c) { };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "(parameter) a: any", "")
	f.VerifyQuickInfoAt(t, "2", "(parameter) b: any", "")
	f.VerifyQuickInfoAt(t, "3", "(parameter) c: any", "")
}
