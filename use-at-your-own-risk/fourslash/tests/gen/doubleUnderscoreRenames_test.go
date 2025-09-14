package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestDoubleUnderscoreRenames(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: fileA.ts
[|export function [|{| "contextRangeIndex": 0 |}__foo|]() {
}|]

// @Filename: fileB.ts
[|import { [|{| "contextRangeIndex": 2 |}__foo|] as bar } from "./fileA";|]

bar();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "__foo")
}
