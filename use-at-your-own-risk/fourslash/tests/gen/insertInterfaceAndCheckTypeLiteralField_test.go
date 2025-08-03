package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestInsertInterfaceAndCheckTypeLiteralField(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*addC*/
interface G<T, U> { }
var v2: G<{ a: /*checkParam*/C }, C>;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "addC")
	f.Insert(t, "interface C { }")
	f.GoToMarker(t, "checkParam")
	f.VerifyQuickInfoExists(t)
}
