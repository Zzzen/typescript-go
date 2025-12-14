package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestInsertInterfaceAndCheckTypeLiteralField(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*addC*/
interface G<T, U> { }
var v2: G<{ a: /*checkParam*/C }, C>;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "addC")
	f.Insert(t, "interface C { }")
	f.GoToMarker(t, "checkParam")
	f.VerifyQuickInfoExists(t)
}
