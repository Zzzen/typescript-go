package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormattingMappedType(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*generic*/type t  < T  > =   {
/*map*/   [   P   in   keyof    T  ]   :   T  [  P  ]
};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "generic")
	f.VerifyCurrentLineContent(t, `type t<T> = {`)
	f.GoToMarker(t, "map")
	f.VerifyCurrentLineContent(t, `    [P in keyof T]: T[P]`)
}
