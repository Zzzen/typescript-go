package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameBindingElementInitializerExternal(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|const [|{| "contextRangeIndex": 0 |}external|] = true;|]

function f({
    lvl1 = [|external|],
    nested: { lvl2 = [|external|]},
    oldName: newName = [|external|]
}) {}

const {
    lvl1 = [|external|],
    nested: { lvl2 = [|external|]},
    oldName: newName = [|external|]
} = obj;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "external")
}
