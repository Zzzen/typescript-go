package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameDestructuringAssignmentNestedInArrayLiteral(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    [|[|{| "contextRangeIndex": 0 |}property1|]: number;|]
    property2: string;
}
var elems: I[], p1: number, [|[|{| "contextRangeIndex": 2 |}property1|]: number|];
[|[{ [|{| "contextRangeIndex": 4 |}property1|]: p1 }] = elems;|]
[|[{ [|{| "contextRangeIndex": 6 |}property1|] }] = elems;|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[5], f.Ranges()[3], f.Ranges()[7])
}
