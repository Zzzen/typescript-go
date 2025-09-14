package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameThis(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f([|this|]) {
    return [|this|];
}
this/**/;
const _ = { [|[|{| "contextRangeIndex": 2 |}this|]: 0|] }.[|this|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyRenameFailed(t, nil /*preferences*/)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[0], f.Ranges()[1], f.Ranges()[3], f.Ranges()[4])
}
