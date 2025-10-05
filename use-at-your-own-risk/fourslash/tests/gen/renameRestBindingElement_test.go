package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameRestBindingElement(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    a: number;
    b: number;
    c: number;
}
function foo([|{ a, ...[|{| "contextRangeIndex": 0 |}rest|] }: I|]) {
    [|rest|];
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, &ls.UserPreferences{UseAliasesForRename: core.TSTrue}, f.Ranges()[1])
}
