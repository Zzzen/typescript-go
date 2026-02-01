package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestMultilineCommentBeforeOpenBrace(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function test() /*1*//* %^ */
{
    if (true) /*2*//* %^ */
    {
    }
}
function a() {
    /* %^ */ }/*3*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `function test() /* %^ */ {`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `    if (true) /* %^ */ {`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `}`)
}
