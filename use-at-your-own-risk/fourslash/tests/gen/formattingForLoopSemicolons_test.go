package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormattingForLoopSemicolons(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/for (;;) { }
/*2*/for (var x;x<0;x++) { }
/*3*/for (var x ;x<0 ;x++) { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `for (; ;) { }`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `for (var x; x < 0; x++) { }`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `for (var x; x < 0; x++) { }`)
	opts444 := f.GetOptions()
	opts444.FormatCodeSettings.InsertSpaceAfterSemicolonInForStatements = core.TSFalse
	f.Configure(t, opts444)
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `for (;;) { }`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `for (var x;x < 0;x++) { }`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `for (var x;x < 0;x++) { }`)
}
