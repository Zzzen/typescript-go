package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsClassWithStaticThisAccess(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|class /*0*/[|{| "isWriteAccess": true, "isDefinition": true, "contextRangeIndex": 0 |}C|] {
    static s() {
        /*1*/[|this|];
    }
    static get f() {
        return /*2*/[|this|];

        function inner() { this; }
        class Inner { x = this; }
    }
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "0", "1", "2")
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1])
}
