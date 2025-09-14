package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameBindingElementInitializerProperty(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f([|{[|{| "contextRangeIndex": 0 |}required|], optional = [|required|]}: {[|[|{| "contextRangeIndex": 3 |}required|]: number,|] optional?: number}|]) {
    console.log("required", [|required|]);
    console.log("optional", optional);
}

f({[|[|{| "contextRangeIndex": 6 |}required|]: 10|]});`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[2], f.Ranges()[5], f.Ranges()[4], f.Ranges()[7])
}
