package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormattingOnModuleIndentation(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `  namespace     Foo    {
    export    namespace    A  .   B  .   C     {      }/**/
               }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToBOF(t)
	f.VerifyCurrentLineContent(t, `namespace Foo {`)
	f.GoToMarker(t, "")
	f.VerifyCurrentLineContent(t, `    export namespace A.B.C { }`)
	f.GoToEOF(t)
	f.VerifyCurrentLineContent(t, `}`)
}
