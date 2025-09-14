package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameStringPropertyNames(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var o = {
    [|[|{| "contextRangeIndex": 0 |}prop|]: 0|]
};

o = {
    [|"[|{| "contextRangeIndex": 2 |}prop|]": 1|]
};

o["[|prop|]"];
o['[|prop|]'];
o.[|prop|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "prop")
}
