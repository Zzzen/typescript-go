package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameLocationsForFunctionExpression02(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f() {

}
var x = [|function [|{| "contextRangeIndex": 0 |}f|](g: any, h: any) {

    let helper = function f(): any { f(); }

    let foo = () => [|f|]([|f|], g);
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "f")
}
