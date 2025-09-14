package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameDestructuringDeclarationInForOf(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    [|[|{| "contextRangeIndex": 0 |}property1|]: number;|]
    property2: string;
}
var elems: I[];

for ([|let { [|{| "contextRangeIndex": 2 |}property1|] } of elems|]) {
    [|property1|]++;
}
for ([|let { [|{| "contextRangeIndex": 5 |}property1|]: p2 } of elems|]) {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[6], f.Ranges()[3], f.Ranges()[4])
}
