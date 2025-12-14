package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormattingSpaceAfterCommaBeforeOpenParen(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo(a,(b))/*1*/
foo(a,(<b>c).d)/*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, ";")
	f.VerifyCurrentLineContentIs(t, "foo(a, (b));")
	f.GoToMarker(t, "2")
	f.Insert(t, ";")
	f.VerifyCurrentLineContentIs(t, "foo(a, (<b>c).d);")
}
