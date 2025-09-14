package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameAcrossMultipleProjects(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: a.ts
[|var [|{| "contextRangeIndex": 0 |}x|]: number;|]
//@Filename: b.ts
/// <reference path="a.ts" />
[|x|]++;
//@Filename: c.ts
/// <reference path="a.ts" />
[|x|]++;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "x")
}
